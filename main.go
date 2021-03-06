package main

import (
	"embed"
	_ "embed"
	"log"
	"net/http"
)

//go:embed webpage/*
var webpage embed.FS

func main() {
	wsServer := NewServer()
	go wsServer.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		Serve(wsServer, w, r)
	})
	http.Handle("/", http.FileServer(http.FS(webpage)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
