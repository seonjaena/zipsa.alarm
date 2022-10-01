package zipsa_alarm

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"time"
	"zipsa.alarm/zlog"
)

var log = zlog.Instance()

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
	Topic struct {
		StringValue string `json:"stringValue"`
	} `json:"topic"`
	Token struct {
		StringValue string `json:"stringValue"`
	} `json:"token"`
}

func Main(ctx context.Context, e FirestoreEvent) error {
	createTime := e.Value.CreateTime.Format("2006-01-02 15:04:05")
	updateTime := e.Value.UpdateTime.Format("2006-01-02 15:04:05")
	titleValue := e.Value.Fields.Title.StringValue
	bodyValue := e.Value.Fields.Body.StringValue
	topic := e.Value.Fields.Topic.StringValue
	token := e.Value.Fields.Token.StringValue

	log.Infof("CreateTime = %s", createTime)
	log.Infof("UpdateTime = %s", updateTime)
	log.Infof("Title = %s", titleValue)
	log.Infof("Body = %s", bodyValue)
	log.Infof("Topic = %s", topic)
	log.Infof("Token = %s", token)

	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Errorf("Error Initializing App: %s", err.Error())
		return err
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Errorf("Error Getting Messaging Client: %s", err.Error())
		return err
	}

	var message *messaging.Message
	if topic != "" {
		message = &messaging.Message{
			Data: map[string]string{
				"title":      titleValue,
				"body":       bodyValue,
				"createTime": createTime,
				"updateTime": updateTime,
			},
			Topic: topic,
		}
	} else {
		message = &messaging.Message{
			Data: map[string]string{
				"score": "850",
				"time":  "2:45",
			},
			Token: token,
		}
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Errorf("Error Sending Alarm: %s", err.Error())
		return err
	}
	// Response is a message ID string.
	log.Infof("Successfully sent message = %s", response)

	log.Infof("response = %s", response)
	log.Infof("title = %s", message.Data["title"])
	log.Infof("body = %s", message.Data["body"])
	log.Infof("createTime = %s", message.Data["createTime"])
	log.Infof("updateTime = %s", message.Data["updateTime"])
	log.Infof("topic = %s", topic)

	return nil
}
