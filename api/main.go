package main

import (
	// "fmt"
	"log"
	// "net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var initialized = false
var ginLambda *ginadapter.GinLambda

func setupRouter() *gin.Engine {
	router := gin.Default()

	// the base path is dynamic, so we have to match pretty much anything here
	routerGroup := router.Group("/:basePath")
	{
		// curl "${ServiceEndpoint}/ping"
		routerGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	return router
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if !initialized {
		// stdout and stderr are sent to AWS CloudWatch Logs
		log.Printf("Gin cold start")
		ginLambda = ginadapter.New(setupRouter())
		initialized = true
	}

	return ginLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
