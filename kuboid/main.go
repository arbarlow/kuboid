package main

import (
	"log"
	"net/http"

	"github.com/arbarlow/kuboid"
	"github.com/neelance/graphql-go/relay"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	prometheusAddr = kingpin.Flag("prometheus-addr", "addr of prometheus server").Required().String()
)

func main() {
	kingpin.Parse()

	// res, err := kuboid.PromQuery(*prometheusAddr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	http.HandleFunc("/", graphiQLHandler())

	http.Handle("/query", &relay.Handler{Schema: kuboid.Schema})

	log.Print("Starting GraphQL server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
