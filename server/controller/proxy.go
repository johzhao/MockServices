package controller

import (
	"go.uber.org/zap"
	"mock.services/server/service"
)

func NewProxyController(logger *zap.Logger, proxyService service.ProxyService) ProxyController {
	return ProxyController{
		logger:       logger,
		proxyService: proxyService,
	}
}

type ProxyController struct {
	logger       *zap.Logger
	proxyService service.ProxyService
}
