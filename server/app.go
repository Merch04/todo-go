package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"todo/auth"
	authhttp "todo/auth/delivery/http"
	authpostgres "todo/auth/repository/postgres"
	authusecase "todo/auth/usecase"
	"todo/todo"
	tdhttp "todo/todo/delivery/http"
	todopostgres "todo/todo/repository/postgres"
	todousecase "todo/todo/usecase"
)

type App struct {
	httpServer *http.Server

	authUC auth.UseCase
	todoUC todo.UseCase
}

func NewApp() *App {
	db := InitDB()

	userRepo := authpostgres.NewUserRepository(db)
	todoRepo := todopostgres.NewTodoRepository(db)

	return &App{
		authUC: authusecase.NewAuthUseCase(
			userRepo,
			viper.GetString("auth.hash_salt"),
			[]byte(viper.GetString("auth.singing_key")),
			viper.GetDuration("auth.token_ttl"),
		),
		todoUC: todousecase.NewTodoUseCase(
			todoRepo,
		),
	}

}

func (a *App) Run(port string) error {
	router := gin.Default()
	router.Use(gin.Recovery(), gin.Logger())

	authhttp.RegisterHTTPEndpoints(router, a.authUC)

	authMiddleware := authhttp.NewAuthMiddleware(a.authUC)
	api := router.Group("/api", authMiddleware)

	tdhttp.RegisterHTTPEndpoints(api, a.todoUC)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)

}

func InitDB() *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		viper.GetString("db.host"),
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.dbname"),
		viper.GetInt32("db.port"),
		viper.GetString("db.sslmode"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error to connect to database: %v", err)
	}

	err = db.AutoMigrate(&authpostgres.User{}, &todopostgres.Task{})
	if err != nil {
		log.Fatalf("Error to migrate models: %v", err)
	}
	return db
}
