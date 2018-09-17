package database_test

import (
	"testing"

	"github.com/tvacare/web-crawler/database"
)

func TestCreateConnection(t *testing.T) {
	// Create database connection
	err := database.NewDB()

	if err != nil {
		t.Errorf("Error trying to get a new database connection %v", err)
	}
}
