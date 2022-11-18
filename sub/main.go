package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var ctx = context.Background()

func main() {
	opt, err := redis.ParseURL("redis://default:KwtdpJJ5W7l09hf1ak0VDO6j6gDfdHhd@redis-16342.c295.ap-southeast-1-1.ec2.cloud.redislabs.com:16342")
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)
	subscriber := rdb.Subscribe(ctx, "send-user-data")

	user := User{}

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal([]byte(msg.Payload), &user); err != nil {
			panic(err)
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Printf("%+v\n", user)
	}
}
