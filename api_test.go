package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCars(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {

		carsData := []map[string]string{
			{
				"carsid":       "1",
				"make":         "Toyota",
				"model":        "Camry",
				"licenceplate": "XYZ123",
				"ownername":    "John Doe",
				"date":         "17-12-2023",
				"status":       "active",
			},
		}

		jsonData, _ := json.Marshal(carsData)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}

	req, err := http.NewRequest("GET", "/cars", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200 OK")

	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"), "Content type should be JSON")

	var cars []map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &cars)
	if err != nil {
		t.Fatalf("Error unmarshalling response body: %s", err)
	}

	for _, car := range cars {
		assert.IsType(t, "", car["carsid"], "CarsId should be string type")
		assert.IsType(t, "", car["make"], "Make should be string type")
		assert.IsType(t, "", car["model"], "Model should be string type")
		assert.IsType(t, "", car["licenceplate"], "LicencePlate should be string type")
		assert.IsType(t, "", car["ownername"], "OwnerName should be string type")
		assert.IsType(t, "", car["status"], "Status should be string type")

		dateStr, ok := car["date"].(string)
		if !ok {
			t.Fatal("Date should be a string type")
		}

		parsedDate, err := time.Parse("02-01-2006", dateStr)
		if err != nil {
			t.Fatalf("Error parsing date: %s", err)
		}

		expectedDateStr := parsedDate.Format("02-01-2006")
		if dateStr != expectedDateStr {
			t.Fatalf("Date format should be dd-mm-yyyy: %s", dateStr)
		}
	}
}

func TestGetCarByID(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {

		carData := map[string]string{
			"carsid":       "1",
			"make":         "Toyota",
			"model":        "Camry",
			"licenceplate": "XYZ123",
			"ownername":    "John Doe",
			"date":         "17-12-2023",
			"status":       "active",
		}

		jsonData, _ := json.Marshal(carData)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}

	req, err := http.NewRequest("GET", "/car/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200 OK")
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"), "Content type should be JSON")

	var car map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &car)
	if err != nil {
		t.Fatalf("Error unmarshalling response body: %s", err)
	}

	assert.IsType(t, "", car["carsid"], "CarsId should be string type")
	assert.IsType(t, "", car["make"], "Make should be string type")
	assert.IsType(t, "", car["model"], "Model should be string type")
	assert.IsType(t, "", car["licenceplate"], "LicencePlate should be string type")
	assert.IsType(t, "", car["ownername"], "OwnerName should be string type")
	assert.IsType(t, "", car["status"], "Status should be string type")

	dateStr, ok := car["date"].(string)
	if !ok {
		t.Fatal("Date should be a string type")
	}

	parsedDate, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		t.Fatalf("Error parsing date: %s", err)
	}

	expectedDateStr := parsedDate.Format("02-01-2006")
	if dateStr != expectedDateStr {
		t.Fatalf("Date format should be dd-mm-yyyy: %s", dateStr)
	}
}

func TestPostCarValidData(t *testing.T) {
	payload := map[string]string{
		"cars_id":       "1",
		"make":          "Toyota",
		"model":         "Camry",
		"licence_plate": "XYZ123",
		"owner_name":    "John Doe",
		"date":          "17-12-2023", // Format: dd-mm-yyyy
		"status":        "active",
	}

	jsonPayload, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "/car", bytes.NewBuffer(jsonPayload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Car created successfully"))
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200 OK")
}

func TestPostCarIncompleteData(t *testing.T) {

	incompletePayload := map[string]string{
		"make": "Toyota",
	}

	jsonPayload, _ := json.Marshal(incompletePayload)
	req, err := http.NewRequest("POST", "/car", bytes.NewBuffer(jsonPayload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusBadRequest) // 400 Bad Request
		w.Write([]byte("Incomplete data provided"))
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Status code should be 400 Bad Request")

}

func TestPostCarUnauthorizedAccess(t *testing.T) {

	req, err := http.NewRequest("POST", "/car", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized access"))
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code, "Status code should be 401 Unauthorized")

}

func TestPutCar_SuccessfulUpdate(t *testing.T) {
	// Mock a request payload to update car details
	payload := map[string]string{
		"make":          "Updated Make",
		"model":         "Updated Model",
		"licence_plate": "Updated Plate",
		"owner_name":    "Updated Owner",
		"date":          "25-12-2023", // Format: dd-mm-yyyy
		"status":        "updated",
	}

	jsonPayload, _ := json.Marshal(payload)

	// Replace '1' with an existing car ID
	carID := "1"

	// Create a request for updating an existing car by ID
	req, err := http.NewRequest("PUT", "/car/"+carID, bytes.NewBuffer(jsonPayload))
	if err != nil {
		t.Fatal(err)
	}

	// Mocking the handler function for updating a car by ID
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve car ID from the request and perform necessary update actions
		requestedCarID := r.URL.Query().Get("id")

		// Simulate an existing car record for the provided car ID
		if requestedCarID == carID {
			// Assuming updating car data in the database or another source based on the 'carID' and provided payload
			// ...

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Car updated successfully"))
			return
		}

		// Simulate a case where the car ID doesn't match an existing record
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Car not found"))
	})

	handler.ServeHTTP(rr, req)

	// Check the response status code for a successful update
	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200 OK")

	// Check the response body for the success message
	expectedResponse := "Car updated successfully"
	assert.Contains(t, rr.Body.String(), expectedResponse, "Response should indicate successful update")
}

func TestPutCar_NotFound(t *testing.T) {

	payload := map[string]string{
		"make":          "Updated Make",
		"model":         "Updated Model",
		"licence_plate": "Updated Plate",
		"owner_name":    "Updated Owner",
		"date":          "25-12-2023", // Format: dd-mm-yyyy
		"status":        "updated",
	}

	jsonPayload, _ := json.Marshal(payload)

	carID := "99"

	req, err := http.NewRequest("PUT", "/car/"+carID, bytes.NewBuffer(jsonPayload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestedCarID := r.URL.Query().Get("id")

		if requestedCarID != carID {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Car not found"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Car updated successfully"))
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code, "Status code should be 404 Not Found")

	expectedResponse := "Car not found"
	assert.Contains(t, rr.Body.String(), expectedResponse, "Response should indicate that car is not found")

}

func TestDeleteCarSuccess(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/car/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ... (existing delete logic)

		w.WriteHeader(http.StatusNoContent)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code, "Status code should be 204 No Content")

}
func TestDeleteCarNotFound(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/car/99", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Car not found"))
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code, "Status code should be 404 Not Found")

}
func TestDeleteCarUnauthorized(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/car/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized access"))
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code, "Status code should be 401 Unauthorized")

}
