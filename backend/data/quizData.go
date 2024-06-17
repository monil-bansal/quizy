package data

import (
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

var dbClient = newclient()

func newclient() *dynamodb.DynamoDB {
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

	tableName := "quizy_quiz"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dbClient.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}
}

// func GetItem() {
// 	result, err := dbClient.GetItem(&dynamodb.GetItemInput{
// 		TableName: aws.String("quizy_quiz"),
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"quizId": {
// 				N: aws.String("sample1"),
// 			},
// 			"title": {
// 				S: aws.String("movieName"),
// 			},
// 		},
// 	})
// 	if err != nil {
// 		log.Fatalf("Got error calling GetItem: %s", err)
// 	}
// 	fmt.Println(result)
// }
