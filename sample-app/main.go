package main

import (
	"io"
	"log"
	"net/http"
	"runtime"
	"time"
)

func main() {
	http.HandleFunc("/hello-world", handleWorldRequest)
	http.HandleFunc("/hello-universe", handleUniverseRequest)
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
