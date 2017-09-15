package services

import (
	"gitlab.com/locatemybeer/lmb-back/app"
	"gitlab.com/locatemybeer/lmb-back/models"
)

// beerStyleDao specifies the interface of the beerStyle DAO needed by BeerStyleService.
type beerStyleDao interface {
	// Get returns the beerStyle with the specified beerStyle ID.
	Get(rs app.RequestScope, id uint) (*models.BeerStyle, error)
	// Count returns the number of beerStyles.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of beerStyles with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.BeerStyle, error)
	// Create saves a new beerStyle in the storage.
	Create(rs app.RequestScope, beerStyle *models.BeerStyle) error
	// Update updates the beerStyle with given ID in the storage.
	Update(rs app.RequestScope, beerStyle *models.BeerStyle) error
	// Delete removes the beerStyle with given ID from the storage.
	Delete(rs app.RequestScope, id uint) error
}

// BeerStyleService provides services related with beerStyles.
type BeerStyleService struct {
	dao beerStyleDao
}

// NewBeerStyleService creates a new BeerStyleService with the given beerStyle DAO.
func NewBeerStyleService(dao beerStyleDao) *BeerStyleService {
	return &BeerStyleService{dao}
}

// Get returns the beerStyle with the specified the beerStyle ID.
func (s *BeerStyleService) Get(rs app.RequestScope, id uint) (*models.BeerStyle, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new BeerStyle.
func (s *BeerStyleService) Create(rs app.RequestScope, model *models.BeerStyle) (*models.BeerStyle, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.ID)
}

// Update updates the beerStyle with the specified ID.
func (s *BeerStyleService) Update(rs app.RequestScope, id uint, model *models.BeerStyle) (*models.BeerStyle, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the beerStyle with the specified ID.
func (s *BeerStyleService) Delete(rs app.RequestScope, id uint) (*models.BeerStyle, error) {
	beerStyle, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return beerStyle, err
}

// Count returns the number of beerStyles.
func (s *BeerStyleService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the beerStyles with the specified offset and limit.
func (s *BeerStyleService) Query(rs app.RequestScope, offset, limit int) ([]models.BeerStyle, error) {
	return s.dao.Query(rs, offset, limit)
}
