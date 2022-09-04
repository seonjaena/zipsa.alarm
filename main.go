package zipsa_alarm

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"log"
	"time"
)

type FirestoreEvent struct {
	OldValue   FirestoreValue `json:"oldValue"`
	Value      FirestoreValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

type FirestoreValue struct {
	CreateTime time.Time     `json:"createTime"`
	Fields     FirestoreData `json:"fields"`
	Name       string        `json:"name"`
	UpdateTime time.Time     `json:"updateTime"`
}

type FirestoreData struct {
	Title struct {
		StringValue string `json:"stringValue"`
	} `json:"title"`
	Body struct {
		StringValue string `json:"stringValue"`
	} `json:"body"`
}

func Main(ctx context.Context, e FirestoreEvent) error {
	createTime := e.Value.CreateTime
	updateTime := e.Value.UpdateTime
	titleValue := e.Value.Fields.Title.StringValue
	bodyValue := e.Value.Fields.Body.StringValue

	fmt.Println("CreateTime= ", createTime)
	fmt.Println("UpdateTime= ", updateTime)
	fmt.Println("Title= ", titleValue)
	fmt.Println("Body= ", bodyValue)

	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error Initializing App: %v\n", err)
		return err
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("Error Getting Messaging Client: %v\n", err)
		return err
	}

	topic := "highScores"

	message := &messaging.Message{
		Data: map[string]string{
			"title":      titleValue,
			"body":       bodyValue,
			"createTime": createTime.Format("2006-01-02 15:04:05"),
			"updateTime": updateTime.Format("2006-01-02 15:04:05"),
		},
		Topic: topic,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)

	fmt.Println("response = " + response + "\ntitle = " + message.Data["title"] + "\nbody = " + message.Data["body"] + "\ncreateTime = " + message.Data["createTime"] + "\nupdateTime = " + message.Data["updateTime"] + "\ntopic = " + message.Topic)

	return nil
}
