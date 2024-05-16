package main

import (
	"rapikan/configs"
	"rapikan/routes"
	"os"
)

func main() {
	configs.InitDB()
	e := routes.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start("0.0.0.0:" + port))
}
