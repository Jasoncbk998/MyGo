package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	r, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	//std := student{
	//	Name:   "kebo",
	//	Gender: "man",
	//	Height: "180",
	//}
	//_, err = r.Do("MSet", "a1", 100, "a2", 300)
	//ints, _ := redis.Ints(r.Do("MGet", "a1", "a2"))
	//for _,v :=range ints{
	//	fmt.Println(v)
	//}

	//_, err = r.Do("HMset", redis.Args{"std1"}.AddFlat(std)...)
	//if err!=nil{
	//	log.Fatalf("mset err :%s",err)
	//}

	strings, _ := redis.Strings(r.Do("HGETALL", "std1"))
	for k,v :=range strings{
		fmt.Println(k,v)
	}

}

type student struct {
	Name   string `redis:"name"`
	Gender string`redis:"gender"`
	Height string`redis:"height"`
}
