package ports

import "gorm.io/gorm"

type DatabaseConnection interface {
	GormDB() *gorm.DB
}
