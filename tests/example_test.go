package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ctrixcode/go-chi-postgres/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateExample(t *testing.T) {
	s, mock := NewTestServer()
	defer mock.ExpectationsWereMet()

	// Create request body
	reqBody := models.CreateExampleRequest{
		Name:        "Test Example",
		LuckyNumber: 42.0,
		IsPremium:   true,
	}
	body, _ := json.Marshal(reqBody)

	// Set up mock expectations
	rows := sqlmock.NewRows([]string{"id", "name", "lucky_number", "is_premium", "created_at", "updated_at"}).
		AddRow(uuid.New(), "Test Example", 42.0, true, time.Now(), time.Now())

	mock.ExpectQuery(`INSERT INTO examples`).
		WithArgs("Test Example", 42.0, true).
		WillReturnRows(rows)

	// Create Request
	req, _ := http.NewRequest("POST", "/examples/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Execute Request
	s.RegisterRoutes().ServeHTTP(rr, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, rr.Code)

	var response map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, "Test Example", response["data"].(map[string]interface{})["name"])
}

func TestListExamples(t *testing.T) {
	s, mock := NewTestServer()
	defer mock.ExpectationsWereMet()

	// Set up mock expectations
	rows := sqlmock.NewRows([]string{"id", "name", "lucky_number", "is_premium", "created_at", "updated_at"}).
		AddRow(uuid.New(), "Example 1", 10.0, true, time.Now(), time.Now()).
		AddRow(uuid.New(), "Example 2", 20.0, false, time.Now(), time.Now())

	mock.ExpectQuery(`SELECT \* FROM examples`).
		WillReturnRows(rows)

	// Create Request
	req, _ := http.NewRequest("GET", "/examples/", nil)
	rr := httptest.NewRecorder()

	// Execute Request
	s.RegisterRoutes().ServeHTTP(rr, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Check if data exists and is the right type
	if response["data"] != nil {
		data := response["data"].([]interface{})
		assert.Equal(t, 2, len(data))
	} else {
		t.Logf("Response body: %s", rr.Body.String())
		t.Fatal("response data is nil")
	}
}

func TestGetExample(t *testing.T) {
	s, mock := NewTestServer()
	defer mock.ExpectationsWereMet()

	testID := uuid.New()

	// Set up mock expectations
	rows := sqlmock.NewRows([]string{"id", "name", "lucky_number", "is_premium", "created_at", "updated_at"}).
		AddRow(testID, "Test Example", 42.0, true, time.Now(), time.Now())

	mock.ExpectQuery(`SELECT \* FROM examples`).
		WithArgs(testID).
		WillReturnRows(rows)

	// Create Request
	req, _ := http.NewRequest("GET", "/examples/"+testID.String(), nil)
	rr := httptest.NewRecorder()

	// Execute Request
	s.RegisterRoutes().ServeHTTP(rr, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, "Test Example", response["data"].(map[string]interface{})["name"])
}

func TestUpdateExample(t *testing.T) {
	s, mock := NewTestServer()
	defer mock.ExpectationsWereMet()

	testID := uuid.New()
	name := "Updated Example"
	reqBody := models.UpdateExampleRequest{
		Name: &name,
	}
	body, _ := json.Marshal(reqBody)

	// Set up mock expectations
	rows := sqlmock.NewRows([]string{"id", "name", "lucky_number", "is_premium", "created_at", "updated_at"}).
		AddRow(testID, "Updated Example", 42.0, true, time.Now(), time.Now())

	mock.ExpectQuery(`UPDATE examples`).
		WithArgs(sqlmock.AnyArg(), "Updated Example", testID).
		WillReturnRows(rows)

	// Create Request
	req, _ := http.NewRequest("PUT", "/examples/"+testID.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Execute Request
	s.RegisterRoutes().ServeHTTP(rr, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, "Updated Example", response["data"].(map[string]interface{})["name"])
}

func TestDeleteExample(t *testing.T) {
	s, mock := NewTestServer()
	defer mock.ExpectationsWereMet()

	testID := uuid.New()

	// Set up mock expectations
	mock.ExpectExec(`DELETE FROM examples`).
		WithArgs(testID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Create Request
	req, _ := http.NewRequest("DELETE", "/examples/"+testID.String(), nil)
	rr := httptest.NewRecorder()

	// Execute Request
	s.RegisterRoutes().ServeHTTP(rr, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rr.Code)
}
