package repository

import (
	"context"
	"testing"
	"usersvc/internal/domain"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type testDBAdapter struct{ DB *gorm.DB }
func (a *testDBAdapter) GormDB() *gorm.DB { return a.DB }

func newInMemoryDB(t *testing.T) *testDBAdapter {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil { t.Fatalf("open db: %v", err) }
	if err := db.AutoMigrate(&domain.User{}); err != nil { t.Fatalf("migrate: %v", err) }
	return &testDBAdapter{DB: db}
}

func TestUserRepository_GetByID(t *testing.T) {
	a := newInMemoryDB(t)
	db := a.GormDB()
	if err := db.Create(&domain.User{ID: 1, Email: "a@b.com", Name: "Bob"}).Error; err != nil { t.Fatalf("seed: %v", err) }

	repo := NewUserRepository(a)
	u, err := repo.GetByID(context.Background(), "1")
	assert.NoError(t, err)
	assert.Equal(t, "Bob", u.Name)
}
