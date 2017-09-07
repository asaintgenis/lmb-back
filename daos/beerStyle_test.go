package daos

import (
	"testing"

	"gitlab.com/locatemybeer/lmb-back/app"
	"gitlab.com/locatemybeer/lmb-back/models"
	"gitlab.com/locatemybeer/lmb-back/testdata"
	"github.com/stretchr/testify/assert"
)

func TestBeerStyleDAO(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewBeerStyleDAO()

	{
		// Get
		testDBCall(db, func(rs app.RequestScope) {
			beerStyle, err := dao.Get(rs, 2)
			assert.Nil(t, err)
			if assert.NotNil(t, beerStyle) {
				assert.Equal(t, 2, beerStyle.Id)
			}
		})
	}

	{
		// Create
		testDBCall(db, func(rs app.RequestScope) {
			beerStyle := &models.BeerStyle{
				Id:   1000,
				Name: "tester",
			}
			err := dao.Create(rs, beerStyle)
			assert.Nil(t, err)
			assert.NotEqual(t, 1000, beerStyle.Id)
			assert.NotZero(t, beerStyle.Id)
		})
	}

	{
		// Update
		testDBCall(db, func(rs app.RequestScope) {
			beerStyle := &models.BeerStyle{
				Id:   2,
				Name: "tester",
			}
			err := dao.Update(rs, beerStyle.Id, beerStyle)
			assert.Nil(t, err)
		})
	}

	{
		// Update with error
		testDBCall(db, func(rs app.RequestScope) {
			beerStyle := &models.BeerStyle{
				Id:   2,
				Name: "tester",
			}
			err := dao.Update(rs, 99999, beerStyle)
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
			beerStyles, err := dao.Query(rs, 1, 3)
			assert.Nil(t, err)
			assert.Equal(t, 2, len(beerStyles))
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