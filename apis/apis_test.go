package apis

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
	"gitlab.com/locatemybeer/lmb-back/app"
	"github.com/stretchr/testify/assert"
	"github.com/jinzhu/gorm"
)

type apiTestCase struct {
	tag      string
	method   string
	url      string
	body     string
	status   int
}

func newRouter() *routing.Router {
	logger := logrus.New()
	logger.Level = logrus.PanicLevel

	//connecting to DB
	db, err := gorm.Open("postgres", app.Config.DSN)
	if err != nil {
		panic(err)
	}

	router := routing.New()

	router.Use(
		app.Init(logger, db),
		content.TypeNegotiator(content.JSON),
	)

	return router
}

func testAPI(router *routing.Router, method, URL, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, URL, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	return res
}

func runAPITests(t *testing.T, router *routing.Router, tests []apiTestCase) {
	for _, test := range tests {
		res := testAPI(router, test.method, test.url, test.body)
		assert.Equal(t, test.status, res.Code, test.tag)
	}
}
