package main

import (
	"github.com/smakaroni/epulse-cats/pkg/app"
	"github.com/smakaroni/epulse-cats/pkg/config"
	"log"
)

func main() {

	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("failed to load config:", err)
	}

	a := app.App{}
	a.Init(&c)

	a.Run(":8080")

}
