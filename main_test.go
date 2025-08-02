package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
func TestAuth(t *testing.T) {
	router := setupTest()

	// Test login
	loginData := map[string]string{"username": "admin", "password": "password"}
	loginResponse := make(map[string]string)
	
}
func TestGetAlbums(t *testing.T) {
	router := setupTest()

	// Test getting all albums
	req, _ := http.NewRequest("GET", "/albums", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestGetAlbumByID(t *testing.T) {
	router := setupTest()

	// Test getting album by ID
	req, _ := http.NewRequest("GET", "/albums/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	
}
func TestCreateAlbum(t *testing.T) {
	router := setupTest()

	// Test creating a new album
	albumData := album{
		Title: "Test Album",
		Artist: "Test Artist",
		Price: 10.99,
	}
}
func TestUpdateAlbum(t *testing.T) {
	router := setupTest()

	// Test updating an album
	albumData := album{
		Title: "Updated Album",
		Artist: "Updated Artist",
		Price: 20.99,
	}
}
func TestDeleteAlbum(t *testing.T) {
	router := setupTest()

	// Test deleting an album
	req, _ := http.NewRequest("DELETE", "/albums/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
