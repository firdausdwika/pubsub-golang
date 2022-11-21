package main

import (
	"flag"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"

	"fmt"
	"os"
)

func main() {
	emailPtr := flag.String("e", "", "The email address of the user subscribing to the topic")
	topicPtr := flag.String("t", "", "The ARN of the topic to which the user subscribes")

	flag.Parse()

	if *emailPtr == "" || *topicPtr == "" {
		fmt.Println("You must supply an email address and topic ARN")
		fmt.Println("Usage: go run SnsSubscribe.go -e EMAIL -t TOPIC-ARN")
		os.Exit(1)
	}

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String("ap-southeast-1")},
		Profile: "dev",
	}))

	svc := sns.New(sess)

	result, err := svc.Subscribe(&sns.SubscribeInput{
		Endpoint:              emailPtr,
		Protocol:              aws.String("email"),
		ReturnSubscriptionArn: aws.Bool(true), // Return the ARN, even if user has yet to confirm
		TopicArn:              topicPtr,
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(*result.SubscriptionArn)
}
