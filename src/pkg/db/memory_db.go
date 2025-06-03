package db

import (
	"errors"
	"mini_service_citatnik/src/internal/models"
	"sync"
	"sync/atomic"
)

var (
	ErrNotFound = errors.New("not found")
)

type memoryDb struct {
	mu            sync.RWMutex
	quotes        []*models.Quote
	indexByAuthor map[string][]*models.Quote
	counter       int64
}

func NewMemoryDb() IMemoryDb {
	return &memoryDb{
		indexByAuthor: make(map[string][]*models.Quote),
	}
}

func (m *memoryDb) InsertQuote(quote *models.Quote) {
	m.mu.Lock()
	defer m.mu.Unlock()

	quote.Id = atomic.AddInt64(&m.counter, 1)
	m.quotes = append(m.quotes, quote)
	m.indexByAuthor[quote.Author] = append(m.indexByAuthor[quote.Author], quote)
}

func (m *memoryDb) DeleteQuote(id int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, quote := range m.quotes {
		if quote.Id == id {
			m.quotes = append(m.quotes[:i], m.quotes[i+1:]...)

			author := quote.Author
			quotes := m.indexByAuthor[author]
			for j, q := range quotes {
				if q.Id == id {
					m.indexByAuthor[author] = append(quotes[:j], quotes[j+1:]...)
					break
				}
			}
			return nil
		}
	}
	return ErrNotFound
}

func (m *memoryDb) GetByAuthor(authorName string) []*models.Quote {
	m.mu.RLock()
	defer m.mu.RUnlock()

	quotes := m.indexByAuthor[authorName]
	result := make([]*models.Quote, len(quotes))
	for i, quote := range quotes {
		copy := *quote
		result[i] = &copy
	}
	return result
}

func (m *memoryDb) Data() []*models.Quote {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]*models.Quote, len(m.quotes))
	for i, quote := range m.quotes {
		copy := *quote
		result[i] = &copy
	}
	return result
}

func (m *memoryDb) Conn() (interface{}, error) {
	return m, nil
}

func (m *memoryDb) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.counter = 0
	m.quotes = nil

	return nil
}
