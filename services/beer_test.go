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
		Content: "rrr",
	})
	if assert.Nil(t, err) && assert.NotNil(t, beer) {
		assert.Equal(t, uint(4), beer.ID)
		assert.Equal(t, "ddd", beer.Name)
		assert.Equal(t, "rrr", beer.Content)
	}

	// validation error
	// Empty Name
	_, err = s.Create(nil, &models.Beer{
		Name: "",
		Content: "rrr",
	})
	assert.NotNil(t, err)

	// validation error
	// Empty Content
	_, err = s.Create(nil, &models.Beer{
		Name: "ddd",
		Content: "",
	})
	assert.NotNil(t, err)
}

func TestBeerService_Update(t *testing.T) {
	s := NewBeerService(newMockBeerDAO())
	beer, err := s.Update(nil, 2, &models.Beer{
		Name: "ddd",
		Content: "rrr",
	})
	if assert.Nil(t, err) && assert.NotNil(t, beer) {
		assert.Equal(t, uint(2), beer.ID)
		assert.Equal(t, "ddd", beer.Name)
		assert.Equal(t,"rrr", beer.Content)
	}

	// validation error
	// Empty Name
	_, err = s.Update(nil, 2, &models.Beer{
		Name: "",
		Content: "rrr",
	})
	assert.NotNil(t, err)

	// validation error
	// Empty Content
	_, err = s.Update(nil, 2, &models.Beer{
		Name: "ddd",
		Content: "",
	})
	assert.NotNil(t, err)
}

func TestBeerService_Delete(t *testing.T) {
	s := NewBeerService(newMockBeerDAO())
	beer, err := s.Delete(nil, 2)
	if assert.Nil(t, err) && assert.NotNil(t, beer) {
		assert.Equal(t, uint(2), beer.ID)
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
	mockBeerDAO := &mockBeerDAO{
		records: []models.Beer{
			{Name: "aaa", Content: "aaa content"},
			{Name: "bbb", Content: "bbb content"},
			{Name: "ccc", Content: "ccc content"},
		},
	}
	mockBeerDAO.prepareMockData()
	return mockBeerDAO
}

type mockBeerDAO struct {
	records []models.Beer
}

func (m *mockBeerDAO) prepareMockData() {
	for index := range m.records {
		m.records[index].ID = uint(index + 1)
	}
}

func (m *mockBeerDAO) Get(rs app.RequestScope, id uint) (*models.Beer, error) {
	for _, record := range m.records {
		if record.ID == id {
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
	if beer.ID != 0 {
		return errors.New("Id cannot be set")
	}
	beer.ID = uint(len(m.records) + 1)
	m.records = append(m.records, *beer)
	return nil
}

func (m *mockBeerDAO) Update(rs app.RequestScope, id uint, beer *models.Beer) error {
	beer.ID = id
	for i, record := range m.records {
		if record.ID == id {
			m.records[i] = *beer
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockBeerDAO) Delete(rs app.RequestScope, id uint) error {
	for i, record := range m.records {
		if record.ID == id {
			m.records = append(m.records[:i], m.records[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
