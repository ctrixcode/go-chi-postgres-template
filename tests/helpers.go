package tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ctrixcode/go-chi-postgres/internal/config"
	"github.com/ctrixcode/go-chi-postgres/internal/server"
	"github.com/jmoiron/sqlx"
)

type mockDB struct {
	db *sqlx.DB
}

func (m *mockDB) Health() map[string]string {
	return map[string]string{"status": "up"}
}

func (m *mockDB) Close() error {
	return m.db.Close()
}

func (m *mockDB) GetDB() *sqlx.DB {
	return m.db
}

// NewTestServer creates a server with a mock database for testing
func NewTestServer() *server.Server {
	cfg := &config.Config{
		Port: 8080,
	}

	// Create a mock database connection
	sqlDB, _, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	sqlxDB := sqlx.NewDb(sqlDB, "sqlmock")

	db := &mockDB{
		db: sqlxDB,
	}

	return server.NewServer(cfg, db)
}

// NewTestServerWithMock creates a server and returns both the server and the sqlmock
// This allows tests to set expectations on the mock
func NewTestServerWithMock() (*server.Server, sqlmock.Sqlmock) {
	cfg := &config.Config{
		Port: 8080,
	}

	// Create a mock database connection
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	sqlxDB := sqlx.NewDb(sqlDB, "sqlmock")

	db := &mockDB{
		db: sqlxDB,
	}

	return server.NewServer(cfg, db), mock
}
