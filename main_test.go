package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestGetUrl(t *testing.T) {
	req, err := http.NewRequest("GET", "/url/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "1")

	rr := httptest.NewRecorder()

	repo := MockUrlRepository{urls: map[string]string{
		"1": "https://example.com/short",
	}}

	getUrl(rr, req, &repo)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("Expected status %v, got %v", http.StatusSeeOther, status)
	}
}

func TestCreateUrl(t *testing.T) {
	req, err := http.NewRequest("POST", "/url", strings.NewReader("url=https://example.com"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()

	repo := MockUrlRepository{urls: map[string]string{}}

	createUrl(rr, req, &repo)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Expected status %v, got %v", http.StatusCreated, status)
	}

	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("Expected Content-Type %v, got %v", "application/json", contentType)
	}

	var id string
	err = json.NewDecoder(rr.Body).Decode(&id)
	if err != nil {
		t.Fatal(err)
	}

	if id != "1" {
		t.Errorf("Expected %s, got %s", "1", id)
	}
}

type MockUrlRepository struct {
	urls map[string]string
}

func (m *MockUrlRepository) Find(id string) (string, error) {
	i, ok := m.urls[id]
	if ok {
		return i, nil
	} else {
		return "nil", errors.New("not found")
	}
}

func (m *MockUrlRepository) Create(path string) (string, error) {
	l := strconv.Itoa(len(m.urls) + 1)
	m.urls[l] = path
	return l, nil
}
