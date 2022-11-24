package api

import (
	"github.com/gin-gonic/gin"
)

func InitAccountController(server *Server, router *gin.Engine) {
	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccountById)
	router.DELETE("/account/:id", server.deleteAccount)
	router.GET("/accounts", server.listAccounts)
}
