package main

import (
	"context"
	"log"
	"os"

	"github.com/Manas8803/The-PUC-Project__BackEnd/auth-service/main-app/routes"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//	@title			Auth API
//	@version		1.0
//	@description	This is an auth api for an application.

// @BasePath	/api/v1
var ginLambda *ginadapter.GinLambda

func init() {

	mode := os.Getenv("RELEASE_MODE")
	if mode != "prod" {
		return
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")
	//* Passing the router to all user(auth-service) routes.
	routes.UserRoute(api)

	ginLambda = ginadapter.New(router)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("NOT ABLE TO FIND .env FILE..\nContinuing...")
	}
	mode := os.Getenv("RELEASE_MODE")
	if mode == "testing"{
		TestRun()
		return
	}

	lambda.Start(Handler)
}

func TestRun() {
	router := gin.Default()

	api := router.Group("/api/v1")
	router.Use(cors.Default())
	routes.UserRoute(api)
	router.Run("localhost:8000")
}
