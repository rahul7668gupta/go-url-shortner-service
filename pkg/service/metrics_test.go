package service_test

import (
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
	"github.com/rahul7668gupta/go-url-shortner/testing/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUrlShortnerService_Metrics(t *testing.T) {
	t.Run("metrics found - success", func(t *testing.T) {
		loggerObj := logger.InitLogger()
		repo := new(mocks.IShortnerRepo)

		// setup env
		os.Setenv(constants.SHORT_URL_DOMAIN, "http://localhost:8080/")
		// set up request and reponse objects for testing
		req := httptest.NewRequest("GET", "/metrics", nil)
		w := httptest.NewRecorder()

		expectedDomains := []dto.Metrics{
			{
				DomainName: "https://www.infracloud.io",
				Count:      5,
			},
			{
				DomainName: "https://www.udemy.com",
				Count:      3,
			},
			{
				DomainName: "https://www.amazon.com",
				Count:      1,
			},
		}

		// mock all repo calls
		repo.On("GetMetrics", mock.Anything).Return(expectedDomains, nil)

		// Intitialise service
		srv := service.NewUrlShortnerService(repo, loggerObj)
		// call the test func
		srv.Metrics(w, req)
		// asserts
		assert.Equal(t, http.StatusOK, w.Code)
		var metricsRaw json.RawMessage
		err := json.NewDecoder(w.Body).Decode(&metricsRaw)
		if err != nil {
			t.Fatal(err)
		}
		var metrics []dto.Metrics
		err = json.Unmarshal(metricsRaw, &metrics)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEmpty(t, metrics)
		assert.Equal(t, expectedDomains, metrics)
	})

	t.Run("metrics repo error - failure", func(t *testing.T) {
		loggerObj := logger.InitLogger()
		repo := new(mocks.IShortnerRepo)

		// setup env
		os.Setenv(constants.SHORT_URL_DOMAIN, "http://localhost:8080/")
		// set up request and reponse objects for testing
		req := httptest.NewRequest("GET", "/r/{code}", nil)

		w := httptest.NewRecorder()

		expectedError := errors.New("some random error")

		// mock all repo calls
		repo.On("GetMetrics", mock.Anything).Return(nil, expectedError)

		// Intitialise service
		srv := service.NewUrlShortnerService(repo, loggerObj)
		// call the test func
		srv.Metrics(w, req)
		// asserts
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "Error retrieving metrics\n", w.Body.String())
	})
}
