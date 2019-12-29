package home

import (
	"net/http"

	"github.com/Raybird/whale/internal/controllers"
	"github.com/gin-gonic/gin"
)

// Ctl ...
type Ctl struct {
	Server *controllers.Server
}

// Home ...
func (ctl *Ctl) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome To This Awesome API",
	})
}
