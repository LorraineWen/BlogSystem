package main

import (
	"blogsystem/common"
	"blogsystem/internal/router"
	"log"
	"net/http"
)

func init() {
	common.Load()
}
func main() {
	router.Routers()
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
