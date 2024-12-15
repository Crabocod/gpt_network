package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"web.app/internal/config"
)

var DB *sqlx.DB

func Connect() error {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Data.Database.Host,
		config.Data.Database.Port,
		config.Data.Database.User,
		config.Data.Database.Password,
		config.Data.Database.Name,
		config.Data.Database.SSLMode,
	)
	var err error
	DB, err = sqlx.Connect("postgres", dbInfo)
	if err != nil {
		return err
	}
	return nil
}
