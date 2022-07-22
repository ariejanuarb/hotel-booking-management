package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"golang-rest-api/app"
	"golang-rest-api/controller"
	"golang-rest-api/helper"
	"golang-rest-api/middleware"
	"golang-rest-api/model/domain"
	"golang-rest-api/repository"
	"golang-rest-api/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/booking-hotel")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	facilityRepository := repository.NewFacilityRepository()
	facilityService := service.NewFacilityService(facilityRepository, db, validate)
	facilityController := controller.NewFacilityController(facilityService)
	router := app.NewRouter(facilityController)

	return middleware.NewAuthMiddleware(router)
}

func truncateFacility(db *sql.DB) {
	db.Exec("TRUNCATE facilityController")
}

func TestCreateFacilitySuccess(t *testing.T) {
	db := setupTestDB()
	truncateFacility(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/facilitas", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateFacilityFailed(t *testing.T) {
	db := setupTestDB()
	truncateFacility(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateFacilitySuccess(t *testing.T) {
	db := setupTestDB()
	truncateFacility(db)

	tx, _ := db.Begin()
	facilityRepository := repository.NewFacilityRepository()
	facility := facilityRepository.Save(context.Background(), tx, domain.Facility{
		Type:        "Gadget",
		Description: "Terjual",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(facility.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, facility.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateFacilityFailed(t *testing.T) {
	db := setupTestDB()
	truncateFacility(db)

	tx, _ := db.Begin()
	facilityRepository := repository.NewFacilityRepository()
	facility := facilityRepository.Save(context.Background(), tx, domain.Facility{
		Type:        "Gadget",
		Description: "Terjual",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/facilitas/"+strconv.Itoa(facility.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestGetFacilitySuccess(t *testing.T) {
	db := setupTestDB()
	truncateFacility(db)

	tx, _ := db.Begin()
	facilityRepository := repository.NewFacilityRepository()
	facility := facilityRepository.Save(context.Background(), tx, domain.Facility{
		Type: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(facility.Id), nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, facility.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, facility.Type, responseBody["data"].(map[string]interface{})["type"])
}

func TestGetFacilityFailed(t *testing.T) {
	db := setupTestDB()
	truncateFacility(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/404", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestFacilitySuccess(t *testing.T) {
	db := setupTestDB()
	truncateFacility(db)

	tx, _ := db.Begin()
	facilityRepository := repository.NewFacilityRepository()
	facility := facilityRepository.Save(context.Background(), tx, domain.Facility{
		Type:        "Gadget",
		Description: "Terjual",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(facility.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateFacility(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/facilitass/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestListFacilitasSuccess(t *testing.T) {
	db := setupTestDB()
	truncateFacility(db)

	tx, _ := db.Begin()
	facilityRepository := repository.NewFacilityRepository()
	facilitas1 := facilityRepository.Save(context.Background(), tx, domain.Facility{
		Type: "Gadget",
	})
	facilitas2 := facilityRepository.Save(context.Background(), tx, domain.Facility{
		Type: "Computer",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/fasilitas", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	fmt.Println(responseBody)

	var facilitas = responseBody["data"].([]interface{})

	facilityResponse1 := facilitas[0].(map[string]interface{})
	facilityResponse2 := facilitas[1].(map[string]interface{})

	assert.Equal(t, facilitas1.Id, int(facilityResponse1["id"].(float64)))
	assert.Equal(t, facilitas1.Type, facilityResponse1["type"])

	assert.Equal(t, facilitas2.Id, int(facilityResponse2["id"].(float64)))
	assert.Equal(t, facilitas2.Type, facilityResponse2["type"])
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateFacility(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/facilitas", nil)
	request.Header.Add("X-API-Key", "SALAH")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}
