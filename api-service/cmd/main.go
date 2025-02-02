package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/Crabocod/gpt_network/api-service/internal/config"
	"github.com/Crabocod/gpt_network/api-service/internal/db"
	"github.com/Crabocod/gpt_network/api-service/internal/handlers"
	"github.com/Crabocod/gpt_network/api-service/internal/logger"
	"github.com/Crabocod/gpt_network/api-service/internal/middlewares"
	pb "github.com/Crabocod/gpt_network/api-service/internal/proto"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
)

var configPath = flag.String("config-path", "./config.toml", "configuration path")

func main() {
	flag.Parse()
	if err := config.LoadConfig(*configPath); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	if err := logger.LoadLogger(); err != nil {
		log.Fatalf("Error loading logrus: %v", err)
	}

	err := db.Connect()
	if err != nil {
		logger.Logrus.Fatalf("Error connecting to database: %v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/auth/registration/", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/auth/login/", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/auth/refresh/", handlers.RefreshTokenHandler).Methods("POST")
	r.Handle("/auth/logout/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.LogoutHandler))).Methods("POST")

	r.Handle("/users/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.GetUserHandler))).Methods("GET")

	r.Handle("/posts/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.GetPostsHandler))).Methods("GET")
	r.Handle("/posts/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.CreatePostHandler))).Methods("POST")
	r.Handle("/posts/{id}/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.UpdatePostHandler))).Methods("PUT")
	r.Handle("/posts/{id}/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.DeletePostHandler))).Methods("DELETE")

	r.Handle("/posts/{post_id}/comments/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.GetCommentsHandler))).Methods("GET")
	r.Handle("/posts/{post_id}/comments/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.CreateCommentHandler))).Methods("POST")
	r.Handle("/comments/{id}/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.UpdateCommentHandler))).Methods("PUT")
	r.Handle("/comments/{id}/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.DeleteCommentHandler))).Methods("DELETE")

	r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/docs/swagger.yaml"),
	))

	originsOk := gorillaHandlers.AllowedOrigins([]string{"http://localhost:8080"})
	methodsOk := gorillaHandlers.AllowedMethods([]string{"HEAD", "GET", "POST", "PUT", "DELETE", "OPTIONS"})
	headersOk := gorillaHandlers.AllowedHeaders([]string{"Origin", "Content-Type", "Authorization", "X-Requested-With", "access-control-expose-headers"})
	handlerWithCORS := gorillaHandlers.CORS(originsOk, headersOk, methodsOk)(r)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		logger.Logrus.Info("Starting HTTP server on :85...")
		if err := http.ListenAndServe(":85", handlerWithCORS); err != nil {
			logger.Logrus.Fatalf("HTTP server failed to start: %v", err)
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
