package daos

import (
"gitlab.com/locatemybeer/lmb-back/app"
"gitlab.com/locatemybeer/lmb-back/models"
)

// BeerDAO persists beer data in database
type BarDAO struct{}

// NewBeerDAO creates a new BeerDAO
func NewBarDAO() *BarDAO {
	return &BarDAO{}
}

// Get reads the beer with the specified ID from the database.
func (dao *BarDAO) Get(rs app.RequestScope, id int) (*models.Bar, error) {
	var bar models.Bar
	err := rs.Tx().Select().Model(id, &bar)
	return &bar, err
}

// Create saves a new beer record in the database.
// The Beer.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *BarDAO) Create(rs app.RequestScope, bar *models.Bar) error {
	bar.Id = 0
	return rs.Tx().Model(bar).Insert()
}

// Update saves the changes to a beer in the database.
func (dao *BarDAO) Update(rs app.RequestScope, id int, bar *models.Bar) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	bar.Id = id
	return rs.Tx().Model(bar).Exclude("Id").Update()
}

// Delete deletes a beer with the specified ID from the database.
func (dao *BarDAO) Delete(rs app.RequestScope, id int) error {
	bar, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(bar).Delete()
}

// Count returns the number of the beer records in the database.
func (dao *BarDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("bar").Row(&count)
	return count, err
}

// Query retrieves the beer records with the specified offset and limit from the database.
func (dao *BarDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Bar, error) {
	bars := []models.Bar{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&bars)
	return bars, err
}
