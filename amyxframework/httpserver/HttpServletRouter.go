package httpserver

import (
	"reflect"
	"strings"
)

type HttpServletRouter struct {
	children    map[string]*HttpServletRouter
	executeList []*HttpServletExecute
}

func (router *HttpServletRouter) SaveChild(handler *HttpServletExecute) {
	url := strings.ReplaceAll(handler.Path, "//", "/")
	if strings.HasPrefix(url, "/") {
		url = string([]byte(url)[1:])
	}
	urlBlocks := strings.Split(url, "/")
	router.saveChild(urlBlocks, 0, handler)
}

func (router *HttpServletRouter) saveChild(urlBlocks []string, pos int, handler *HttpServletExecute) {
	block := urlBlocks[pos]

	if strings.HasPrefix(block, "{") && strings.HasSuffix(block, "}") {
		block = "*"
	}

	if router.children == nil {
		router.children = make(map[string]*HttpServletRouter)
	}

	if _, exist := router.children[block]; !exist {
		router.children[block] = &HttpServletRouter{}
	}

	child := router.children[block]
	if pos < len(urlBlocks)-1 {
		child.saveChild(urlBlocks, pos+1, handler)
	} else {
		child.addExecute(handler)
	}
}

func (router *HttpServletRouter) addExecute(handler *HttpServletExecute) {
	if router.executeList == nil {
		router.executeList = []*HttpServletExecute{}
	}
	router.executeList = append(router.executeList, handler)
}

func (router *HttpServletRouter) FindChild(url string) *HttpServletRouter {
	url = strings.ReplaceAll(url, "//", "/")
	if strings.HasPrefix(url, "/") {
		url = string([]byte(url)[1:])
	}
	urlBlocks := strings.Split(url, "/")
	return router.findChild(urlBlocks, 0)
}

func (router *HttpServletRouter) findChild(urlBlocks []string, pos int) *HttpServletRouter {
	if router.children == nil {
		return nil
	}

	block := urlBlocks[pos]
	child, exist := router.children[block]
	if !exist {
		child, exist = router.children["*"]
		if !exist {
			child, exist = router.children["**"]
			return child
		}
	}

	if pos < len(urlBlocks)-1 {
		return child.findChild(urlBlocks, pos+1)
	}
	return child
}

func (router *HttpServletRouter) execute(request *HttpServletRequest, response *HttpServletResponse) {
	child := router.FindChild(request.GetUrl())
	if child != nil && child.executeList != nil {
		for _, execute := range child.executeList {
			if strings.Compare(request.GetMethod(), HttpMethodOptions) == 0 {
				response.Ok()
			} else if strings.Compare(request.GetMethod(), execute.Method) == 0 {
				callFunc(execute.Handler, request, response)
			}
		}
	} else {
		response.Error(HttpStatusNotfound)
	}
}

func callFunc(target interface{}, params ...interface{}) []reflect.Value {
	f := reflect.ValueOf(target)
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result := f.Call(in)
	return result
}
