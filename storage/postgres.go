package storage

import (
	"github.com/akhil-is-watching/goshort/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresStore struct {
	db *gorm.DB
}

func NewPostgresStore() *PostgresStore {
	db, err := gorm.Open(postgres.Open(""))

	if err != nil {
		panic("DB Execution Failed")
	}

	err = db.AutoMigrate(&types.Goshort{})
	if err != nil {
		panic("DB Migration failed")
	}

	return &PostgresStore{
		db: db,
	}
}
