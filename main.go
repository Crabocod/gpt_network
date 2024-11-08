package main

import (
	"fmt"
	"log"
	"net/http"

	"web.app/internal/config"
	"web.app/internal/db"
	"web.app/internal/handlers"
	"web.app/internal/sessions"

	"github.com/gorilla/mux"
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
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("GET", "POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET")
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	// Настройка сессий
	sessions.InitSessionStore()

	// Запуск сервера
	fmt.Println("Starting server on :85...")
	if err := http.ListenAndServe(":85", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
