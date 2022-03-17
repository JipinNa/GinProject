package server

import "github.com/gin-gonic/gin"

func HelloWorld(ctx *gin.Context) {
	ctx.Writer.Write([]byte("ok"))
}
