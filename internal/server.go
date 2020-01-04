package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Raybird/whale/internal/models"
	"github.com/Raybird/whale/internal/modules"
	"github.com/Raybird/whale/internal/routes"
	"github.com/Raybird/whale/internal/seed"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	ginprometheus "github.com/zsais/go-gin-prometheus"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var server = Server{}

// Run ...
func Run() {
	var err error
	err = godotenv.Load()

	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))

	seed.Load(server.DB)

	server.Run(":8080")
}

// Server ...
type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

// Initialize ...
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}
	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
			DbHost, DbUser, DbName, DbPassword)
		log.Println(DBURL)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}) //database migration

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	// prometheus
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	base := modules.BaseCtrl{
		DB:     server.DB,
		Router: r,
	}
	routes.InitializeRoutes(&base)
	server.Router = r

}

// Run ...
func (server *Server) Run(addr string) {

	srv := &http.Server{
		Addr:    addr, // listen and serve on 0.0.0.0:8080
		Handler: server.Router,
	}

	fmt.Println("Listening to port: " + addr)
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		fmt.Println("Listening to port 8080")
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	s := <-quit
	log.Printf("Shutdown Server: %s ...", s)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
