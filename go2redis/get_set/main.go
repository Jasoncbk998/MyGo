package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

var dial redis.Conn

func main() {

	dial, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatalf("redis dial err : %s", err)
	}
	defer dial.Close()

	//redisSet("aaa", 10)
	//设置过期时间
	//_, err = dial.Do("expire", "abc", 10)
	//if err != nil {
	//	log.Fatalf("expire err:%s",err)
	//}
	//Send + Flush + Receive = Do
	getInt := redisGetInt("abc")
	println(getInt)

	//ints, err := redis.Ints(dial.Do("mget", "abc", "efg"))
	//for _, v := range ints {
	//	fmt.Println(v)
	//}

	// 批量 set
	//_, err = dial.Do("MSet", "abc", 100, "efg", 300)
	//if err != nil {
	//	log.Fatalf("MSet err :%s", err)
	//}

	//i2, err :=

	//if err !=nil{
	//	log.Fatalf("MGet err:%s",err)
	//}
	//for _,v :=range i2{
	//	fmt.Println(v)
	//}

}

func redisGetInt(key string) int {
	i, err := redis.Int(dial.Do("Get", "aaa"))
	if err != nil {
		log.Fatalf("get int err :%s", err)
	}
	return i
}

func redisMGet() []int {
	ints, err := redis.Ints(dial.Do("mget", "abc", "efg"))
	if err != nil {
		log.Fatalf("set err :%s", err)
	}
	return ints
}

func redisSet(key string, value int) {
	_, err := dial.Do("Set", "aaa", 10)
	if err != nil {
		log.Fatalf("set err :%s", err)
	}
}
