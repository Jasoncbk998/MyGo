package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

func main() {

	r, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	var tmp_data []*RewardInfo
	for i := 0; i < 10; i++ {
		tmp_data = append(tmp_data, &RewardInfo{
			Type:      1,
			Uin:       uint64(i),
			OrderId:   "1234567",
			Username:  "sbweijun",
			Desc:      "sb",
			Cmoney:    0,
			ScriptId:  uint32(1),
			AuthorUin: 123,
		})
	}

	marshal, _ := json.Marshal(tmp_data)

	if _, err := r.Do("SET", "json1", marshal, "EX", 86400); err != nil {
		log.Fatalf("set json1 err:%s", err)
	}
	if  data , err := redis.Bytes(r.Do("GET", "json1")) ;err!=nil{
		log.Fatalf("set json1 err:%s", err)
	}else {

		fmt.Println("json1:",string(data))
	}

}

type RewardInfo struct {
	Type      uint32 `json:Type`
	Uin       uint64 `json:Uin`
	OrderId   string `json:OrderId`
	Username  string `json:Username`
	Desc      string `json:Desc`
	Cmoney    uint64 `json:Cmoney`
	ScriptId  uint32 `json:ScriptId`
	AuthorUin uint64 `json:AuthorUin`
}
