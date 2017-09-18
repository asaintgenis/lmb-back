package apis

import (
	"net/http"
	"testing"

	"gitlab.com/locatemybeer/lmb-back/daos"
	"gitlab.com/locatemybeer/lmb-back/services"
	"gitlab.com/locatemybeer/lmb-back/testdata"
)

func TestBeerStyle(t *testing.T) {
	db := testdata.ResetDB()
	testdata.CreateBeerStyleData(db)
	router := newRouter()
	ServeBeerStyleResource(&router.RouteGroup, services.NewBeerStyleService(daos.NewBeerStyleDAO()))

	runAPITests(t, router, []apiTestCase{
		{"t1 - get a beer", "GET", "/beersStyles/2", "", http.StatusOK},
		{"t2 - get a nonexisting beer", "GET", "/beersStyles/99999", "", http.StatusNotFound},
		{"t3 - create an beer", "POST", "/beersStyles", `{"name":"Qiang"}`, http.StatusOK},
		{"t4 - create an beer with validation error", "POST", "/beersStyles", `{"name":""}`, http.StatusBadRequest},
		{"t5 - update an beer", "PUT", "/beersStyles/2", `{"name":"Qiang"}`, http.StatusOK},
		{"t6 - update an beer with validation error", "PUT", "/beersStyles/2", `{"name":""}`, http.StatusBadRequest},
		{"t7 - update a nonexisting beer", "PUT", "/beersStyles/99999", "{}", http.StatusNotFound},
		{"t8 - delete an beer", "DELETE", "/beersStyles/2", ``, http.StatusOK},
		{"t9 - delete a nonexisting beer", "DELETE", "/beersStyles/99999", "", http.StatusNotFound},
		{"t10 - get a list of beers", "GET", "/beersStyles?page=1&per_page=2", "", http.StatusOK},
	})
}
