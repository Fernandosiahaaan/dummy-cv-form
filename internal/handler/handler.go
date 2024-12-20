package handler

import (
	"context"
	"dummy-cv-form/infrastructure/redis"
	"dummy-cv-form/internal/repository"
	"dummy-cv-form/internal/service"
)

type ParamHandler struct {
	Ctx     context.Context
	Redis   *redis.Redis
	Repo    *repository.Repository
	Service *service.Service
}

type Handler struct {
	Ctx     context.Context
	cancel  context.CancelFunc
	service *service.Service
	repo    *repository.Repository
	Redis   *redis.Redis
}

func NewHandler(param *ParamHandler) *Handler {
	handlerCtx, handlerCancel := context.WithCancel(param.Ctx)
	return &Handler{
		Ctx:     handlerCtx,
		cancel:  handlerCancel,
		Redis:   param.Redis,
		repo:    param.Repo,
		service: param.Service,
	}
}

func (h *Handler) Close() {
	h.cancel()
}
