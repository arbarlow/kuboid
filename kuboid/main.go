package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/arbarlow/kuboid"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	prometheusAddr = kingpin.Flag("prometheus-addr", "addr of prometheus server").Required().String()
)

func main() {
	kingpin.Parse()

	res, err := kuboid.PromQuery(*prometheusAddr)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}
