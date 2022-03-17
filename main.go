package main

import (
	"MockUmf/route"
	"MockUmf/server"
	"github.com/gin-gonic/gin"
	"reflect"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	for _, route := range route.Routes {
		v := reflect.ValueOf(router)
		vMethod := v.MethodByName(route.Method)
		args := []reflect.Value{reflect.ValueOf(route.Path), reflect.ValueOf(route.HandlerFunc)}
		vMethod.Call(args)
		/*switch route.Method {
		case http.MethodGet:
			router.GET(route.Path, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Path, route.HandlerFunc)
		case http.MethodPatch:

		}*/
	}
	router.Run(server.ListenAddr)
}
