package service

import (
	"context"
	"mini_service_citatnik/src/internal/models"
)

type IQuoteRepository interface {
	Insert(ctx context.Context, quote *models.Quote) error
	GetAll(ctx context.Context) ([]*models.Quote, error)
	GetById(ctx context.Context, id int64) (*models.Quote, error)
	GetByAuthor(ctx context.Context, authorName string) ([]*models.Quote, error)
	Delete(ctx context.Context, id int64) error
}

type QuoteService struct {
	repository IQuoteRepository
}

func NewQuoteService(repository IQuoteRepository) *QuoteService {
	return &QuoteService{repository: repository}
}

func (q *QuoteService) Insert(ctx context.Context, quote *models.Quote) error {
	err := q.repository.Insert(ctx, quote)
	if err != nil {
		return err
	}

	return err
}

func (q *QuoteService) GetAll(ctx context.Context) ([]*models.Quote, error) {
	quotes, err := q.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

func (q *QuoteService) GetById(ctx context.Context, id int64) (*models.Quote, error) {
	quote, err := q.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func (q *QuoteService) GetByAuthor(ctx context.Context, authorName string) ([]*models.Quote, error) {
	quotes, err := q.repository.GetByAuthor(ctx, authorName)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

func (q *QuoteService) Delete(ctx context.Context, id int64) error {
	err := q.repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
