package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/grafana/pyroscope-go"
)

func main() {
	//// These 2 lines are only required if you're using mutex or block profiling
	//// Read the explanation below for how to set these rates:
	//runtime.SetMutexProfileFraction(5)
	//runtime.SetBlockProfileRate(5)

	// setup pyroscope
	pyroscope.Start(pyroscope.Config{
		ApplicationName: "1brc.golang.app",
		ServerAddress:   "http://localhost:4040",
		Logger:          pyroscope.StandardLogger,
		Tags: map[string]string{
			"hostname": os.Getenv("HOSTNAME"),
		},
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})

	// server http
	http.HandleFunc("/v1", v1)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func v1(w http.ResponseWriter, r *http.Request) {
	response := "hello world"
	w.Write([]byte(response))
}
