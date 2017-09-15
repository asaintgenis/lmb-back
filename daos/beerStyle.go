package daos

import (
	"gitlab.com/locatemybeer/lmb-back/app"
	"gitlab.com/locatemybeer/lmb-back/models"
)

// BeerStyle persists beerStyle data in database
type BeerStyleDAO struct{}

// NewBeerStyleDAO creates a new BeerStyleDAO
func NewBeerStyleDAO() *BeerStyleDAO {
	return &BeerStyleDAO{}
}

// Get reads the beerStyle with the specified ID from the database.
func (dao *BeerStyleDAO) Get(rs app.RequestScope, id uint) (*models.BeerStyle, error) {
	var beerStyle models.BeerStyle
	err := rs.DB().Where("id = ?", id).First(&beerStyle).Error
	return &beerStyle, err
}

// Create saves a new beerStyle record in the database.
// The Beer.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *BeerStyleDAO) Create(rs app.RequestScope, beerStyle *models.BeerStyle) error {
	beerStyle.ID = 0
	return rs.DB().Create(&beerStyle).Error
}

// Update saves the changes to a beerStyle in the database.
func (dao *BeerStyleDAO) Update(rs app.RequestScope, beerStyle *models.BeerStyle) error {
	return rs.DB().Save(&beerStyle).Error
}

// Delete deletes a beerStyle with the specified ID from the database.
func (dao *BeerStyleDAO) Delete(rs app.RequestScope, id uint) error {
	beerStyle, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.DB().Delete(&beerStyle).Error
}

// Count returns the number of the beerStyle records in the database.
func (dao *BeerStyleDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.DB().Table("beerStyles").Count(&count).Error
	return count, err
}

// Query retrieves the beerStyle records with the specified offset and limit from the database.
func (dao *BeerStyleDAO) Query(rs app.RequestScope, offset, limit int) ([]models.BeerStyle, error) {
	beerStyles := []models.BeerStyle{}
	err := rs.DB().Limit(limit).Offset(offset).Find(&beerStyles).Error
	return beerStyles, err
}