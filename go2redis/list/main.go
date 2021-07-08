package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

func main() {
	dial, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatalf("redis.Dial err :%s", err)
	}
	defer dial.Close()

	_, err = dial.Do("lpush", "mylist", 200,100)
	if err != nil {
		log.Fatalf("dial do lupsh err :%s", err)
	}
	s, err := redis.String(dial.Do("LRange", "mylist",0,-1))
	if err != nil {
		log.Fatalf("dial do lpush: %s", err)
	}
	fmt.Println(s)

	//list操作
	_, _ = dial.Do("lpush", "jason", "a", "b", "c")
	//lrange
	//values, _ := redis.Values(r.Do("lrange", "jason", "0", "-1"))
	strings, _ := redis.Strings(dial.Do("lrange", "jason", "0", "-1"))
	for _, v := range strings {
		fmt.Println(v)
	}
	//for _, v := range values {
	//	fmt.Println(v)
	//	fmt.Println(string(v.([]byte)))
	//}





}
