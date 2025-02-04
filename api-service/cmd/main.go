package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/Crabocod/gpt_network/api-service/internal/db"
	"github.com/Crabocod/gpt_network/api-service/internal/handlers"
	"github.com/Crabocod/gpt_network/api-service/internal/logger"
	"github.com/Crabocod/gpt_network/api-service/internal/pkg/app"
	pb "github.com/Crabocod/gpt_network/api-service/internal/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

var configPath = flag.String("config-path", "./config.toml", "configuration path")

func main() {
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		config := app.NewConfig()
		_, err = toml.DecodeFile(*configPath, config)
		if err != nil {
			log.Fatal(err)
		}

		s := app.New(config)
		if err = s.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()
		server := grpc.NewServer()
		pb.RegisterApiServiceServer(server, &handlers.ApiService{})

		listener, err := net.Listen("tcp", ":50052")
		if err != nil {
			logger.Logrus.Fatalf("Error starting gRPC listener: %v", err)
		}

		logger.Logrus.Info("Starting gRPC server on 50052...")
		if err := server.Serve(listener); err != nil {
			logger.Logrus.Fatalf("Error starting gRPC server: %v", err)
		}
	}()

	wg.Wait()
}
