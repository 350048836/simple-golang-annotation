package httpserver

import (
	"../utils/jsonutils"
	"../utils/logger"
)

type HttpServletResponse struct {
	status      int
	body        []byte
	contentType string
	headers     map[string]string
}

func (response *HttpServletResponse) Ok() *HttpServletResponse {
	response.status = HttpStatusOk
	return response
}

func (response *HttpServletResponse) Error(status int) *HttpServletResponse {
	response.status = status
	return response
}

func (response *HttpServletResponse) Byte(msg []byte) *HttpServletResponse {
	response.body = msg
	return response
}

func (response *HttpServletResponse) String(msg string) *HttpServletResponse {
	response.body = []byte(msg)
	return response
}

func (response *HttpServletResponse) Object(msg interface{}) *HttpServletResponse {
	data, err := jsonutils.Obj2Json(msg)
	if err != nil {
		logger.Error("HttpServletResponse.Object Error: %v", err)
	}
	response.body = []byte(data)
	response.contentType = HttpContentTypeJson
	return response
}

func (response *HttpServletResponse) Header(key string, value string) *HttpServletResponse {
	if response.headers == nil {
		response.headers = make(map[string]string)
	}
	response.headers[key] = value
	return response
}
