package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/Crabocod/gpt_network/api-service/internal/app/apiserver"
	"github.com/Crabocod/gpt_network/api-service/internal/app/grpcserver"
	"github.com/Crabocod/gpt_network/api-service/internal/config"
	"github.com/Crabocod/gpt_network/api-service/internal/db"
	"log"
	"sync"
)

var configPath = flag.String("config-path", "./configs/apiserver.toml", "configuration path")

func main() {
	flag.Parse()

	conf := config.NewConfig()
	_, err := toml.DecodeFile(*configPath, conf)
	if err != nil {
		log.Fatal(err)
	}

	pg, err := db.Connect(conf)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		s := apiserver.New(pg, conf)
		if err = s.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()

		s := grpcserver.New(pg, conf)
		if err = s.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}
