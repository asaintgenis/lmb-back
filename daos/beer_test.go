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
	dao := NewBeerDAO()

	{
		// Get
		testDBCall(db, func(rs app.RequestScope) {
			beer, err := dao.Get(rs, 2)
			assert.Nil(t, err)
			if assert.NotNil(t, beer) {
				assert.Equal(t, 2, beer.Id)
			}
		})
	}

	{
		// Create
		testDBCall(db, func(rs app.RequestScope) {
			beer := &models.Beer{
				Id:   1000,
				Name: "tester",
			}
			err := dao.Create(rs, beer)
			assert.Nil(t, err)
			assert.NotEqual(t, 1000, beer.Id)
			assert.NotZero(t, beer.Id)
		})
	}

	{
		// Update
		testDBCall(db, func(rs app.RequestScope) {
			beer := &models.Beer{
				Id:   2,
				Name: "tester",
			}
			err := dao.Update(rs, beer.Id, beer)
			assert.Nil(t, err)
		})
	}

	{
		// Update with error
		testDBCall(db, func(rs app.RequestScope) {
			beer := &models.Beer{
				Id:   2,
				Name: "tester",
			}
			err := dao.Update(rs, 99999, beer)
			assert.NotNil(t, err)
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
