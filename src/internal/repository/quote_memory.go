package repository

import (
	"context"
	"mini_service_citatnik/src/internal/models"
	"mini_service_citatnik/src/internal/service"
	"mini_service_citatnik/src/pkg/db"
)

type QuoteMemoryRepository struct {
	memoryDb db.IMemoryDb
}

func NewQuoteRepository(memoryDb db.IMemoryDb) (service.IQuoteRepository, error) {
	var conn, err = memoryDb.Conn()
	if err != nil {
		return nil, err
	}
	mem := conn.(db.IMemoryDb)

	return &QuoteMemoryRepository{memoryDb: mem}, nil
}

func (r *QuoteMemoryRepository) Insert(ctx context.Context, quote *models.Quote) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	r.memoryDb.InsertQuote(quote)
	return nil
}

func (r *QuoteMemoryRepository) GetAll(ctx context.Context) ([]*models.Quote, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	return r.memoryDb.Data(), nil
}

func (r *QuoteMemoryRepository) GetById(ctx context.Context, id int64) (*models.Quote, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	quotes := r.memoryDb.Data()
	for _, quote := range quotes {
		if quote.Id == id {
			return quote, nil
		}
	}

	return nil, db.ErrNotFound
}

func (r *QuoteMemoryRepository) GetByAuthor(ctx context.Context, authorName string) ([]*models.Quote, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	result := r.memoryDb.GetByAuthor(authorName)
	if len(result) == 0 {
		return nil, db.ErrNotFound
	}

	return result, nil
}

func (r *QuoteMemoryRepository) Delete(ctx context.Context, id int64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return r.memoryDb.DeleteQuote(id)
}
