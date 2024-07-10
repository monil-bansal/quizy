package data

import (
	"fmt"
	"log"
	. "quizy/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	// IS CLASSED AT THE START OF THE PROGRAM
	dbClient = newclient()

	// TODO: study if we should create 2 different db clients or use the same client for different tables. (both will work but want to know what are the pros and cons).
	quizTableName = "quizy_quiz"
	userTableName = "quizy_user"
)

func newclient() *dynamodb.DynamoDB {
	fmt.Println("CREATING DB CLIENT")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	return dynamodb.New(sess)
}

func AddQuiz(quiz Quiz) {
	av, err := dynamodbattribute.MarshalMap(quiz)
	if err != nil {
		log.Fatalf("Got error marshalling new quiz item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(quizTableName),
	}

	_, err = dbClient.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem in AddQuiz: %s", err)
	}
}

func AddUser(user User) {
	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		log.Fatalf("Got error marshalling new user item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(userTableName),
	}

	_, err = dbClient.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem in AddUSer: %s", err)
	}
}

// TODO: Add pagination
func GetAllQuiz(invalidateAnswer bool) []Quiz {
	params := &dynamodb.ScanInput{
		TableName: aws.String(quizTableName),
	}

	result, err := dbClient.Scan(params)
	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}

	quizList := []Quiz{}

	for _, i := range result.Items {
		quiz := Quiz{}
		err = dynamodbattribute.UnmarshalMap(i, &quiz)

		if err != nil {
			fmt.Println(i)
			log.Fatalf("Got error unmarshalling: %s", err)
		}

		if invalidateAnswer {
			quiz = removeAnswer(quiz)
		}

		quizList = append(quizList, quiz)
	}

	return quizList
}

func GetQuiz(quizId string, invalidateAnswer bool) Quiz {
	result, err := dbClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(quizTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"quizId": {
				S: aws.String(quizId),
			},
		},
	})
	if err != nil {
		log.Fatalf("Got error calling GetItem in GetQuiz: %s", err)
	}

	quiz := Quiz{}
	if result.Item == nil {
		msg := "Could not find '" + quizId + "'"
		fmt.Errorf(msg)
		return quiz
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &quiz)
	if err != nil {
		log.Fatalf("Failed to unmarshal Quiz Record, %v", err)
	}

	if invalidateAnswer {
		quiz = removeAnswer(quiz)
	}

	return quiz

}

func GetUser(userId string) *User {
	result, err := dbClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(userTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"userId": {
				S: aws.String(userId),
			},
		},
	})
	if err != nil {
		log.Fatalf("Got error calling GetItem in GetUser: %s", err)
	}
	if result.Item == nil {
		return nil
	}

	user := User{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		log.Fatalf("Failed to unmarshal User Record, %v", err)
	}

	return &user
}

func removeAnswer(quiz Quiz) Quiz {
	questions := []Question{}
	for _, j := range quiz.Questions {
		j.Answer = -1
		questions = append(questions, j)
	}
	quiz.Questions = questions
	return quiz
}
