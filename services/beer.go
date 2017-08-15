package services

import (
	"gitlab.com/locatemybeer/lmb-back/app"
	"gitlab.com/locatemybeer/lmb-back/models"
)

// artistDAO specifies the interface of the beer DAO needed by BeerService.
type beerDao interface {
	// Get returns the beer with the specified beer ID.
	Get(rs app.RequestScope, id int) (*models.Beer, error)
	// Count returns the number of beers.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of beers with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.Beer, error)
	// Create saves a new beer in the storage.
	Create(rs app.RequestScope, artist *models.Beer) error
	// Update updates the beer with given ID in the storage.
	Update(rs app.RequestScope, id int, artist *models.Beer) error
	// Delete removes the beer with given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

// BeerService provides services related with beers.
type BeerService struct {
	dao beerDao
}

// NewArtistService creates a new BeerService with the given artist DAO.
func NewBeerService(dao beerDao) *BeerService {
	return &BeerService{dao}
}

// Get returns the beer with the specified the artist ID.
func (s *BeerService) Get(rs app.RequestScope, id int) (*models.Beer, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new Beer.
func (s *BeerService) Create(rs app.RequestScope, model *models.Beer) (*models.Beer, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

// Update updates the beer with the specified ID.
func (s *BeerService) Update(rs app.RequestScope, id int, model *models.Beer) (*models.Beer, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the beer with the specified ID.
func (s *BeerService) Delete(rs app.RequestScope, id int) (*models.Beer, error) {
	artist, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return artist, err
}

// Count returns the number of beers.
func (s *BeerService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the beers with the specified offset and limit.
func (s *BeerService) Query(rs app.RequestScope, offset, limit int) ([]models.Beer, error) {
	return s.dao.Query(rs, offset, limit)
}
