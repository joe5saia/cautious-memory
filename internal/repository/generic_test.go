package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Create a gorm model for testing
type testProfile struct {
	gorm.Model
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FirstName string    `gorm:"type:varchar(255)"`
	LastName  string    `gorm:"type:varchar(255)"`
}

func TestNewGormRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening gorm database", err)
	}

	repo := NewGormRepository[testProfile](gdb)

	// Type assertion to access the db field
	gormRepo, ok := repo.(*gormRepository[testProfile])
	if !ok {
		t.Fatalf("expected *gormRepository[testProfile], got %T", repo)
	}

	assert.Equal(t, gdb, gormRepo.db)
}
