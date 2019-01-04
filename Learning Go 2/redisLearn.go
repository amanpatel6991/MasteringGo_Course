package main

import (
	"github.com/go-redis/redis"
	"fmt"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "password", // no password set
		DB:       0,  // use default DB
	})

	defer client.Close()

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	//err = client.Set("token1" , "aman" , time.Second*1000).Err()
	//if err != nil {
	//	panic(err)
	//}

	val, err := client.Get("token1").Result()
	if err != nil {
		panic(err)
	}
	//
	//val1, err1 := client.TTL("token1").Result()
	//if err1 != nil {
	//	panic(err1)
	//}

	fmt.Println("val", val)
}
