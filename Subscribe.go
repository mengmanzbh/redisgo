package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	psc := redis.PubSubConn{c}
	psc.PSubscribe("__keyevent@0__:expired")
	for {
		switch v := psc.Receive().(type) {
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case redis.Message: //单个订阅subscribe
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.PMessage: //模式订阅psubscribe
			fmt.Printf("PMessage: %s %s %s\n", v.Pattern, v.Channel, v.Data)
		case error:
			return

		}
	}

}
