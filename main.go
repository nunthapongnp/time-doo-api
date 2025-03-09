package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nunthapongnp/time-doo-api/config"
	"github.com/nunthapongnp/time-doo-api/internal/controllers"
	"github.com/nunthapongnp/time-doo-api/internal/repositories"
	"github.com/nunthapongnp/time-doo-api/internal/services"
	"github.com/nunthapongnp/time-doo-api/routes"
)

func main() {
	godotenv.Load()
	env := config.LoadEnv()

	// Firebase/Firestore
	fsClient, authClient, err := config.InitFirebase(env)
	if err != nil {
		log.Fatalf("Firebase init error: %v", err)
	}

	// Redis
	redisClient := config.InitRedis(env)
	redisClient.FlushAllAsync(redisClient.Context())

	authController := controllers.NewAuthController(authClient)

	taskRepo := repositories.NewTaskRepository(fsClient, redisClient)
	taskService := services.NewTaskService(taskRepo)
	taskController := controllers.NewTaskController(taskService)

	subTaskRepo := repositories.NewSubTaskRepository(fsClient, redisClient)
	subTaskService := services.NewSubTaskService(subTaskRepo)
	subTaskController := controllers.NewSubTaskController(subTaskService)

	router := gin.Default()

	// Public routes
	routes.SetupPublicRoutes(router, authClient, authController)
	routes.SetupApiRoutes(router, authClient, taskController, subTaskController)

	port := env.Port
	if port == "" {
		port = "3000"
	}
	log.Fatal(router.Run(":" + port))
}
