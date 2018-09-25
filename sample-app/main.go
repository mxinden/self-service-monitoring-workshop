package main

import (
	"io"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "sample_app_api_requests_total",
			Help: "A counter for requests to the wrapped handler.",
		},
		[]string{"code", "method"},
	)

	// TODO: Don't use global registry
	prometheus.MustRegister(counter)

	worldInstrumented := promhttp.InstrumentHandlerCounter(counter, http.HandlerFunc(handleWorldRequest))
	universeIntrumented := promhttp.InstrumentHandlerCounter(counter, http.HandlerFunc(handleUniverseRequest))

	http.HandleFunc("/hello-world", worldInstrumented)
	http.HandleFunc("/hello-universe", universeIntrumented)
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWorldRequest(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func handleUniverseRequest(w http.ResponseWriter, req *http.Request) {
	wasteCPUCycles()

	io.WriteString(w, "Hello, universe!\n")
}

// wasteCPUCycles generates arbitrary CPU heavy load. Taken from
// https://stackoverflow.com/a/41084841 .
func wasteCPUCycles() {
	done := make(chan int)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}
	time.Sleep(100 * time.Millisecond)
	close(done)
}
