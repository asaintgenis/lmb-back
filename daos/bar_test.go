package daos

import (
	"testing"

	"gitlab.com/locatemybeer/lmb-back/app"
	"gitlab.com/locatemybeer/lmb-back/models"
	"gitlab.com/locatemybeer/lmb-back/testdata"
	"github.com/stretchr/testify/assert"
)

func TestBarDAO(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewBarDAO()

	{
		// Get
		testDBCall(db, func(rs app.RequestScope) {
			bar, err := dao.Get(rs, 2)
			assert.Nil(t, err)
			if assert.NotNil(t, bar) {
				assert.Equal(t, 2, bar.Id)
			}
		})
	}

	{
		// Create
		testDBCall(db, func(rs app.RequestScope) {
			bar := &models.Bar{
				Id:   1000,
				Name: "tester",
			}
			err := dao.Create(rs, bar)
			assert.Nil(t, err)
			assert.NotEqual(t, 1000, bar.Id)
			assert.NotZero(t, bar.Id)
		})
	}

	{
		// Update
		testDBCall(db, func(rs app.RequestScope) {
			bar := &models.Bar{
				Id:   2,
				Name: "tester",
			}
			err := dao.Update(rs, bar.Id, bar)
			assert.Nil(t, err)
		})
	}

	{
		// Update with error
		testDBCall(db, func(rs app.RequestScope) {
			bar := &models.Bar{
				Id:   2,
				Name: "tester",
			}
			err := dao.Update(rs, 99999, bar)
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

