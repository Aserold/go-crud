package main

import (
	"log"
	"os"

	"github.com/Aserold/go-crud/config"
	"github.com/Aserold/go-crud/internal/server"
	"github.com/Aserold/go-crud/pkg/db/postgres"
	"github.com/Aserold/go-crud/pkg/utils"
)

func main() {
	log.Println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	log.Println("Config loaded")

	psqlDB, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		log.Fatalf("Postgresql init: %s", err)
	} else {
		log.Printf("Postgres connected, Status: %#v", psqlDB.Stats())
	}
	defer psqlDB.Close()

	s := server.NewServer(cfg, psqlDB)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
