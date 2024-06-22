package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	. "quizy/data"
	. "quizy/login"
	. "quizy/model"
)

func getRandomId() string {
	id, er := uuid.NewRandom()
	if er != nil {
		panic(er)
	}
	return id.String()
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println(" I AM HERE")
		// log.Fatal("HI MONIL")
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
		quiz.QuizId = getRandomId()

		AddQuiz(quiz)
		c.String(http.StatusOK, "ok")
	})

	// Get quiz list
	r.GET("/quiz", func(c *gin.Context) {
		quizList := GetAllQuiz(true /* invalidateAnswer */)
		c.JSON(http.StatusOK, quizList)
	})

	// Get user value
	r.GET("/quiz/:quizId", func(c *gin.Context) {
		quizId := c.Params.ByName("quizId")
		quiz := GetQuiz(quizId, true /* invalidateAnswer */)
		c.JSON(http.StatusOK, quiz)
	})

	r.POST("/submit/:quizId", func(c *gin.Context) {
		quizId := c.Params.ByName("quizId")
		originalQuiz := GetQuiz(quizId, false /* invalidateAnswer */)

		var submittedQuiz Quiz

		// Call BindJSON to bind the received JSON to
		// newAlbum.
		if err := c.BindJSON(&submittedQuiz); err != nil {
			// TODO: replace fmt with log.fatal here and elsewhere
			fmt.Println("ERROR WHILE PARSING QUIZ DURING CREATION")
			return
		}

		/*
			TODO:
				1. tell which questions were answered correctly
				2. Persist the data about submission
		*/
		score := 0
		for i := 0; i < len(originalQuiz.Questions); i++ {
			if originalQuiz.Questions[i].Answer == submittedQuiz.Questions[i].Answer {
				score++
			}
		}

		c.String(http.StatusOK, strconv.Itoa(score))
	})

	r.POST("/createUser", func(c *gin.Context) {
		var user User

		// Call BindJSON to bind the received JSON to
		// newAlbum.
		if err := c.BindJSON(&user); err != nil {
			fmt.Println("ERROR WHILE PARSING QUIZ DURING CREATION")
			return
		}

		password := user.Password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		user.Password = string(hashedPassword)

		// TODO: Check if user already exists if needed -> if dynamoDb already doesn't handle it as it is also the primary key.
		// curUser := GetUser(user.Email)

		// if curUser != nil {
		// 	c.String(http.StatusConflict, "user with email already exists")
		// }

		AddUser(user)
	})

	r.POST("/login", func(c *gin.Context) {
		var user User

		// Call BindJSON to bind the received JSON to
		// newAlbum.
		if err := c.BindJSON(&user); err != nil {
			fmt.Println("ERROR WHILE PARSING QUIZ DURING CREATION")
			return
		}

		password := user.Password

		existingUser := GetUser(user.Email)
		if existingUser == nil {
			c.String(http.StatusBadRequest, "Error login in. Please check the credentials.")
			return
		}
		hashedPassword := existingUser.Password

		err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		if err != nil {
			c.String(http.StatusBadRequest, "Error login in. Please check the credentials.")
			return
		}

		token, err := CreateToken(existingUser.Email)

		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		// TODO: use JWT for keeping user logged in.
		c.String(http.StatusOK, token)
	})

	/*
		NOTE TO SELF: keeping it for studying later (maybe). Code from documentation
	*/
	// /*
	// 	Authorized group (uses gin.BasicAuth() middleware)
	// 	Same than:
	// 	authorized := r.Group("/")
	// 	authorized.Use(gin.BasicAuth(gin.Credentials{
	// 		"foo":  "bar",
	// 		"manu": "123",
	// 	}))
	// */
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
	// 	user := c.MustGet(gin.AuthUserKey).(string)

	// 	// Parse JSON
	// 	var json struct {
	// 		Value string `json:"value" binding:"required"`
	// 	}

	// 	if c.Bind(&json) == nil {
	// 		db[user] = json.Value
	// 		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	// 	}
	// })

	return r
}

func main() {
	r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	// r.Run(":8080")

	_ = r.RunTLS(":443", "cert.pem", "key.pem")
}
