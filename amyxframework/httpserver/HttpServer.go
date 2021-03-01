package httpserver

import (
	"../utils/logger"
	"../utils/stringutils"
	"net/http"
)

func Start(port string, router *HttpServletRouter) {
	if stringutils.IsEmpty(port) {
		port = "8080"
	}
	logger.Info("Http Server Start On Port: %s", port)
	setRouter(router)
	http.HandleFunc("/", DefaultHandler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil && err != http.ErrServerClosed {
		logger.Error("Http Server Start Failed: %v", err)
	}
	logger.Info("Http Server Exit")
}
