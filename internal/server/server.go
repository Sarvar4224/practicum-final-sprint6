package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

// функция создания нового сервера
func NewServer(logger *log.Logger) *Server {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		Server: httpServer,
	}
}

func (s *Server) Start() {
	s.Logger.Println("Сервер запущен на порту 8080")
	err := s.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.Logger.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
