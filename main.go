package main

import (
	log "github.com/sirupsen/logrus"
	"value-survey/api"
	"value-survey/internal/config"
	"value-survey/internal/handler"
	"value-survey/internal/persistence"
	"value-survey/internal/server"
	"value-survey/pkg/storage"
)

func main() {

	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	if err := handler.InitHandler(); err != nil {
		log.Fatal(err)
	}

	s := server.NewHttpServer()

	if err := storage.InitDefault(config.GetConfig().Persistence); err != nil {
		log.Fatal(err)
	}

	if err := persistence.InitRDB(storage.GetDefault()); err != nil {
		log.Fatal(err)
	}

	if err := api.RegisterRouter(s); err != nil {
		log.Fatal(err)
	}

	if err := s.Run(":8080"); err != nil {
		log.Fatal(err)
	}

	return
}
