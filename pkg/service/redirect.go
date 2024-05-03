package service

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *UrlShortnerService) Redirect(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	vars := mux.Vars(r)
	shortCode := vars["code"]
	orignalUrl, err := s.repo.GetOriginalUrlForShortCode(ctx, shortCode)
	if err != nil {
		s.logger.Sugar().Errorf("short url not found %s, status code %d", err.Error(), http.StatusNotFound)
		http.Error(w, "Invalid short url", http.StatusNotFound)
		return
	}
	s.logger.Sugar().Infof("redirecting short code %s to original url %s", shortCode, orignalUrl)
	http.Redirect(w, r, orignalUrl, http.StatusMovedPermanently)
}
