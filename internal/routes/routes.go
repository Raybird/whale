package routes

import (
	"net/http"

	"github.com/Raybird/whale/internal/controllers"
	"github.com/Raybird/whale/internal/controllers/home"
	"github.com/Raybird/whale/internal/controllers/login"
	"github.com/Raybird/whale/internal/controllers/users"
	"github.com/gin-gonic/gin"
)

// InitializeRoutes ...
func InitializeRoutes(s *controllers.Server) {

	s.Router.GET("/healthz", Heartbeat)

	homeCtl := home.Ctl{Server: s}
	s.Router.GET("/", homeCtl.Home)

	loginCtl := login.Ctl{Server: s}
	s.Router.POST("/login", loginCtl.Login)

	userCtl := users.Ctl{Server: s}
	s.Router.POST("/users", userCtl.CreateUser)
	s.Router.GET("/users", userCtl.GetUsers)
}

// Heartbeat ...
func Heartbeat(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
