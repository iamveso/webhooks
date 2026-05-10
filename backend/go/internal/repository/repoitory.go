package repository

type WebhooksRepository interface{}

type WebhookRepository struct{}

func NewRepository() WebhooksRepository {
	return &WebhookRepository{}
}
