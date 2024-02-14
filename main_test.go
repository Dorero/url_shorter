package main

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"url_shorter/model"
)

func TestGetUrl(t *testing.T) {
	url := model.Url{Id: uuid.New(), Path: "http://example.com"}
	mockRepo := &MockUrlRepository{urls: []model.Url{url}}
	str := "/url/" + url.Id.String()
	req, err := http.NewRequest("GET", str, nil)

	req.SetPathValue("id", url.Id.String())
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	getUrl(w, req, mockRepo)
	assert.Equal(t, http.StatusOK, w.Code)

	var result model.Url
	err = json.NewDecoder(w.Body).Decode(&result)

	assert.NoError(t, err)
	assert.Equal(t, url, result)

}

func TestCreateUrl(t *testing.T) {
	path := "http://example.com"
	mockRepo := &MockUrlRepository{urls: []model.Url{}}

	req, err := http.NewRequest("POST", "/url", strings.NewReader("url="+path))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()

	createUrl(w, req, mockRepo)

	assert.Equal(t, http.StatusCreated, w.Code)

	var result model.Url
	err = json.NewDecoder(w.Body).Decode(&result)

	assert.NoError(t, err)
	assert.Equal(t, path, result.Path)
}

type MockUrlRepository struct {
	urls []model.Url
}

func (m *MockUrlRepository) Find(id string) (model.Url, error) {
	for _, v := range m.urls {
		result, parse := uuid.Parse(id)

		if parse != nil {
			return model.Url{}, parse
		}

		if v.Id == result {
			return v, nil
		}
	}

	return model.Url{}, errors.New("Not found")
}

func (m *MockUrlRepository) Create(path string) (model.Url, error) {
	url := model.Url{Id: uuid.New(), Path: path}
	m.urls = append(m.urls, url)
	return url, nil
}
