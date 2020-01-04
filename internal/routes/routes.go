package routes

import (
	"net/http"

	"github.com/Raybird/whale/internal/modules"
	"github.com/Raybird/whale/internal/modules/home"
	"github.com/Raybird/whale/internal/modules/login"
	"github.com/Raybird/whale/internal/modules/users"
	"github.com/gin-gonic/gin"
)

// InitializeRoutes ...
func InitializeRoutes(base *modules.BaseCtrl) {

	base.Router.GET("/healthz", Heartbeat)

	homeCtl := home.Ctrl{Base: base}
	base.Router.GET("/", homeCtl.Home)

	loginCtl := login.Ctrl{Base: base}
	base.Router.POST("/login", loginCtl.Login)

	userCtl := users.Ctrl{Base: base}
	base.Router.POST("/users", userCtl.CreateUser)
	base.Router.GET("/users", userCtl.GetUsers)
}

// Heartbeat ...
func Heartbeat(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
