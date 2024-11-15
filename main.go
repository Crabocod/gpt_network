package main

import (
	"fmt"
	"log"
	"net/http"

	"web.app/internal/config"
	"web.app/internal/db"
	"web.app/internal/handlers"
	"web.app/internal/middlewares"

	gorillaHandlers "github.com/gorilla/handlers"
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
	r.HandleFunc("/auth/registration/", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/auth/login/", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/auth/refresh/", handlers.RefreshTokenHandler).Methods("POST")
	r.Handle("/auth/logout/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.LogoutHandler))).Methods("POST")

	r.Handle("/users/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.GetUserHandler))).Methods("GET")

	r.Handle("/posts/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.GetPostsHandler))).Methods("GET")
	r.Handle("/posts/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.CreatePostHandler))).Methods("POST")
	r.Handle("/posts/{id}/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.UpdatePostHandler))).Methods("PUT")
	r.Handle("/posts/{id}/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.DeletePostHandler))).Methods("DELETE")

	r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/docs/swagger.yaml"),
	))

	originsOk := gorillaHandlers.AllowedOrigins([]string{"http://localhost:8080"})
	methodsOk := gorillaHandlers.AllowedMethods([]string{"HEAD", "GET", "POST", "PUT", "DELETE", "OPTIONS"})
	headersOk := gorillaHandlers.AllowedHeaders([]string{"Origin", "Content-Type", "Authorization", "X-Requested-With", "access-control-expose-headers"})
	handlerWithCORS := gorillaHandlers.CORS(originsOk, headersOk, methodsOk)(r)

	// Запуск сервера
	fmt.Println("Starting server on :85...")
	if err := http.ListenAndServe(":85", handlerWithCORS); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
