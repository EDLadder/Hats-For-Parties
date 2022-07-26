package main

import (
	"log"
	"net/http"

	"github.com/EDLadder/Hats-For-Parties/config"
	"github.com/EDLadder/Hats-For-Parties/logs"
	"github.com/EDLadder/Hats-For-Parties/routes"
	"github.com/fatih/color"
	"github.com/rs/cors"
)

func main() {
	port := config.GetEnvVariable("PORT")
	color.Cyan("🌏 Server running on localhost:" + port)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	router := routes.Routes()

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PATCH"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	handler := c.Handler(router)
	http.ListenAndServe(":"+port, logs.LogRequest(handler))
}
