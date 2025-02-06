package db

import (
	"fmt"
	"github.com/Crabocod/gpt_network/api-service/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(c *config.Config) (*sqlx.DB, error) {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
		c.Database.SSLMode,
	)
	var err error
	db, err := sqlx.Connect("postgres", dbInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
