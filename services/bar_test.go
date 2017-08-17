package services

import (
	"errors"
	"testing"

	"gitlab.com/locatemybeer/lmb-back/app"
	"gitlab.com/locatemybeer/lmb-back/models"
	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, 4, bar.Id)
		assert.Equal(t, "ddd", bar.Name)
	}

	// dao error
	_, err = s.Create(nil, &models.Bar{
		Id:   100,
		Name: "ddd",
	})
	assert.NotNil(t, err)

	// validation error
	_, err = s.Create(nil, &models.Bar{
		Name: "",
	})
	assert.NotNil(t, err)
}

func TestBarService_Update(t *testing.T) {
	s := NewBarService(newMockBarDAO())
	bar, err := s.Update(nil, 2, &models.Bar{
		Name: "ddd",
	})
	if assert.Nil(t, err) && assert.NotNil(t, bar) {
		assert.Equal(t, 2, bar.Id)
		assert.Equal(t, "ddd", bar.Name)
	}

	// dao error
	_, err = s.Update(nil, 100, &models.Bar{
		Name: "ddd",
	})
	assert.NotNil(t, err)

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
		assert.Equal(t, 2, bar.Id)
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
	return &mockBarDAO{
		records: []models.Bar{
			{Id: 1, Name: "aaa"},
			{Id: 2, Name: "bbb"},
			{Id: 3, Name: "ccc"},
		},
	}
}

type mockBarDAO struct {
	records []models.Bar
}

func (m *mockBarDAO) Get(rs app.RequestScope, id int) (*models.Bar, error) {
	for _, record := range m.records {
		if record.Id == id {
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
	if bar.Id != 0 {
		return errors.New("Id cannot be set")
	}
	bar.Id = len(m.records) + 1
	m.records = append(m.records, *bar)
	return nil
}

func (m *mockBarDAO) Update(rs app.RequestScope, id int, bar *models.Bar) error {
	bar.Id = id
	for i, record := range m.records {
		if record.Id == id {
			m.records[i] = *bar
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockBarDAO) Delete(rs app.RequestScope, id int) error {
	for i, record := range m.records {
		if record.Id == id {
			m.records = append(m.records[:i], m.records[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
