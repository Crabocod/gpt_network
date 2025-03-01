package apiserver

import (
	"context"
	"errors"
	"github.com/Crabocod/gpt_network/api-service/internal/app/controller/rest"
	"github.com/Crabocod/gpt_network/api-service/internal/app/service"
	"github.com/Crabocod/gpt_network/api-service/internal/app/store/postgresql"
	"github.com/Crabocod/gpt_network/api-service/internal/config"
	"github.com/Crabocod/gpt_network/api-service/internal/middlewares"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"time"
)

type APIServer struct {
	logger *logrus.Logger
	config *config.Config
	router *mux.Router
	db     *sqlx.DB
	server *http.Server
}

func New(db *sqlx.DB, config *config.Config) *APIServer {
	s := &APIServer{
		logger: logrus.New(),
		router: mux.NewRouter(),
		config: config,
		db:     db,
	}

	if err := s.configureLogger(); err != nil {
		s.logger.Fatal(err)
	}

	s.configureRouter()

	s.server = &http.Server{
		Addr:    s.config.ApiServer.BindAddr,
		Handler: getCORS(s.router),
	}

	return s
}

func (s *APIServer) Run(ctx context.Context) error {
	s.logger.Info("Starting HTTP server on ", s.config.ApiServer.BindAddr)

	go func() {
		<-ctx.Done()
		s.logger.Info("Shutting down HTTP server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := s.server.Shutdown(shutdownCtx); err != nil {
			s.logger.Error("HTTP server shutdown error: ", err)
		}
	}()

	if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *APIServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *APIServer) configureRouter() {
	store := postgresql.New(s.db)
	business := service.NewService(store)
	controller := rest.NewController(*business)

	s.router.HandleFunc("/auth/registration/", controller.UserController.RegisterHandler).Methods("POST")
	s.router.HandleFunc("/auth/login/", controller.UserController.LoginHandler).Methods("POST")
	s.router.HandleFunc("/auth/refresh/", controller.UserController.RefreshTokenHandler).Methods("POST")
	s.router.Handle("/auth/logout/", middlewares.AuthMiddleware(http.HandlerFunc(controller.UserController.LogoutHandler))).Methods("POST")
	s.router.Handle("/users/", middlewares.AuthMiddleware(http.HandlerFunc(controller.UserController.GetUserHandler))).Methods("GET")

	s.router.Handle("/posts/", middlewares.AuthMiddleware(http.HandlerFunc(controller.PostController.GetPostsHandler))).Methods("GET")
	s.router.Handle("/posts/", middlewares.AuthMiddleware(http.HandlerFunc(controller.PostController.CreatePostHandler))).Methods("POST")
	s.router.Handle("/posts/{id}/", middlewares.AuthMiddleware(http.HandlerFunc(controller.PostController.UpdatePostHandler))).Methods("PUT")
	s.router.Handle("/posts/{id}/", middlewares.AuthMiddleware(http.HandlerFunc(controller.PostController.DeletePostHandler))).Methods("DELETE")

	s.router.Handle("/posts/{post_id}/comments/", middlewares.AuthMiddleware(http.HandlerFunc(controller.CommentController.GetCommentsHandler))).Methods("GET")
	s.router.Handle("/posts/{post_id}/comments/", middlewares.AuthMiddleware(http.HandlerFunc(controller.CommentController.CreateCommentHandler))).Methods("POST")
	s.router.Handle("/comments/{id}/", middlewares.AuthMiddleware(http.HandlerFunc(controller.CommentController.UpdateCommentHandler))).Methods("PUT")
	s.router.Handle("/comments/{id}/", middlewares.AuthMiddleware(http.HandlerFunc(controller.CommentController.DeleteCommentHandler))).Methods("DELETE")

	s.router.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))
	s.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/docs/swagger.yaml"),
	))
}

func (s *APIServer) configureLogger() error {
	Formatter := new(logrus.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	Formatter.ForceColors = true
	s.logger.SetFormatter(Formatter)

	level, err := logrus.ParseLevel(s.config.Logger.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func getCORS(r *mux.Router) http.Handler {
	originsOk := gorillaHandlers.AllowedOrigins([]string{"http://localhost:8080"})
	methodsOk := gorillaHandlers.AllowedMethods([]string{"HEAD", "GET", "POST", "PUT", "DELETE", "OPTIONS"})
	headersOk := gorillaHandlers.AllowedHeaders([]string{"Origin", "Content-Type", "Authorization", "X-Requested-With", "access-control-expose-headers"})

	return gorillaHandlers.CORS(originsOk, headersOk, methodsOk)(r)
}
