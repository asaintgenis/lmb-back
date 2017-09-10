package daos

import (
	"testing"

	"gitlab.com/locatemybeer/lmb-back/app"
	"gitlab.com/locatemybeer/lmb-back/models"
	"gitlab.com/locatemybeer/lmb-back/testdata"
	"github.com/stretchr/testify/assert"
)

func TestBeerDAO(t *testing.T) {
	db := testdata.ResetDB()
	testdata.CreateBeerData(db)
	defer db.Close()


	dao := NewBeerDAO()

	{
		// Get
		testDBCall(db, func(rs app.RequestScope) {
			beer, err := dao.Get(rs, 2)
			assert.Nil(t, err)
			if assert.NotNil(t, beer) {
				assert.Equal(t, uint(2), beer.ID)
				assert.Equal(t, "bbb", beer.Name)
			}
		})
	}

	{
		// Create
		testDBCall(db, func(rs app.RequestScope) {
			beer := &models.Beer{
				Name: "tester",
				Content: "test content",
			}
			err := dao.Create(rs, beer)
			assert.Nil(t, err)
			assert.NotZero(t, beer.ID)
		})
	}

	{
		// Update
		testDBCall(db, func(rs app.RequestScope) {
			beer, err := dao.Get(rs, 2)
			assert.Nil(t, err)
			beer.Content = "modified test"
			err = dao.Update(rs, beer)
			assert.Nil(t, err)

			beer, err = dao.Get(rs, 2)
			assert.Nil(t, err)
			assert.Equal(t, "modified test", beer.Content)
		})
	}

	{
		// Delete
		testDBCall(db, func(rs app.RequestScope) {
			err := dao.Delete(rs, 2)
			assert.Nil(t, err)
		})
	}

	{
		// Delete with error
		testDBCall(db, func(rs app.RequestScope) {
			err := dao.Delete(rs, 99999)
			assert.NotNil(t, err)
		})
	}

	{
		// Query
		testDBCall(db, func(rs app.RequestScope) {
			beers, err := dao.Query(rs, 1, 3)
			assert.Nil(t, err)
			assert.Equal(t, 2, len(beers))
		})
	}

	{
		// Count
		testDBCall(db, func(rs app.RequestScope) {
			count, err := dao.Count(rs)
			assert.Nil(t, err)
			assert.NotZero(t, count)
		})
	}
}
