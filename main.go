package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/v1", v1)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func v1(w http.ResponseWriter, r *http.Request) {
	response := "hello world"
	w.Write([]byte(response))
}
