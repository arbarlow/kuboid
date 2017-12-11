package kuboid

import (
	"fmt"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/prometheus/client_golang/api"
)

func TestPromQuery(t *testing.T) {
	r, err := recorder.New("fixtures/prometheus")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()

	client, err = api.NewClient(api.Config{
		Address:      "http://localhost:9090",
		RoundTripper: r,
	})
	if err != nil {
		t.Fatal(err)
	}

	res, err := PromQuery("http://localhost:9090")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("res = %+v\n", res)
}
