package daos

import (
	"time"

	"github.com/jinzhu/gorm"
	"gitlab.com/locatemybeer/lmb-back/app"
)

func testDBCall(db *gorm.DB, f func(rs app.RequestScope)) {
	rs := mockRequestScope(db)

	defer func() {
		rs.DB().Rollback()
	}()

	f(rs)
}

type requestScope struct {
	app.Logger
	db *gorm.DB
}

func mockRequestScope(db *gorm.DB) app.RequestScope {
	db = db.Begin()
	return &requestScope{
		db: db,
	}
}

func (rs *requestScope) UserID() string {
	return "tester"
}

func (rs *requestScope) SetUserID(id string) {
}

func (rs *requestScope) RequestID() string {
	return "test"
}

func (rs *requestScope) DB() *gorm.DB {
	return rs.db
}

func (rs *requestScope) SetDB(db *gorm.DB) {
	rs.db = db
}

func (rs *requestScope) Now() time.Time {
	return time.Now()
}
