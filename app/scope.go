package app

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

// RequestScope contains the application-specific information that are carried around in a request.
type RequestScope interface {
	Logger
	// UserID returns the ID of the user for the current request
	UserID() string
	// SetUserID sets the ID of the currently authenticated user
	SetUserID(id string)
	// RequestID returns the ID of the current request
	RequestID() string
	// Set DB connection
	SetDB(db *gorm.DB)
	// Get DB connection
	DB() *gorm.DB
	// Now returns the timestamp representing the time when the request is being processed
	Now() time.Time
}

type requestScope struct {
	Logger              // the logger tagged with the current request information
	now       time.Time // the time when the request is being processed
	requestID string    // an ID identifying one or multiple correlated HTTP requests
	userID    string    // an ID identifying the current user
	db        *gorm.DB  // the db connection
}

func (rs *requestScope) UserID() string {
	return rs.userID
}

func (rs *requestScope) SetUserID(id string) {
	rs.Logger.SetField("UserID", id)
	rs.userID = id
}

func (rs *requestScope) RequestID() string {
	return rs.requestID
}

func (rs *requestScope) SetDB(db *gorm.DB) {
	rs.db = db
}

func (rs *requestScope) DB() *gorm.DB {
	return rs.db
}

func (rs *requestScope) Now() time.Time {
	return rs.now
}

// newRequestScope creates a new RequestScope with the current request information.
func newRequestScope(now time.Time, logger *logrus.Logger, db *gorm.DB, request *http.Request) RequestScope {
	l := NewLogger(logger, logrus.Fields{})
	requestID := request.Header.Get("X-Request-Id")
	if requestID != "" {
		l.SetField("RequestID", requestID)
	}
	return &requestScope{
		Logger:    l,
		now:       now,
		requestID: requestID,
		db:        db,
	}
}
