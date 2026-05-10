package service

import "github.com/iamveso/webhooks/backend/go/internal/repository"

type WebhooksService interface{}

type WebhookService struct {
	repo repository.WebhooksRepository
}

func NewService(repo repository.WebhooksRepository) WebhooksService {
	return &WebhookService{
		repo: repo,
	}
}
