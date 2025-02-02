package main

import (
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
func main() {
	blogServer := http.Server{Addr: "localhost:8080"}
	http.HandleFunc("/", indexHandler)
	if err := blogServer.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
