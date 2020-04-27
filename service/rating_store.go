package service

import "sync"

//RatingStore is an interface to store laptop ratings
type RatingStore interface {
	// Add adds a new laptop score to the store and returns its rating
	Add(laptopID string, score float64) (*Rating, error)
}

// Rating contains the rating information of a laptop
type Rating struct {
	Count uint32
	Sun   float64
}

// InMemoryRatingStore stores laptop ratings im memory
type InMemoryRatingStore struct {
	mutex  sync.Mutex
	rating map[string]*Rating
}

// NewInMemoryRatingStore return a new InMemoryRatingStore
// Inicializa la tienda
func NewInMemoryRatingStore() *InMemoryRatingStore {
	return &InMemoryRatingStore{
		rating: make(map[string]*Rating),
	}
}

// Add adds a new laptop score to the and returns its rating
func (store *InMemoryRatingStore) Add(laptopID string, score float64) (*Rating, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	rating := store.rating[laptopID]
	if rating == nil {
		rating = &Rating{
			Count: 1,
			Sun:   score,
		}
	} else {
		rating.Count++
		rating.Sun += score
	}

	store.rating[laptopID] = rating
	return rating, nil
}
