package services

import (
	"errors"
	"testing"

	"gitlab.com/locatemybeer/lmb-back/app"
	"gitlab.com/locatemybeer/lmb-back/models"
	"github.com/stretchr/testify/assert"
)

func TestNewBeerService(t *testing.T) {
	dao := newMockBeerDAO()
	s := NewBeerService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestBeerService_Get(t *testing.T) {
	s := NewBeerService(newMockBeerDAO())
	beer, err := s.Get(nil, 1)
	if assert.Nil(t, err) && assert.NotNil(t, beer) {
		assert.Equal(t, "aaa", beer.Name)
	}

	beer, err = s.Get(nil, 100)
	assert.NotNil(t, err)
}

func TestBeerService_Create(t *testing.T) {
	s := NewBeerService(newMockBeerDAO())
	beer, err := s.Create(nil, &models.Beer{
		Name: "ddd",
	})
	if assert.Nil(t, err) && assert.NotNil(t, beer) {
		assert.Equal(t, 4, beer.Id)
		assert.Equal(t, "ddd", beer.Name)
	}

	// dao error
	_, err = s.Create(nil, &models.Beer{
		Id:   100,
		Name: "ddd",
	})
	assert.NotNil(t, err)

	// validation error
	_, err = s.Create(nil, &models.Beer{
		Name: "",
	})
	assert.NotNil(t, err)
}

func TestBeerService_Update(t *testing.T) {
	s := NewBeerService(newMockBeerDAO())
	beer, err := s.Update(nil, 2, &models.Beer{
		Name: "ddd",
	})
	if assert.Nil(t, err) && assert.NotNil(t, beer) {
		assert.Equal(t, 2, beer.Id)
		assert.Equal(t, "ddd", beer.Name)
	}

	// dao error
	_, err = s.Update(nil, 100, &models.Beer{
		Name: "ddd",
	})
	assert.NotNil(t, err)

	// validation error
	_, err = s.Update(nil, 2, &models.Beer{
		Name: "",
	})
	assert.NotNil(t, err)
}

func TestBeerService_Delete(t *testing.T) {
	s := NewBeerService(newMockBeerDAO())
	beer, err := s.Delete(nil, 2)
	if assert.Nil(t, err) && assert.NotNil(t, beer) {
		assert.Equal(t, 2, beer.Id)
		assert.Equal(t, "bbb", beer.Name)
	}

	_, err = s.Delete(nil, 2)
	assert.NotNil(t, err)
}

func TestBeerService_Query(t *testing.T) {
	s := NewBeerService(newMockBeerDAO())
	result, err := s.Query(nil, 1, 2)
	if assert.Nil(t, err) {
		assert.Equal(t, 2, len(result))
	}
}

func newMockBeerDAO() beerDao {
	return &mockBeerDAO{
		records: []models.Beer{
			{Id: 1, Name: "aaa"},
			{Id: 2, Name: "bbb"},
			{Id: 3, Name: "ccc"},
		},
	}
}

type mockBeerDAO struct {
	records []models.Beer
}

func (m *mockBeerDAO) Get(rs app.RequestScope, id int) (*models.Beer, error) {
	for _, record := range m.records {
		if record.Id == id {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *mockBeerDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Beer, error) {
	return m.records[offset : offset+limit], nil
}

func (m *mockBeerDAO) Count(rs app.RequestScope) (int, error) {
	return len(m.records), nil
}

func (m *mockBeerDAO) Create(rs app.RequestScope, beer *models.Beer) error {
	if beer.Id != 0 {
		return errors.New("Id cannot be set")
	}
	beer.Id = len(m.records) + 1
	m.records = append(m.records, *beer)
	return nil
}

func (m *mockBeerDAO) Update(rs app.RequestScope, id int, beer *models.Beer) error {
	beer.Id = id
	for i, record := range m.records {
		if record.Id == id {
			m.records[i] = *beer
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockBeerDAO) Delete(rs app.RequestScope, id int) error {
	for i, record := range m.records {
		if record.Id == id {
			m.records = append(m.records[:i], m.records[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
