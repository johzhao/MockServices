package service

import "go.uber.org/zap"

type ProxyService interface {
}

func NewProxyService(logger *zap.Logger) ProxyService {
	return &proxyService{
		logger: logger,
	}
}

type proxyService struct {
	logger *zap.Logger
}
