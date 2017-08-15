package daos

import (
	"gitlab.com/locatemybeer/lmb-back/app"
	"gitlab.com/locatemybeer/lmb-back/models"
)

// ArtistDAO persists artist data in database
type BeerDAO struct{}

// NewArtistDAO creates a new ArtistDAO
func NewArtistDAO() *BeerDAO {
	return &BeerDAO{}
}

// Get reads the artist with the specified ID from the database.
func (dao *BeerDAO) Get(rs app.RequestScope, id int) (*models.Beer, error) {
	var artist models.Beer
	err := rs.Tx().Select().Model(id, &artist)
	return &artist, err
}

// Create saves a new artist record in the database.
// The Artist.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *BeerDAO) Create(rs app.RequestScope, artist *models.Beer) error {
	artist.Id = 0
	return rs.Tx().Model(artist).Insert()
}

// Update saves the changes to an artist in the database.
func (dao *BeerDAO) Update(rs app.RequestScope, id int, artist *models.Beer) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	artist.Id = id
	return rs.Tx().Model(artist).Exclude("Id").Update()
}

// Delete deletes an artist with the specified ID from the database.
func (dao *BeerDAO) Delete(rs app.RequestScope, id int) error {
	artist, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(artist).Delete()
}

// Count returns the number of the artist records in the database.
func (dao *BeerDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("artist").Row(&count)
	return count, err
}

// Query retrieves the artist records with the specified offset and limit from the database.
func (dao *BeerDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Beer, error) {
	artists := []models.Beer{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&artists)
	return artists, err
}
