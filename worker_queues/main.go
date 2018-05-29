package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	NWorkers = flag.Int("n", 4, "The number of workers to start")
	HTTPAddr = flag.String("http", "127.0.0.1:8000", "Address to listen for HTTP requests on")
)

func main() {
	flag.Parse()

	fmt.Println("Starting the dispatcher")
	StartDispatcher(*NWorkers)

	fmt.Println("Registering the web collector")
	http.HandleFunc("/work", WebCollect)

	fmt.Println("Starting the Rabbit collector")
	RabbitCollect()

	fmt.Println("HTTP server listening on", *HTTPAddr)
	if err := http.ListenAndServe(*HTTPAddr, nil); err != nil {
		fmt.Println(err.Error())
	}
}
