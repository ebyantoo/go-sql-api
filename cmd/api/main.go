package main

import (
	"go-sql-api/internal/database"
	"go-sql-api/internal/exercise"
	"go-sql-api/internal/middleware"
	"go-sql-api/internal/user"

	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	db := database.NewDatabaseConn()

	// define service
	exerciseService := exercise.NewExerciseService(db)
	userService := user.NewUserService(db)

	// route exercises
	route.GET("/exercises/:exerciseId", middleware.Authentication(userService), exerciseService.GetExercise)
	route.GET("/exercises/:exerciseId/score", middleware.Authentication(userService), exerciseService.GetUserScore)
	route.POST("/exercises", middleware.Authentication(userService), exerciseService.CreateExercise)

	// route user
	route.POST("/register", userService.Register)
	route.POST("/login", userService.Login)

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("Your Application Name"),
		newrelic.ConfigLicense("19df67a2b6761d4516949197e727cd75c01eNRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigInfoLogger(w)
	)

	// listen and serve on 0.0.0.0:8080
	route.Run()
}
