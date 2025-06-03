package db

import (
	"mini_service_citatnik/src/internal/models"
)

type IMemoryDb interface {
	DataBase

	InsertQuote(quote *models.Quote)
	DeleteQuote(id int64) error
	GetByAuthor(authorName string) []*models.Quote
	Data() []*models.Quote
}
