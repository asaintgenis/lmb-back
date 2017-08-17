package apis

import (
	"net/http"
	"testing"

	"gitlab.com/locatemybeer/lmb-back/daos"
	"gitlab.com/locatemybeer/lmb-back/services"
	"gitlab.com/locatemybeer/lmb-back/testdata"
)

func TestBar(t *testing.T) {
	testdata.ResetDB()
	router := newRouter()
	ServeBarResource(&router.RouteGroup, services.NewBarService(daos.NewBarDAO()))

	notFoundError := `{"error_code":"NOT_FOUND", "message":"NOT_FOUND"}`
	nameRequiredError := `{"error_code":"INVALID_DATA","message":"INVALID_DATA","details":[{"field":"name","error":"cannot be blank"}]}`

	runAPITests(t, router, []apiTestCase{
		{"t1 - get a bar", "GET", "/bars/2", "", http.StatusOK, `{"id":2,"name":"Guinness tavern"}`},
		{"t2 - get a nonexisting bar", "GET", "/bars/99999", "", http.StatusNotFound, notFoundError},
		{"t3 - create an bar", "POST", "/bars", `{"name":"Qiang"}`, http.StatusOK, `{"id": 4, "name":"Qiang"}`},
		{"t4 - create an bar with validation error", "POST", "/bars", `{"name":""}`, http.StatusBadRequest, nameRequiredError},
		{"t5 - update an bar", "PUT", "/bars/2", `{"name":"Qiang"}`, http.StatusOK, `{"id": 2, "name":"Qiang"}`},
		{"t6 - update an bar with validation error", "PUT", "/bars/2", `{"name":""}`, http.StatusBadRequest, nameRequiredError},
		{"t7 - update a nonexisting bar", "PUT", "/bars/99999", "{}", http.StatusNotFound, notFoundError},
		{"t8 - delete an bar", "DELETE", "/bars/2", ``, http.StatusOK, `{"id": 2, "name":"Qiang"}`},
		{"t9 - delete a nonexisting bar", "DELETE", "/bars/99999", "", http.StatusNotFound, notFoundError},
		{"t10 - get a list of bars", "GET", "/bars?page=1&per_page=2", "", http.StatusOK, `{"page":1, "per_page":2, "page_count":2,"total_count":3,"items":[{"id":1,"name":"La Kolok"},{"id":3,"name":"King georges"}]}`},
	})
}
