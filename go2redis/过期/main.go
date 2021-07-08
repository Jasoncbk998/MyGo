package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

func main() {
	r, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	//两种get
	do, err := r.Do("Get", "abc")
	fmt.Println(string(do.([]byte))) //断言

	r.Send("Get", "abc")
	r.Flush()
	receive, err := r.Receive()
	fmt.Println(string(receive.([]byte)))

	_, err = r.Do("Set", "jason1", "测试是否过期")
	if err != nil {
		panic(err)
	}

	//24*3600 一天
	_, err = r.Do("Set", "jason2", "测试是否过期", "EX", 3) //3秒后过期
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Duration(4) * time.Second)
	b, err := redis.Bool(r.Do("EXISTS", "jason2"))
	//验证过期
	if b {
		fmt.Println("jason2 存在")
	} else {
		fmt.Println("jason2 不存在")
	}
	//删除一个key
	_, err = r.Do("DEL", "jason")
	b2, _ := redis.Bool(r.Do("EXISTS", "jason"))
	if !b2 {
		fmt.Println("删除key:jason成功")
	}
	if err != nil {
		log.Fatalf("del key err:%s", err)
	}





	//s := student{}
	//redis.ScanStruct(values, s)
	//fmt.Println(s)

	//r.Do("SET", "school1:teacher", "Teacher Y")
}

