package httpserver

import "net/http"

type HttpServletRequest struct {
	httpRequest *http.Request
}

func (request *HttpServletRequest) GetUrl() string {
	return request.httpRequest.RequestURI
}

func (request *HttpServletRequest) GetMethod() string{
	return request.httpRequest.Method
}