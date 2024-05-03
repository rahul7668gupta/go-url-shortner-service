package service_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/rahul7668gupta/go-url-shortner/pkg/constants"
	"github.com/rahul7668gupta/go-url-shortner/pkg/logger"
	"github.com/rahul7668gupta/go-url-shortner/pkg/service"
	"github.com/rahul7668gupta/go-url-shortner/testing/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUrlShortnerService_Redirect(t *testing.T) {
	t.Run("original url for short code found - success", func(t *testing.T) {
		loggerObj := logger.InitLogger()
		repo := new(mocks.IShortnerRepo)
		shortCode := "5Wm"
		urlToShorten := "https://www.amazon.com"
		// setup env
		os.Setenv(constants.SHORT_URL_DOMAIN, "http://localhost:8080/")
		// set up request and reponse objects for testing
		req := httptest.NewRequest("GET", "/r/{code}", nil)
		req = mux.SetURLVars(req, map[string]string{
			"code": shortCode,
		})

		w := httptest.NewRecorder()

		// mock all repo calls
		repo.On("GetOriginalUrlForShortCode", mock.Anything, shortCode).Return(urlToShorten, nil)

		// Intitialise service
		srv := service.NewUrlShortnerService(repo, loggerObj)
		// call the test func
		srv.Redirect(w, req)
		// asserts
		assert.Equal(t, http.StatusMovedPermanently, w.Code)
		assert.Equal(t, urlToShorten, w.Header().Get("Location"))
	})

	t.Run("original url for short code not found - failure", func(t *testing.T) {
		loggerObj := logger.InitLogger()
		repo := new(mocks.IShortnerRepo)
		shortCode := "5Wm"

		// setup env
		os.Setenv(constants.SHORT_URL_DOMAIN, "http://localhost:8080/")
		// set up request and reponse objects for testing
		req := httptest.NewRequest("GET", "/r/{code}", nil)
		req = mux.SetURLVars(req, map[string]string{
			"code": shortCode,
		})

		w := httptest.NewRecorder()

		expectedError := errors.New("some random error")

		// mock all repo calls
		repo.On("GetOriginalUrlForShortCode", mock.Anything, shortCode).Return("", expectedError)

		// Intitialise service
		srv := service.NewUrlShortnerService(repo, loggerObj)
		// call the test func
		srv.Redirect(w, req)
		// asserts
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, "Invalid short url\n", w.Body.String())
	})
}
