package core

import (
	"github.com/garyburd/redigo/redis"
	"time"
)
//启用redis 连接池
var (
	// RedisPool RedisPool连接池实例
	RedisPool *redis.Pool
)


func initRedisPool() *redis.Pool {
	pool := &redis.Pool{
		IdleTimeout:  300 * time.Second,
		Wait:         true,
	}

	return pool
}
//执行redis命令
func exec(command string, args ...interface{}) (interface{}, error) {
	c, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer c.Close()
	//选择数据库
	c.Do("SELECT", 0)
	//设置键值
	return c.Do(command, args...)
}




