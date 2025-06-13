package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
);

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer cancel()

	mux := http.NewServeMux();
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request, ) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	go func() {
		log.Printf("HTTP server listening on 127.0.0.1:8001")
		if err := http.ListenAndServe("127.0.0.1:8001", mux); err != nil {
			log.Fatalf("Could not start the HTTP server on 127.0.0.1:8001: %v", err)
		}
	}()

	<-ctx.Done()
}