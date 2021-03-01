package httpserver

type HttpServletExecute struct {
	Path    string
	Method  string
	ContentType string
	Handler interface{}
}
