package service

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/rahul7668gupta/go-url-shortner/pkg/constants"
	"github.com/rahul7668gupta/go-url-shortner/pkg/dto"
	"github.com/rahul7668gupta/go-url-shortner/pkg/repository"
	"github.com/rahul7668gupta/go-url-shortner/pkg/utils"
	"go.uber.org/zap"
)

type UrlShortnerService struct {
	repo   repository.IShortnerRepo
	logger *zap.Logger
}

type IUrlShortnerService interface {
	Metrics(w http.ResponseWriter, r *http.Request)
	Redirect(w http.ResponseWriter, r *http.Request)
	ShortenURL(w http.ResponseWriter, r *http.Request)
}

func NewUrlShortnerService(repo repository.IShortnerRepo, logger *zap.Logger) *UrlShortnerService {
	return &UrlShortnerService{
		repo:   repo,
		logger: logger,
	}
}

func (s *UrlShortnerService) ShortenURL(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var request dto.Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		s.logger.Sugar().Errorf("unable to decode request body %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the url is already shortened
	shortCode, urlFound := s.repo.LookupURL(ctx, request.URL)
	if urlFound {
		s.logger.Sugar().Infof("found the short code for url %s, shortCode %s,status code %d", request.URL, shortCode, http.StatusOK)
		response := dto.Reponse{
			ShortenedUrl: os.Getenv(constants.SHORT_URL_DOMAIN) + "r/" + shortCode,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Generate a shortened path using sha256 hash
	shortCodeCounter, err := s.repo.IncrementCounterForShortCode(ctx)
	if err != nil {
		s.logger.Sugar().Errorf("error getting short code counter %s, statusCode %d", err.Error(), http.StatusInternalServerError)
		http.Error(w, "Error getting short code counter", http.StatusInternalServerError)
		return
	}

	// Convert counter value to base62 chars
	shortCode = utils.GetShortCodeFromId(shortCodeCounter)

	// Store the mapping in Redis
	if err := s.repo.CreateShortCodeRecord(ctx, shortCode, request.URL); err != nil {
		s.logger.Sugar().Errorf("error saving shortener record to redis %s statusCode %d", err.Error(), http.StatusInternalServerError)
		http.Error(w, "Error saving shortner record", http.StatusInternalServerError)
		return
	}

	// Parse the url to get the domain
	domain, err := utils.GetUrlDomain(request.URL)
	if err != nil {
		s.logger.Sugar().Errorf("unable to parse domain %s, status %d", err.Error(), http.StatusInternalServerError)
		http.Error(w, "Error parsing received url", http.StatusInternalServerError)
		return
	}

	// Create and index on original url
	if err = s.repo.CreateIndexOnOriginalUrl(ctx, request.URL, shortCode); err != nil {
		s.logger.Sugar().Errorf("error creating URL index %s, status %d", err.Error(), http.StatusInternalServerError)
		http.Error(w, "Error creating url index", http.StatusInternalServerError)
		return
	}

	// Increment the domain counter
	_ = s.repo.IncrementDomainCounter(ctx, domain) //unsure if the error should be handled here or not

	// Return the shortened URL
	response := dto.Reponse{
		ShortenedUrl: os.Getenv(constants.SHORT_URL_DOMAIN) + "r/" + shortCode,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
