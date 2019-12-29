package routes

import (
	"github.com/Raybird/whale/internal/controllers"
	"github.com/Raybird/whale/internal/controllers/home"
	"github.com/Raybird/whale/internal/controllers/login"
	"github.com/Raybird/whale/internal/controllers/users"
)

// InitializeRoutes ...
func InitializeRoutes(s *controllers.Server) {

	homeCtl := home.Ctl{Server: s}
	s.Router.GET("/", homeCtl.Home)

	loginCtl := login.Ctl{Server: s}
	s.Router.POST("/login", loginCtl.Login)

	userCtl := users.Ctl{Server: s}
	s.Router.POST("/users", userCtl.CreateUser)
	s.Router.GET("/users", userCtl.GetUsers)
}
