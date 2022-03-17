package server

import "github.com/gin-gonic/gin"

type Handler func(ctx *gin.Context)
