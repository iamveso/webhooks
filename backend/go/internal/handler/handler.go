package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/iamveso/webhooks/backend/go/internal/service"
)

type Handler struct {
	service service.WebhooksService
}

func NewHandler(service service.WebhooksService) *Handler {
	return &Handler{
		service: service,
	}
}

type ResponseMeta struct {
	RequestID string `json:"request_id"`
}

type Response[T any] struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    T            `json:"data,omitempty"`
	Meta    ResponseMeta `json:"meta"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

func respondJson(w http.ResponseWriter, code int, response any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(response) //intentionally ignore error foe dev ease, ideally return an error here
}

func RespondSuccess(w http.ResponseWriter, r *http.Request, code int, data any) {
	respondJson(w, code, Response[any]{
		Status:  "success",
		Message: "successful request",
		Data:    data,
		Meta: ResponseMeta{
			RequestID: middleware.GetReqID(r.Context()),
		},
	})
}

func RespondFailed(w http.ResponseWriter, r *http.Request, code int, message string) {
	respondJson(w, code, Response[any]{
		Status:  "error",
		Message: message,
		Meta: ResponseMeta{
			RequestID: middleware.GetReqID(r.Context()),
		},
	})
}
func (h *Handler) StartServer() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("endpoint up and running\n"))
	})
	// subscribe to events
	r.Post("/webhooks/subscribe", h.SubscribeToWebhookEvents)

	// create/post a webhook
	r.Post("/webhooks", h.CreateWebhook)

	// get a list of processed webhooks under your username
	r.Get("/webhooks", h.GetWebhookList)

	// get webhook by ID
	r.Get("/webhooks/{id}", h.GetWebhookById)

	return http.ListenAndServe(":8080", r)
}

func (h *Handler) SubscribeToWebhookEvents(w http.ResponseWriter, r *http.Request) {
	RespondSuccess(w, r, http.StatusOK, "great stuff")
}

func (h *Handler) CreateWebhook(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) GetWebhookList(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) GetWebhookById(w http.ResponseWriter, r *http.Request) {}
