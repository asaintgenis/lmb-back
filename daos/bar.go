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
func (dao *BarDAO) Get(rs app.RequestScope, id uint) (*models.Bar, error) {
	var bar models.Bar
	err := rs.DB().Where("id = ?", id).First(&bar).Error
	return &bar, err
}

// Create saves a new beer record in the database.
// The Beer.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *BarDAO) Create(rs app.RequestScope, bar *models.Bar) error {
	bar.ID = 0
	return rs.DB().Create(&bar).Error
}

// Update saves the changes to a beer in the database.
func (dao *BarDAO) Update(rs app.RequestScope, bar *models.Bar) error {
	return rs.DB().Save(&bar).Error
}

// Delete deletes a beer with the specified ID from the database.
func (dao *BarDAO) Delete(rs app.RequestScope, id uint) error {
	bar, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.DB().Delete(&bar).Error
}

// Count returns the number of the beer records in the database.
func (dao *BarDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.DB().Table("bars").Count(&count).Error
	return count, err
}

// Query retrieves the beer records with the specified offset and limit from the database.
func (dao *BarDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Bar, error) {
	bars := []models.Bar{}
	err := rs.DB().Limit(limit).Offset(offset).Find(&bars).Error
	return bars, err
}
