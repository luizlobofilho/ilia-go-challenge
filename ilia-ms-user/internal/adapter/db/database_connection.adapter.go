package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConnectionAdapter struct{ DB *gorm.DB }

func NewDatabaseConnectionAdapter(dsn string) (*DatabaseConnectionAdapter, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DatabaseConnectionAdapter{DB: db}, nil
}

func (a *DatabaseConnectionAdapter) GormDB() *gorm.DB { return a.DB }
