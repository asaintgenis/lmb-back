package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/locatemybeer/lmb-back/app"
	"gitlab.com/locatemybeer/lmb-back/models"
)

func TestNewBarService(t *testing.T) {
	dao := newMockBarDAO()
	s := NewBarService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestBarService_Get(t *testing.T) {
	s := NewBarService(newMockBarDAO())
	bar, err := s.Get(nil, 1)
	if assert.Nil(t, err) && assert.NotNil(t, bar) {
		assert.Equal(t, "aaa", bar.Name)
	}

	bar, err = s.Get(nil, 100)
	assert.NotNil(t, err)
}

func TestBarService_Create(t *testing.T) {
	s := NewBarService(newMockBarDAO())
	bar, err := s.Create(nil, &models.Bar{
		Name: "ddd",
	})
	if assert.Nil(t, err) && assert.NotNil(t, bar) {
		assert.Equal(t, uint(4), bar.ID)
		assert.Equal(t, "ddd", bar.Name)
	}

	// validation error
	_, err = s.Create(nil, &models.Bar{
		Name: "",
	})
	assert.NotNil(t, err)
}

func TestBarService_Update(t *testing.T) {
	s := NewBarService(newMockBarDAO())
	barToUpdate := models.Bar{
		Name: "ddd",
	}
	barToUpdate.ID = 2
	bar, err := s.Update(nil, 2, &barToUpdate)
	if assert.Nil(t, err) && assert.NotNil(t, bar) {
		assert.Equal(t, uint(2), bar.ID)
		assert.Equal(t, "ddd", bar.Name)
	}
	// validation error
	_, err = s.Update(nil, 2, &models.Bar{
		Name: "",
	})
	assert.NotNil(t, err)
}

func TestBarService_Delete(t *testing.T) {
	s := NewBarService(newMockBarDAO())
	bar, err := s.Delete(nil, 2)
	if assert.Nil(t, err) && assert.NotNil(t, bar) {
		assert.Equal(t, uint(2), bar.ID)
		assert.Equal(t, "bbb", bar.Name)
	}

	_, err = s.Delete(nil, 2)
	assert.NotNil(t, err)
}

func TestBarService_Query(t *testing.T) {
	s := NewBarService(newMockBarDAO())
	result, err := s.Query(nil, 1, 2)
	if assert.Nil(t, err) {
		assert.Equal(t, 2, len(result))
	}
}

func newMockBarDAO() barDao {
	mockBarDAO := &mockBarDAO{
		records: []models.Bar{
			{Name: "aaa"},
			{Name: "bbb"},
			{Name: "ccc"},
		},
	}
	mockBarDAO.prepareMockData()
	return mockBarDAO
}

type mockBarDAO struct {
	records []models.Bar
}

func (m *mockBarDAO) prepareMockData() {
	for index := range m.records {
		m.records[index].ID = uint(index + 1)
	}
}

func (m *mockBarDAO) Get(rs app.RequestScope, id uint) (*models.Bar, error) {
	for _, record := range m.records {
		if record.ID == id {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *mockBarDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Bar, error) {
	return m.records[offset : offset+limit], nil
}

func (m *mockBarDAO) Count(rs app.RequestScope) (int, error) {
	return len(m.records), nil
}

func (m *mockBarDAO) Create(rs app.RequestScope, bar *models.Bar) error {
	if bar.ID != 0 {
		return errors.New("Id cannot be set")
	}
	bar.ID = uint(len(m.records) + 1)
	m.records = append(m.records, *bar)
	return nil
}

func (m *mockBarDAO) Update(rs app.RequestScope, bar *models.Bar) error {
	for i, record := range m.records {
		if record.ID == bar.ID {
			m.records[i] = *bar
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockBarDAO) Delete(rs app.RequestScope, id uint) error {
	for i, record := range m.records {
		if record.ID == id {
			m.records = append(m.records[:i], m.records[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
