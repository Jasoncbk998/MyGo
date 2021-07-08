package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

/**
redis 存取结构体
*/
func main() {
	top1 := Book{
		BookName:  "james",
		Author:    "lerbon",
		PageCount: "900",
		Press:     "math",
	}

	r, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}

	if _, err := r.Do("HMset", redis.Args{}.Add("top2").AddFlat(&top1)...); err != nil {
		log.Fatalf("hmset err:%s", err)
	}
	t := Book{}

	for _, item := range []string{"top1","top2"} {
		value, _ := redis.Values(r.Do("HGETALL", item))
		redis.ScanStruct(value, &t)
		fmt.Printf("%s[%+v]\n", item, t)
	}

	if  stringsValue, err := redis.Strings(r.Do("HMGET", "top2", "BookName", "Author", "Press")); err!=nil{
		log.Fatalf("err:%s",err)
	}else {
		fmt.Printf("hmget:%+v\n", stringsValue)

	}

}

type Book struct {
	BookName  string
	Author    string
	PageCount string
	Press     string
}
