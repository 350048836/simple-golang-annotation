package controller

import (
	"../httpserver"
)

// 健康检查接口
// @method=GET
// @path=/health
func HealthCheck(request *httpserver.HttpServletRequest, response *httpserver.HttpServletResponse) {

	response.Ok().String("OK")
}
