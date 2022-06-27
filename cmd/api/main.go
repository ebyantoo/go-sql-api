package main

import (
	"go-sql-api/internal/database"
	"go-sql-api/internal/exercise"
	"go-sql-api/internal/middleware"
	"go-sql-api/internal/user"

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
	route.POST("/exercises/:exerciseId/questions", middleware.Authentication(userService), exerciseService.CreateQuestions)
	route.POST("/exercises/:exerciseId/questions/:questionId/answer", middleware.Authentication(userService), exerciseService.CreateAnswer)

	// route user
	route.POST("/register", userService.Register)
	route.POST("/login", userService.Login)

	// listen and serve on 0.0.0.0:8080
	route.Run()
}
