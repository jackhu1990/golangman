package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"time"
)

func main() {
	conn , err := redis.DialTimeout("tcp", "127.0.0.1:6379", 0, 1*time.Second, 1*time.Second)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	size ,err:= conn.Do("DBSIZE")
	fmt.Printf("size is %d \n",size)

	_,err = conn.Do("SET","user:user0",123)
	_,err = conn.Do("SET","user:user1",456)
	_,err = conn.Do("APPEND","user:user0",87)

	user0,err := redis.Int(conn.Do("GET","user:user0"))
	user1,err := redis.Int(conn.Do("GET","user:user1"))

	fmt.Printf("user0 is %d , user1 is %d \n",user0,user1)
}