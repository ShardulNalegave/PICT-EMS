package database

import (
	"github.com/ShardulNalegave/PICT-EMS/database/models"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	DBNAME = "ems.db"
)

func ConnectToDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DBNAME), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Couldn't connect to database")
	}

	db.AutoMigrate(&models.Student{})
	db.AutoMigrate(&models.StaffMember{})

	log.Info().Msg("Connected to Sqlite")
	return db
}
