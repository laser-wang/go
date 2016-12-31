//package test
//package test2
package main

import (
	"fmt"
	"strconv"
	"time"

	l4g "github.com/alecthomas/log4go"
	"github.com/garyburd/redigo/redis"
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
	testRedis()
	//	testLog()
	testLog2()
	testLog3()
}

func testLog3() {

}

func testLog2() {

	logger3 := l4g.Logger{}
	logger3.LoadConfiguration("example.xml") //使用加载配置文件,类似与java的log4j.propertites
	logger3.Log(l4g.FINE, "testLog2", "FINE..1")
	logger3.Log(l4g.TRACE, "testLog2", "TRACE..1")
	logger3.Log(l4g.DEBUG, "testLog2", "DEBUG..1")
	logger3.Log(l4g.INFO, "testLog2", "INFO..1")
	logger3.Log(l4g.ERROR, "testLog2", "error..1")

	//	time.Sleep(time.Second * 10)

	logger3.Log(l4g.ERROR, "testLog2", "error..111111111")

	//	defer logger3.Close() //注:如果不是一直运行的程序,请加上这句话,否则主线程结束后,也不会输出和log到日志文件

	go testLog2Sub1(&logger3)
	go testLog2Sub2(&logger3)

	//	l4g.AddFilter("stdout", l4g.DEBUG, l4g.NewConsoleLogWriter())             //输出到控制台,级别为DEBUG
	//	l4g.AddFilter("file", l4g.DEBUG, l4g.NewFileLogWriter("test.log", false)) //输出到文件,级别为DEBUG,文件名为test.log,每次追加该原文件
	//	l4g.LoadConfiguration("example.xml") //使用加载配置文件,类似与java的log4j.propertites
	//	l4g.ERROR("the time is now :%s -- %s", "213", "sad")
	defer l4g.Close() //注:如果不是一直运行的程序,请加上这句话,否则主线程结束后,也不会输出和log到日志文件

	fmt.Println("testLog2 OK")
	fmt.Println("==========================")
}

func testLog2Sub1(log *l4g.Logger) {
	for i := 0; i < 1; i++ {

		log.Log(l4g.INFO, "testLog2Sub1", "INFO..===="+strconv.Itoa(i))
	}
	log.Log(l4g.INFO, "testLog2Sub1", "INFO..1")
	log.Log(l4g.ERROR, "testLog2Sub1", "error..1")
}
func testLog2Sub2(log *l4g.Logger) {
	for i := 0; i < 1; i++ {

		log.Log(l4g.INFO, "testLog2Sub2", "INFO..*****"+strconv.Itoa(i))
	}
	log.Log(l4g.INFO, "testLog2Sub2", "INFO..1")
	log.Log(l4g.ERROR, "testLog2Sub2", "error..1")
}

//func testLog() {
//	logPath := "c:\tmp"
//	TRACE = fileLogger.NewDefaultLogger(logPath, "trace.log")
//	INFO = fileLogger.NewDefaultLogger(logPath, "info.log")
//	WARN = fileLogger.NewDefaultLogger(logPath, "warn.log")
//	ERROR = fileLogger.NewDefaultLogger(logPath, "error.log")

//	TRACE.SetPrefix("[TRACE] ")
//	INFO.SetPrefix("[INFO] ")
//	WARN.SetPrefix("[WARN] ")
//	ERROR.SetPrefix("[ERROR] ")

//	//	logFile = fileLogger.NewDefaultLogger("/usr/local/aiwuTech/log", "test.log")
//	//	logFile.SetLogLevel(fileLogger.INFO) //trace log will not be print

//	i := 1
//	TRACE.Printf("This is the No[%v] TRACE log using fileLogger that written by aiwuTech.", i)
//	INFO.Printf("This is the No[%v] INFO log using fileLogger that written by aiwuTech.", i)
//	WARN.Printf("This is the No[%v] WARN log using fileLogger that written by aiwuTech.", i)
//	ERROR.Printf("This is the No[%v] ERROR log using fileLogger that written by aiwuTech.", i)

//	defer TRACE.Close()
//	defer INFO.Close()
//	defer WARN.Close()
//	defer ERROR.Close()

//	fmt.Println("testLog OK")
//	fmt.Println("==========================")
//}

func testRedis() {
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
	fmt.Println("testRedis OK")
	fmt.Println("==========================")
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
