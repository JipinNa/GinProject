package route

import (
	"MockUmf/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Route struct {
	name    string
	method  server.Method
	path    string
	handler server.Handler
}

var Routes = []gin.RouteInfo{
	{
		Method:      http.MethodGet,
		Path:        "/v1",
		Handler:     "HelloWorld",
		HandlerFunc: server.HelloWorld,
	},
}
