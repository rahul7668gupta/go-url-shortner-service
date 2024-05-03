package repository

import "context"

func (r *ShortnerRepo) GetOriginalUrlForShortCode(ctx context.Context, shortCode string) (string, error) {
	originalURL, err := r.rdb.Get(ctx, shortCode).Result()
	return originalURL, err
}
