package tests

import (
	"github.com/ctrixcode/go-chi-postgres/internal/config"
	"github.com/ctrixcode/go-chi-postgres/internal/server"
	"github.com/jmoiron/sqlx"
)

type mockDB struct{}

func (m *mockDB) Health() map[string]string {
	return map[string]string{"status": "up"}
}

func (m *mockDB) Close() error {
	return nil
}

func (m *mockDB) GetDB() *sqlx.DB {
	return nil
}

func NewTestServer() *server.Server {
	cfg := &config.Config{
		Port: 8080,
	}
	db := &mockDB{}
	return server.NewServer(cfg, db)
}
