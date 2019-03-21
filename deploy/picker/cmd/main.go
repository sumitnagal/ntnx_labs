package main

import (
	"log"
	"net/http"

	"github.com/gopherland/ntnx_labs/deploy/picker/internal/web"
)

const httpPort = ":4500"

var (
	// Version set via build
	version = "dev"
)

func main() {
	http.HandleFunc("/api/v1/load", web.LoadHandler)
	http.HandleFunc("/api/v1/picker", web.PickHandler)

	log.Printf("Picker[%s] is listening on port [%s]", version, httpPort)
	log.Fatal(http.ListenAndServe(httpPort, nil))
}
