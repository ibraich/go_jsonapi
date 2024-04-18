package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	// postgres://gclcgrjb:w6cbQADQsl_x_qFBepYgeWRBzRrkE2jX@ella.db.elephantsql.com/gclcgrjb
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot Connect to Database")
	}

}
