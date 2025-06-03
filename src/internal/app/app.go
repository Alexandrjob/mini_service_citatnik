package app

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"mini_service_citatnik/src/internal/handler"
	"mini_service_citatnik/src/internal/models"
	"mini_service_citatnik/src/internal/repository"
	"mini_service_citatnik/src/internal/service"
	"mini_service_citatnik/src/pkg/db"
	"net/http"
)

func SetupRouter(quoteService *service.QuoteService) *mux.Router {
	r := mux.NewRouter()
	quoteHandler := handler.NewQuoteHandler(quoteService)

	r.HandleFunc("/quotes", quoteHandler.GetByAuthor).Methods("GET").Queries("author", "{author}")

	r.HandleFunc("/quotes", quoteHandler.GetAllQuotes).Methods("GET")

	r.HandleFunc("/quotes", quoteHandler.CreateQuote).Methods("POST")
	r.HandleFunc("/quotes/random", quoteHandler.GetRandomQuote).Methods("GET")
	r.HandleFunc("/quotes/{id}", quoteHandler.DeleteQuote).Methods("DELETE")

	return r
}

func RunServiceInstance() {
	memoryDb := db.NewMemoryDb()
	quoteRepository, _ := repository.NewQuoteRepository(memoryDb)
	quoteService := service.NewQuoteService(quoteRepository)

	// Добавление тестовых данных
	quoteService.Insert(context.Background(),
		&models.Quote{Author: "Confucius", Text: "Life is simple, but we insist on making it complicated."})
	quoteService.Insert(context.Background(),
		&models.Quote{Author: "Einstein", Text: "Imagination is more important than knowledge."})

	// Запуск сервера
	r := SetupRouter(quoteService)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
