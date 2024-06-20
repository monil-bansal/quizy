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

type TableBasics struct {
	DynamoDbClient *dynamodb.DynamoDB
	TableName      string
}

// IS CLASSED AT THE START OF THE PROGRAM
var dbClient = newclient()

var tableName = "quizy_quiz"

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
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dbClient.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}
}

// TODO: Add pagination
func GetAllQuiz(invalidateAnswer bool) []Quiz {
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
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
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"quizId": {
				S: aws.String(quizId),
			},
		},
	})
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}
	if result.Item == nil {
		msg := "Could not find '" + quizId + "'"
		log.Fatalln(msg)
	}

	quiz := Quiz{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &quiz)
	if err != nil {
		log.Fatalf("Failed to unmarshal Record, %v", err)
	}

	if invalidateAnswer {
		quiz = removeAnswer(quiz)
	}

	return quiz

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
