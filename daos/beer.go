package daos

import (
	"gitlab.com/locatemybeer/lmb-back/app"
	"gitlab.com/locatemybeer/lmb-back/models"
)

// BeerDAO persists beer data in database
type BeerDAO struct{}

// NewBeerDAO creates a new BeerDAO
func NewBeerDAO() *BeerDAO {
	return &BeerDAO{}
}

// Get reads the beer with the specified ID from the database.
func (dao *BeerDAO) Get(rs app.RequestScope, id int) (*models.Beer, error) {
	var beer models.Beer
	err := rs.DB().Where("id = ?", id).First(beer).Error
	return &beer, err
}

// Create saves a new beer record in the database.
// The Beer.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *BeerDAO) Create(rs app.RequestScope, beer *models.Beer) error {
	beer.Id = 0
	return rs.DB().Create(beer).Error
}

// Update saves the changes to a beer in the database.
func (dao *BeerDAO) Update(rs app.RequestScope, id int, beer *models.Beer) error {
	return rs.DB().Save(beer).Error
}

// Delete deletes a beer with the specified ID from the database.
func (dao *BeerDAO) Delete(rs app.RequestScope, id int) error {
	beer, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.DB().Delete(beer).Error
}

// Count returns the number of the beer records in the database.
func (dao *BeerDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.DB().Table("beer").Count(&count).Error
	return count, err
}

// Query retrieves the beer records with the specified offset and limit from the database.
func (dao *BeerDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Beer, error) {
	beers := []models.Beer{}
	err := rs.DB().Limit(limit).Offset(offset).Find(beers).Error
	return beers, err
}
