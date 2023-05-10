package storage

import (
	"fmt"

	"github.com/whiskerside/myshopify/conf"
	"github.com/whiskerside/myshopify/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GenPostgresConnection() (db *gorm.DB) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		conf.Env.DbHost, conf.Env.DbPort, conf.Env.DbUserName, conf.Env.DbName, conf.Env.DbPassword)

	log.Logger.Info().Msg(fmt.Sprintf("Connecting to database dsn: %s", dsn))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		FullSaveAssociations: true,
	})
	if err != nil {
		log.Logger.Fatal().Err(err).Msg("Connect PostgreSQL failed, exit.")
	}
	return
}
