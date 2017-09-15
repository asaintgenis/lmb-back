package apis

import (
	"net/http"
	"testing"

	"gitlab.com/locatemybeer/lmb-back/daos"
	"gitlab.com/locatemybeer/lmb-back/services"
	"gitlab.com/locatemybeer/lmb-back/testdata"
)

func TestBeer(t *testing.T) {
	db := testdata.ResetDB()
	testdata.CreateBeerData(db)
	defer db.Close()

	router := newRouter()

	ServeBeerResource(&router.RouteGroup, services.NewBeerService(daos.NewBeerDAO()))

	runAPITests(t, router, []apiTestCase{
		{"t1 - get a beer", "GET", "/beers/2", "", http.StatusOK},
		{"t2 - get a nonexisting beer", "GET", "/beers/99999", "", http.StatusNotFound},
		{"t3 - create an beer", "POST", "/beers", `{"name":"Qiang","content":"Toto"}`, http.StatusOK},
		{"t4 - create an beer with validation error", "POST", "/beers", `{"name":""}`, http.StatusBadRequest},
		{"t5 - update an beer", "PUT", "/beers/2", `{"name":"Qiang"}`, http.StatusOK},
		{"t6 - update an beer with validation error", "PUT", "/beers/2", `{"name":""}`, http.StatusBadRequest},
		{"t7 - update a nonexisting beer", "PUT", "/beers/99999", "{}", http.StatusNotFound},
		{"t8 - delete an beer", "DELETE", "/beers/2", ``, http.StatusOK},
		{"t9 - delete a nonexisting beer", "DELETE", "/beers/99999", "", http.StatusNotFound},
		{"t10 - get a list of beers", "GET", "/beers?page=1&per_page=2", "", http.StatusOK},
	})
}
