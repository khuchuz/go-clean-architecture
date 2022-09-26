package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	authhttp "github.com/khuchuz/go-clean-architecture/auth/delivery"
	itface "github.com/khuchuz/go-clean-architecture/auth/itface"
	authmongo "github.com/khuchuz/go-clean-architecture/auth/repository"
	authusecase "github.com/khuchuz/go-clean-architecture/auth/usecase"
)

func main() {

	app := NewApp()

	if err := app.Run("8000"); err != nil {
		log.Fatalf("%s", err.Error())
	}
}

type App struct {
	httpServer *http.Server
	authUC     itface.UseCase
}

func NewApp() *App {
	db := initDB()

	userRepo := authmongo.NewUserRepository(db, "users")

	return &App{
		authUC: authusecase.NewAuthUseCase(
			userRepo,
			"hash_salt",
			[]byte("signing_key"),
			86400,
		),
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Set up http handlers
	h := authhttp.NewHandler(a.authUC)

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/sign-up", h.SignUp)
		authEndpoints.POST("/sign-in", h.SignIn)
		authEndpoints.POST("/change-pass", h.ChangePassword)
	}

	// API endpoints
	authMiddleware := authhttp.NewAuthMiddleware(a.authUC)
	_ = router.Group("/api", authMiddleware)

	// HTTP Server
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

func initDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("testdb")
}
