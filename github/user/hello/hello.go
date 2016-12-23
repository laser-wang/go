//package test
//package test2
package main

import (
	"fmt"
	"time"

	"go/github.com/garyburd/redigo/redis"
)

var (
	// 定义常量
	RedisClient *redis.Pool
	REDIS_HOST  string
	REDIS_DB    int
)

/**dfdsfd*/
func main() {
	//	fmt.Printf("hello, world\n")
	test()

}

func test() {
	fmt.Printf("hello, world\n")
	// 从池里获取连接
	c := RedisClient.Get()
	v, err := c.Do("SET", "name", "red1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	v, err = redis.String(c.Do("GET", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	// 用完后将连接放回连接池
	defer c.Close()
}

func init() {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST = "192.168.0.10:6379"
	REDIS_DB = 1
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     10,
		MaxActive:   100,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}
