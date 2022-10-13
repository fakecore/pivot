package sys

import (
	"github.com/gin-gonic/gin"
)

type UserInterface interface {
	CreateUser(ctx *gin.Context)
}
