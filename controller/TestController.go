package controller

import "../amyxframework/httpserver"

// @method=GET
// @path=/test/get
func GetTest(request *httpserver.HttpServletRequest, response *httpserver.HttpServletResponse) {

	response.Ok().String("GetTest Finished")
}

// @method=GET
// @path=/test/post
func PostTest(request *httpserver.HttpServletRequest, response *httpserver.HttpServletResponse) {

	response.Ok().String("PostTest Finished")
}
