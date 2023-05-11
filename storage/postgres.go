package storage

import (
	"fmt"

	"github.com/whiskerside/myshopify/conf"
	"github.com/whiskerside/myshopify/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	logs = log.Logger()
	DB   *gorm.DB
)

func GenPostgresConnection() (db *gorm.DB) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		conf.Env.Database.Host, conf.Env.Database.Port,
		conf.Env.Database.Username, conf.Env.Database.Name,
		conf.Env.Database.Password)

	logs.Info().Msg(fmt.Sprintf("Connecting to database dsn: %s", dsn))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		FullSaveAssociations: true,
	})
	if err != nil {
		logs.Fatal().Err(err).Msg("Connect PostgreSQL failed, exit.")
	}
	return
}
