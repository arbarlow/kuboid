package kuboid

import (
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/prometheus/client_golang/api"
	"github.com/stretchr/testify/assert"
)

func TestPromQuery(t *testing.T) {
	defer setRecorderClient(t, "fixtures/test_prom_query").Stop()

	res, err := PromQuery("http://localhost:9090")
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, res)
}

func setRecorderClient(t *testing.T, path string) *recorder.Recorder {
	r, err := recorder.New(path)
	if err != nil {
		t.Fatal(err)
	}

	r.SetMatcher(func(r *http.Request, i cassette.Request) bool {
		return r.Method == i.Method && r.URL.Path == "/api/v1/query"
	})

	client, err = api.NewClient(api.Config{
		Address:      "http://localhost:9090",
		RoundTripper: r,
	})
	if err != nil {
		t.Fatal(err)
	}

	return r
}
