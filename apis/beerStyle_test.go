package apis

import (
	"net/http"
	"testing"

	"gitlab.com/locatemybeer/lmb-back/daos"
	"gitlab.com/locatemybeer/lmb-back/services"
	"gitlab.com/locatemybeer/lmb-back/testdata"
)

func TestBeerStyle(t *testing.T) {
	testdata.ResetDB()
	router := newRouter()
	ServeBeerStyleResource(&router.RouteGroup, services.NewBeerStyleService(daos.NewBeerStyleDAO()))

	notFoundError := `{"error_code":"NOT_FOUND", "message":"NOT_FOUND"}`
	nameRequiredError := `{"error_code":"INVALID_DATA","message":"INVALID_DATA","details":[{"field":"name","error":"cannot be blank"}]}`

	runAPITests(t, router, []apiTestCase{
		{"t1 - get a beer", "GET", "/beerStyle/2", "", http.StatusOK, `{"id":2, "srm":2, "name":"India Pale Ale", "color":"#D5BC26" "ebc":12}`},
		{"t2 - get a nonexisting beer", "GET", "/beerStyle/99999", "", http.StatusNotFound, notFoundError},
		{"t3 - create an beer", "POST", "/beerStyle", `{"name":"Qiang"}`, http.StatusOK, `{"id": 4, "name":"Qiang"}`},
		{"t4 - create an beer with validation error", "POST", "/beerStyle", `{"name":""}`, http.StatusBadRequest, nameRequiredError},
		{"t5 - update an beer", "PUT", "/beerStyle/2", `{"name":"Qiang"}`, http.StatusOK, `{"id": 2, "name":"Qiang"}`},
		{"t6 - update an beer with validation error", "PUT", "/beerStyle/2", `{"name":""}`, http.StatusBadRequest, nameRequiredError},
		{"t7 - update a nonexisting beer", "PUT", "/beerStyle/99999", "{}", http.StatusNotFound, notFoundError},
		{"t8 - delete an beer", "DELETE", "/beerStyle/2", ``, http.StatusOK, `{"id": 2, "name":"Qiang"}`},
		{"t9 - delete a nonexisting beer", "DELETE", "/beerStyle/99999", "", http.StatusNotFound, notFoundError},
		{"t10 - get a list of beers", "GET", "/beerStyle?page=1&per_page=2", "", http.StatusOK, `{"page":1, "per_page":2, "page_count":2,"total_count":3,"items":[{"id":1,"name":"1664"},{"id":3,"name":"Guinness"}]}`},
	})
}
