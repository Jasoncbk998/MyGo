package main

import "github.com/garyburd/redigo/redis"

func main() {
	r, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer r.Close()






}





