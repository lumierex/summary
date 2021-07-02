package concurrency

import (
	"log"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panicf("http server err %+v", err)
		return
	}

	go pprof()
}

// good case
func pprof() {
	http.ListenAndServe(":8081", nil)
}

// bad case
// go action must be the option for others person
func pprof1() {
	go http.ListenAndServe(":8081", nil)
}
