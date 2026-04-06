package ports

import "gorm.io/gorm"

type Rows interface {
	Next() bool
	Scan(dest ...interface{}) error
	Close()
	Err() error
}

type Result interface{}

// DatabaseConnection abstracts access to GORM DB
type DatabaseConnection interface {
	GormDB() *gorm.DB
}
