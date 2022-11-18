package main

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var ctx = context.Background()

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(user); err != nil {
			panic(err)
		}

		payload, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}

		opt, err := redis.ParseURL("redis://default:KwtdpJJ5W7l09hf1ak0VDO6j6gDfdHhd@redis-16342.c295.ap-southeast-1-1.ec2.cloud.redislabs.com:16342")
		if err != nil {
			panic(err)
		}

		rdb := redis.NewClient(opt)

		if err := rdb.Publish(ctx, "send-user-data", payload).Err(); err != nil {
			panic(err)
		}

		return c.SendStatus(200)
	})

	app.Listen(":3000")
}
