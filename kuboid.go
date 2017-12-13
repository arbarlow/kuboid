package kuboid

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/api"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

const reqsFmt = "sum(rate(request_count[%s])) by (source_service, destination_service, source_version, destination_version)"

type serviceGraph struct {
	Services    map[string]struct{}
	Connections []Connection
}

type Connection struct {
	Source      string
	Destination string
	RPS         float64
}

var client api.Client

func PromQuery(addr string) (*serviceGraph, error) {
	if client == nil {
		var err error
		client, err = api.NewClient(api.Config{Address: addr})
		if err != nil {
			return nil, err
		}
	}

	promAPI := v1.NewAPI(client)
	query := fmt.Sprintf(reqsFmt, "1m")
	val, err := promAPI.Query(context.Background(), query, time.Now())
	if err != nil {
		return nil, err
	}

	switch v := val.(type) {
	case model.Vector:
		s := serviceGraph{
			Services:    map[string]struct{}{},
			Connections: []Connection{},
		}

		for _, sample := range v {
			svc := sample.Metric[model.LabelName("source_service")]
			dest := sample.Metric[model.LabelName("destination_service")]

			s.Services[string(svc)] = struct{}{}
			s.Services[string(dest)] = struct{}{}

			s.Connections = append(s.Connections, Connection{
				Source:      string(svc),
				Destination: string(dest),
				RPS:         float64(sample.Value),
			})
		}

		return &s, nil
	}

	return nil, fmt.Errorf("unknown value type returned from query: %#v", val)
}
