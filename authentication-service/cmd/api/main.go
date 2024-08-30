package main

import (
	"authentication-service/cmd/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const webPort = "8002"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service")

	// TODO connect to DB

	app := Config{}

	srv := &http.Server{
		Addr:    fmt.Sprint(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
