package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	srv := server.NewServer(log.Default())
	err := srv.Server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
