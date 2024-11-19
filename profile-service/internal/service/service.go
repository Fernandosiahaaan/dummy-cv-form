package service

import (
	"context"
	"profiles-service/infrastructure/redis"
	"profiles-service/internal/repository"
)

type ServiceParam struct {
	Repo  *repository.Repository
	Redis *redis.Redis
	Ctx   context.Context
}

type Service struct {
	repo   *repository.Repository
	redis  *redis.Redis
	ctx    context.Context
	cancel context.CancelFunc
}

func NewService(param ServiceParam) *Service {
	serviceCtx, serviceCancel := context.WithCancel(param.Ctx)
	return &Service{
		ctx:    serviceCtx,
		cancel: serviceCancel,
		repo:   param.Repo,
		redis:  param.Redis,
	}
}

func (s *Service) Close() {
	s.cancel()
}
