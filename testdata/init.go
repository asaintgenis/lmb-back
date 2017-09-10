package testdata

import (
	_ "github.com/lib/pq" // initialize posgresql for test
	"gitlab.com/locatemybeer/lmb-back/app"
	"github.com/jinzhu/gorm"
	"gitlab.com/locatemybeer/lmb-back/models"
)

var (
	DB *gorm.DB
)

func init() {
	// the test may be started from the home directory or a subdirectory
	err := app.LoadConfig("./config", "../config")
	if err != nil {
		panic(err)
	}
}

// ResetDB re-create the database schema and re-populate the initial data using the SQL statements in db.sql.
// This method is mainly used in tests.
func ResetDB() *gorm.DB {
	db, err := gorm.Open("postgres", app.Config.DSN)
	if err != nil {
		panic(err)
	}
	db.DropTable(&models.Beer{})
	err = db.AutoMigrate(&models.Beer{}).Error
	if err != nil {
		panic(err)
	}
	return db
}

func CreateBeerData(db *gorm.DB) {
	var records = []models.Beer{
		{Name: "aaa", Content:"aaa content"},
		{Name: "bbb", Content:"bbb content"},
		{Name: "ccc", Content:"ccc content"},
	}

	for _, r := range records {
		db.Create(&r)
	}
}
