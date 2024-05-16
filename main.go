package main

import (
	"rapikan/configs"
	"rapikan/routes"
)

func main() {
	configs.InitDB()
	e := routes.New()

	e.Start(":8080")
}
