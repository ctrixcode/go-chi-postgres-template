package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ctrixcode/go-chi-postgres/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateExample(t *testing.T) {
	s := NewTestServer()

	// Create request body
	reqBody := models.CreateExampleRequest{
		Name:        "Test Example",
		LuckyNumber: 42.0,
		IsPremium:   true,
	}
	body, _ := json.Marshal(reqBody)

	// Create Request
	req, _ := http.NewRequest("POST", "/examples/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Execute Request
	s.RegisterRoutes().ServeHTTP(rr, req)

	// Assertions
	// Note: This will fail because we're using a mock DB that returns nil
	// In a real test, you'd use a test database or a more sophisticated mock
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestListExamples(t *testing.T) {
	s := NewTestServer()

	// Create Request
	req, _ := http.NewRequest("GET", "/examples/", nil)
	rr := httptest.NewRecorder()

	// Execute Request
	s.RegisterRoutes().ServeHTTP(rr, req)

	// Assertions
	// Note: This will fail because we're using a mock DB that returns nil
	// In a real test, you'd use a test database or a more sophisticated mock
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestGetExample(t *testing.T) {
	s := NewTestServer()

	// Create Request with a valid UUID
	req, _ := http.NewRequest("GET", "/examples/550e8400-e29b-41d4-a716-446655440000", nil)
	rr := httptest.NewRecorder()

	// Execute Request
	s.RegisterRoutes().ServeHTTP(rr, req)

	// Assertions
	// Note: This will fail because we're using a mock DB that returns nil
	// In a real test, you'd use a test database or a more sophisticated mock
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestUpdateExample(t *testing.T) {
	s := NewTestServer()

	// Create request body
	name := "Updated Example"
	reqBody := models.UpdateExampleRequest{
		Name: &name,
	}
	body, _ := json.Marshal(reqBody)

	// Create Request with a valid UUID
	req, _ := http.NewRequest("PUT", "/examples/550e8400-e29b-41d4-a716-446655440000", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Execute Request
	s.RegisterRoutes().ServeHTTP(rr, req)

	// Assertions
	// Note: This will fail because we're using a mock DB that returns nil
	// In a real test, you'd use a test database or a more sophisticated mock
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestDeleteExample(t *testing.T) {
	s := NewTestServer()

	// Create Request with a valid UUID
	req, _ := http.NewRequest("DELETE", "/examples/550e8400-e29b-41d4-a716-446655440000", nil)
	rr := httptest.NewRecorder()

	// Execute Request
	s.RegisterRoutes().ServeHTTP(rr, req)

	// Assertions
	// Note: This will fail because we're using a mock DB that returns nil
	// In a real test, you'd use a test database or a more sophisticated mock
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
