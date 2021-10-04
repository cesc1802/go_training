package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cesc1802/go_training/internal/dto/requests"
	"github.com/cesc1802/go_training/internal/services"
	"github.com/gin-gonic/gin"
)

type authController struct{}

func (a authController) Routes(r *gin.Engine) {
	r.GET("/login", loginEndpoint)
}

func loginEndpoint(c *gin.Context) {
	loginPayload := dto_request.LoginPayload{
		UserId: c.Query("user_id"),
		Password: c.Query("password"),
	}
	
	loginResult, err := services.Login(loginPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		var sb strings.Builder
		sb.WriteString("Bearer ")
		sb.WriteString(loginResult.Token)
		c.JSON(http.StatusOK, gin.H{"access_token": fmt.Sprint(sb.String())})
	}
}

func AuthController() authController {
	return authController{}
}