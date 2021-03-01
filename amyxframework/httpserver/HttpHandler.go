package httpserver

import (
	"../utils/logger"
	"net/http"
)

var rootRouter *HttpServletRouter

func setRouter(router *HttpServletRouter) {
	rootRouter = router
}

func DefaultHandler(response http.ResponseWriter, request *http.Request) {
	logger.Info("%s", request.Method, request.RequestURI)

	httpServletRequest := HttpServletRequest{
		httpRequest: request,
	}
	httpServletResponse := HttpServletResponse{}

	//交给路由节点执行并获取结果
	rootRouter.execute(&httpServletRequest, &httpServletResponse)

	if len(httpServletResponse.contentType) > 0 {
		response.Header().Set("Content-Type", httpServletResponse.contentType)
	}
	if httpServletResponse.headers != nil {
		for key, value := range httpServletResponse.headers {
			response.Header().Set(key, value)
		}
	}
	response.WriteHeader(httpServletResponse.status)
	response.Write(httpServletResponse.body)
}
