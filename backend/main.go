package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	// "github.com/aws/aws-sdk-go-v2/aws"
	// "github.com/aws/aws-sdk-go-v2/config"
	// "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	. "quizy/model"
)

var db = make(map[string]Quiz)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("quiz", func(c *gin.Context) {
		var quiz Quiz

		// Call BindJSON to bind the received JSON to
		// newAlbum.
		if err := c.BindJSON(&quiz); err != nil {
			fmt.Println("ERROR WHILE PARSING QUIZ DURING CREATION")
			return
		}

		id, er := uuid.NewRandom()
		if er != nil {
			panic(er)
		}
		db[id.String()] = quiz
		c.String(http.StatusOK, "ok")
	})

	// Get quiz list
	r.GET("/quiz", func(c *gin.Context) {
		c.JSON(http.StatusOK, db)
	})

	// Get user value
	r.GET("/quiz/:quizId", func(c *gin.Context) {
		quizId := c.Params.ByName("name")
		value, ok := db[quizId]
		if ok {
			c.JSON(http.StatusOK, gin.H{"quiz": quizId, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"quiz": quizId, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	// //}))
	// authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"foo":  "bar", // user:foo password:bar
	// 	"manu": "123", // user:manu password:123
	// }))

	// /* example curl for /admin with basicauth header
	//    Zm9vOmJhcg== is base64("foo:bar")

	// 	curl -X POST \
	//   	http://localhost:8080/admin \
	//   	-H 'authorization: Basic Zm9vOmJhcg==' \
	//   	-H 'content-type: application/json' \
	//   	-d '{"value":"bar"}'
	// */
	// authorized.POST("admin", func(c *gin.Context) {
	// 	// user := c.MustGet(gin.AuthUserKey).(string)

	// 	// Parse JSON
	// 	var json struct {
	// 		Value string `json:"value" binding:"required"`
	// 	}

	// 	if c.Bind(&json) == nil {
	// 		// db[user] = json.Value
	// 		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	// 	}
	// })

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
