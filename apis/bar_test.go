package apis

import (
	"net/http"
	"testing"

	"gitlab.com/locatemybeer/lmb-back/daos"
	"gitlab.com/locatemybeer/lmb-back/services"
	"gitlab.com/locatemybeer/lmb-back/testdata"
)

func TestBar(t *testing.T) {
	db := testdata.ResetDB()
	testdata.CreateBarData(db)
	router := newRouter()
	ServeBarResource(&router.RouteGroup, services.NewBarService(daos.NewBarDAO()))

	runAPITests(t, router, []apiTestCase{
		{"t1 - get a bar", "GET", "/bars/2", "", http.StatusOK},
		{"t2 - get a nonexisting bar", "GET", "/bars/99999", "", http.StatusNotFound},
		{"t3 - create an bar", "POST", "/bars", `{"name":"Qiang"}`, http.StatusOK},
		{"t4 - create an bar with validation error", "POST", "/bars", `{"name":""}`, http.StatusBadRequest},
		{"t5 - update an bar", "PUT", "/bars/2", `{"name":"Qiang"}`, http.StatusOK},
		{"t6 - update an bar with validation error", "PUT", "/bars/2", `{"name":""}`, http.StatusBadRequest},
		{"t7 - update a nonexisting bar", "PUT", "/bars/99999", "{}", http.StatusNotFound},
		{"t8 - delete an bar", "DELETE", "/bars/2", ``, http.StatusOK},
		{"t9 - delete a nonexisting bar", "DELETE", "/bars/99999", "", http.StatusNotFound},
		{"t10 - get a list of bars", "GET", "/bars?page=1&per_page=2", "", http.StatusOK},
	})
}
