package home

import (
	"net/http"

	"github.com/Raybird/whale/internal/modules"
	"github.com/gin-gonic/gin"
)

// Ctl ...
type Ctrl struct {
	Base *modules.BaseCtrl
}

// Home ...
func (ctl *Ctrl) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome To This Awesome API",
	})
}
