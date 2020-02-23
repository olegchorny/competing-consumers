package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

const redisListName = "access"

var client *redis.Client

func init() {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		fmt.Printf("REDIS_HOST env var absent!")
		os.Exit(1)
	}

	port := os.Getenv("REDIS_PORT")
	if port == "" {
		fmt.Printf("REDIS_PORT env var absent!")
		os.Exit(1)
	}
	client = redis.NewClient(&redis.Options{Addr: host + ":" + port, PoolSize: 500})
	err := client.Ping().Err()
	if err != nil {
		fmt.Printf("Unable to connect to Redis at %s:%s", host, port)
		os.Exit(1)
	}
}

func main() {
	forever := make(chan bool)
	go func() {
		for {
			access, err := client.LPop(redisListName).Result()
			if err != nil {
				fmt.Println("Unable LPOP from the list. List is empty?", err)
				//			os.Exit(1)
			} else {
				fmt.Println("new message from sender: ", access)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
