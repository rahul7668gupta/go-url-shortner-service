package service_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/rahul7668gupta/go-url-shortner/pkg/constants"
	"github.com/rahul7668gupta/go-url-shortner/pkg/dto"
	"github.com/rahul7668gupta/go-url-shortner/pkg/logger"
	"github.com/rahul7668gupta/go-url-shortner/pkg/service"
	"github.com/rahul7668gupta/go-url-shortner/pkg/utils"
	"github.com/rahul7668gupta/go-url-shortner/testing/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUrlShortnerService_ShortenURL(t *testing.T) {
	t.Run("already shortened url found - success", func(t *testing.T) {
		loggerObj := logger.InitLogger()
		repo := new(mocks.IShortnerRepo)
		shortCode := "abc"
		urlToShorten := "https://www.amazon.com"
		jsonData, err := json.Marshal(dto.Request{URL: urlToShorten})
		if err != nil {
			t.Fatal(err)
		}

		// setup env
		os.Setenv(constants.SHORT_URL_DOMAIN, "http://localhost:8080/")
		// set up request and reponse objects for testing
		req := httptest.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		// mock all repo calls
		repo.On("LookupURL", mock.Anything, urlToShorten).Return(shortCode, true)

		// Intitialise service
		srv := service.NewUrlShortnerService(repo, loggerObj)
		// call the test func
		srv.ShortenURL(w, req)
		// asserts
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		var response dto.Reponse
		err = json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEmpty(t, response.ShortenedUrl)
		assert.Equal(t, "http://localhost:8080/r/abc", response.ShortenedUrl)
	})

	t.Run("short url created - success", func(t *testing.T) {
		loggerObj := logger.InitLogger()
		repo := new(mocks.IShortnerRepo)
		shortCode := utils.GetShortCodeFromId(1)
		urlToShorten := "https://www.amazon.com"
		domain, _ := utils.GetUrlDomain(urlToShorten)
		jsonData, err := json.Marshal(dto.Request{URL: urlToShorten})
		if err != nil {
			t.Fatal(err)
		}

		// setup env
		os.Setenv(constants.SHORT_URL_DOMAIN, "http://localhost:8080/")
		// set up request and reponse objects for testing
		req := httptest.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		// mock all repo calls
		repo.On("LookupURL", mock.Anything, urlToShorten).Return("", false)
		repo.On("IncrementCounterForShortCode", mock.Anything).Return(int64(1), nil)
		repo.On("CreateShortCodeRecord", mock.Anything, "1", urlToShorten).Return(nil)
		repo.On("CreateIndexOnOriginalUrl", mock.Anything, urlToShorten, shortCode).Return(nil)
		repo.On("IncrementDomainCounter", mock.Anything, domain).Return(nil)

		// Intitialise service
		srv := service.NewUrlShortnerService(repo, loggerObj)
		// call the test func
		srv.ShortenURL(w, req)
		// asserts
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		var response dto.Reponse
		err = json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEmpty(t, response.ShortenedUrl)
		assert.Equal(t, "http://localhost:8080/r/"+shortCode, response.ShortenedUrl)
	})

	t.Run("error at IncrementCounterForShortCode - failure", func(t *testing.T) {
		loggerObj := logger.InitLogger()
		repo := new(mocks.IShortnerRepo)
		urlToShorten := "https://www.amazon.com"
		jsonData, err := json.Marshal(dto.Request{URL: urlToShorten})
		if err != nil {
			t.Fatal(err)
		}

		// setup env
		os.Setenv(constants.SHORT_URL_DOMAIN, "http://localhost:8080/")
		// set up request and reponse objects for testing
		req := httptest.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		expectedError := errors.New("some random error")

		// mock all repo calls
		repo.On("LookupURL", mock.Anything, urlToShorten).Return("", false)
		repo.On("IncrementCounterForShortCode", mock.Anything).Return(int64(1), expectedError)

		// Intitialise service
		srv := service.NewUrlShortnerService(repo, loggerObj)
		// call the test func
		srv.ShortenURL(w, req)
		// asserts
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "Error getting short code counter\n", w.Body.String())
	})

	t.Run("error at CreateShortCodeRecord - failure", func(t *testing.T) {
		loggerObj := logger.InitLogger()
		repo := new(mocks.IShortnerRepo)
		urlToShorten := "https://www.amazon.com"
		jsonData, err := json.Marshal(dto.Request{URL: urlToShorten})
		if err != nil {
			t.Fatal(err)
		}

		// setup env
		os.Setenv(constants.SHORT_URL_DOMAIN, "http://localhost:8080/")
		// set up request and reponse objects for testing
		req := httptest.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		expectedError := errors.New("some random error")

		// mock all repo calls
		repo.On("LookupURL", mock.Anything, urlToShorten).Return("", false)
		repo.On("IncrementCounterForShortCode", mock.Anything).Return(int64(1), nil)
		repo.On("CreateShortCodeRecord", mock.Anything, "1", urlToShorten).Return(expectedError)

		// Intitialise service
		srv := service.NewUrlShortnerService(repo, loggerObj)
		// call the test func
		srv.ShortenURL(w, req)
		// asserts
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "Error saving shortner record\n", w.Body.String())
	})

	t.Run("error at CreateIndexOnOriginalUrl - failure", func(t *testing.T) {
		loggerObj := logger.InitLogger()
		repo := new(mocks.IShortnerRepo)
		shortCode := "1"
		urlToShorten := "https://www.amazon.com"
		jsonData, err := json.Marshal(dto.Request{URL: urlToShorten})
		if err != nil {
			t.Fatal(err)
		}

		// setup env
		os.Setenv(constants.SHORT_URL_DOMAIN, "http://localhost:8080/")
		// set up request and reponse objects for testing
		req := httptest.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		expectedError := errors.New("some random error")

		// mock all repo calls
		repo.On("LookupURL", mock.Anything, urlToShorten).Return("", false)
		repo.On("IncrementCounterForShortCode", mock.Anything).Return(int64(1), nil)
		repo.On("CreateShortCodeRecord", mock.Anything, "1", urlToShorten).Return(nil)
		repo.On("CreateIndexOnOriginalUrl", mock.Anything, urlToShorten, shortCode).Return(expectedError)

		// Intitialise service
		srv := service.NewUrlShortnerService(repo, loggerObj)
		// call the test func
		srv.ShortenURL(w, req)
		// asserts
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "Error creating url index\n", w.Body.String())
	})

}
