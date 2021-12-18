package transport

import (
	"mock.services/common/router"
	"mock.services/server/controller"
)

func proxyRouters(proxyController controller.ProxyController) []router.Router {
	return []router.Router{
		//router.NewJSONRouter(http.MethodPost, "/users", command.CreateUserCommand{}, proxyController.CreateUser),
		//router.NewJSONRouter(http.MethodPut, "/users/:user_id", command.UpdateUserCommand{}, proxyController.UpdateUser),
		//router.NewJSONRouter(http.MethodGet, "/users", query.GetUserQuery{}, proxyController.GetUser),
		//router.NewJSONRouter(http.MethodGet, "/users/fail", nil, proxyController.FailedTest),
	}
}
