package main

import (
	log "github.com/sirupsen/logrus"
	"value-survey/api"
	"value-survey/internal/server"
)

func main() {

	s := server.NewHttpServer()

	if err := api.RegisterRouter(s); err != nil {
		log.Fatal(err)
	}

	if err := s.Run(":8080"); err != nil {
		log.Fatal(err)
	}

	return
}
