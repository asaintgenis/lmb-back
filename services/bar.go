package services

import (
	"gitlab.com/locatemybeer/lmb-back/app"
	"gitlab.com/locatemybeer/lmb-back/models"
)

// barDao specifies the interface of the bar DAO needed by BarService.
type barDao interface {
	// Get returns the bar with the specified bar ID.
	Get(rs app.RequestScope, id int) (*models.Bar, error)
	// Count returns the number of bars.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of bars with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.Bar, error)
	// Create saves a new bar in the storage.
	Create(rs app.RequestScope, bar *models.Bar) error
	// Update updates the bar with given ID in the storage.
	Update(rs app.RequestScope, id int, bar *models.Bar) error
	// Delete removes the bar with given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

// BarService provides services related with bars.
type BarService struct {
	dao barDao
}

// NewBarService creates a new BarService with the given bar DAO.
func NewBarService(dao barDao) *BarService {
	return &BarService{dao}
}

// Get returns the bar with the specified the bar ID.
func (s *BarService) Get(rs app.RequestScope, id int) (*models.Bar, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new Bar.
func (s *BarService) Create(rs app.RequestScope, model *models.Bar) (*models.Bar, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

// Update updates the bar with the specified ID.
func (s *BarService) Update(rs app.RequestScope, id int, model *models.Bar) (*models.Bar, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the bar with the specified ID.
func (s *BarService) Delete(rs app.RequestScope, id int) (*models.Bar, error) {
	bar, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return bar, err
}

// Count returns the number of bars.
func (s *BarService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the bars with the specified offset and limit.
func (s *BarService) Query(rs app.RequestScope, offset, limit int) ([]models.Bar, error) {
	return s.dao.Query(rs, offset, limit)
}
