package service

import (
	"context"
	"encoding/json"
	"net/http"
)

func (s *UrlShortnerService) Metrics(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	topDomains, err := s.repo.GetMetrics(ctx)
	if err != nil {
		s.logger.Sugar().Errorf("error retrieving metrics %s, statusCode %d", err.Error(), http.StatusInternalServerError)
		http.Error(w, "Error retrieving metrics", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(topDomains)
}
