package main

import (
	"fmt"
	"log"
	"net/http"

	"web.app/internal/config"
	"web.app/internal/db"
	"web.app/internal/handlers"
	"web.app/internal/middlewares"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// Загрузка конфигурации
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Подключение к базе данных
	err = db.Connect()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Настройка маршрутов
	r := mux.NewRouter()
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/refresh-token", handlers.RefreshTokenHandler).Methods("POST")

	r.Handle("/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.HomeHandler))).Methods("GET")
	r.Handle("/logout", middlewares.AuthMiddleware(http.HandlerFunc(handlers.LogoutHandler))).Methods("POST")

	r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/docs/swagger.yaml"),
	))

	// Запуск сервера
	fmt.Println("Starting server on :85...")
	if err := http.ListenAndServe(":85", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
