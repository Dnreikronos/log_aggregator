package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request from %s", r.RemoteAddr)
		w.Write([]byte("Hello, World!"))
	})

	go func() {
		for {
			log.Printf("Application is running at %s", time.Now().Format(time.RFC3339))
			time.Sleep(10 * time.Second)
		}
	}()

	log.Println("Starting server on :9090")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}
