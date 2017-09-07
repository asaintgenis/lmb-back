package services

import (
	"errors"
	"testing"

	"gitlab.com/locatemybeer/lmb-back/app"
	"gitlab.com/locatemybeer/lmb-back/models"
	"github.com/stretchr/testify/assert"
)

func TestNewBeerStyleService(t *testing.T) {
	dao := newMockBeerStyleDAO()
	s := NewBeerStyleService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestBeerStyleService_Get(t *testing.T) {
	s := NewBeerStyleService(newMockBeerStyleDAO())
	beerStyle, err := s.Get(nil, 1)
	if assert.Nil(t, err) && assert.NotNil(t, beerStyle) {
		assert.Equal(t, "aaa", beerStyle.Name)
	}

	beerStyle, err = s.Get(nil, 100)
	assert.NotNil(t, err)
}

func TestBeerStyleService_Create(t *testing.T) {
	s := NewBeerStyleService(newMockBeerStyleDAO())
	beerStyle, err := s.Create(nil, &models.BeerStyle{
		Name: "ddd",
	})
	if assert.Nil(t, err) && assert.NotNil(t, beerStyle) {
		assert.Equal(t, 4, beerStyle.Id)
		assert.Equal(t, "ddd", beerStyle.Name)
	}

	// dao error
	_, err = s.Create(nil, &models.BeerStyle{
		Id:   100,
		Name: "ddd",
	})
	assert.NotNil(t, err)

	// validation error
	_, err = s.Create(nil, &models.BeerStyle{
		Name: "",
	})
	assert.NotNil(t, err)
}

func TestBeerStyleService_Update(t *testing.T) {
	s := NewBeerStyleService(newMockBeerStyleDAO())
	beerStyle, err := s.Update(nil, 2, &models.BeerStyle{
		Name: "ddd",
	})
	if assert.Nil(t, err) && assert.NotNil(t, beerStyle) {
		assert.Equal(t, 2, beerStyle.Id)
		assert.Equal(t, "ddd", beerStyle.Name)
	}

	// dao error
	_, err = s.Update(nil, 100, &models.BeerStyle{
		Name: "ddd",
	})
	assert.NotNil(t, err)

	// validation error
	_, err = s.Update(nil, 2, &models.BeerStyle{
		Name: "",
	})
	assert.NotNil(t, err)
}

func TestBeerStyleService_Delete(t *testing.T) {
	s := NewBeerStyleService(newMockBeerStyleDAO())
	beerStyle, err := s.Delete(nil, 2)
	if assert.Nil(t, err) && assert.NotNil(t, beerStyle) {
		assert.Equal(t, 2, beerStyle.Id)
		assert.Equal(t, "bbb", beerStyle.Name)
	}

	_, err = s.Delete(nil, 2)
	assert.NotNil(t, err)
}

func TestBeerStyleService_Query(t *testing.T) {
	s := NewBeerStyleService(newMockBeerStyleDAO())
	result, err := s.Query(nil, 1, 2)
	if assert.Nil(t, err) {
		assert.Equal(t, 2, len(result))
	}
}

func newMockBeerStyleDAO() beerStyleDao {
	return &mockBeerStyleDAO{
		records: []models.BeerStyle{
			{Id: 1, Name: "aaa"},
			{Id: 2, Name: "bbb"},
			{Id: 3, Name: "ccc"},
		},
	}
}

type mockBeerStyleDAO struct {
	records []models.BeerStyle
}

func (m *mockBeerStyleDAO) Get(rs app.RequestScope, id int) (*models.BeerStyle, error) {
	for _, record := range m.records {
		if record.Id == id {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *mockBeerStyleDAO) Query(rs app.RequestScope, offset, limit int) ([]models.BeerStyle, error) {
	return m.records[offset : offset+limit], nil
}

func (m *mockBeerStyleDAO) Count(rs app.RequestScope) (int, error) {
	return len(m.records), nil
}

func (m *mockBeerStyleDAO) Create(rs app.RequestScope, beerStyle *models.BeerStyle) error {
	if beerStyle.Id != 0 {
		return errors.New("Id cannot be set")
	}
	beerStyle.Id = len(m.records) + 1
	m.records = append(m.records, *beerStyle)
	return nil
}

func (m *mockBeerStyleDAO) Update(rs app.RequestScope, id int, beerStyle *models.BeerStyle) error {
	beerStyle.Id = id
	for i, record := range m.records {
		if record.Id == id {
			m.records[i] = *beerStyle
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockBeerStyleDAO) Delete(rs app.RequestScope, id int) error {
	for i, record := range m.records {
		if record.Id == id {
			m.records = append(m.records[:i], m.records[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
