package transport

import (
	"context"
	"mock.services/common/router"
	"mock.services/common/router/encoder"
	"mock.services/common/server"
	"mock.services/server/controller"
	"net/http"
)

func SetupProxyRouters(server server.Server, proxyController controller.ProxyController) {
	routers := make([]router.Router, 0)
	routers = append(routers, proxyRouters(proxyController)...)

	routers = append(routers, router.NewCustomRouter(http.MethodGet, "/ping", nil, PingHandler, encoder.NewJSONResponseEncoder()))

	for _, serviceRouter := range routers {
		server.HandleRouter(serviceRouter)
	}
}

//goland:noinspection GoUnusedParameter
func PingHandler(ctx context.Context, req interface{}) (interface{}, error) {
	return map[string]string{
		"message": "pong",
	}, nil
}
