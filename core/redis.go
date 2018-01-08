package core

import (
	"github.com/garyburd/redigo/redis"
)

//执行redis命令
func exec(command string, args ...interface{}) (interface{}, error) {
	c, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer c.Close()
	//选择数据库
	c.Do("SELECT", 0)
	//设置键值
	return c.Do(command, args...)
}
