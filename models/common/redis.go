package common

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	// 定义常量
	RedisClient		*redis.Pool
	REDIS_HOST	string
	REDIS_PORT	string
)

/**
 * @author    wangtao
 * @date      2018/11/2 9:37
 * @param     nil
 * @return    nil
 * @desc      初始化redis连接池
 */
func init() {
	// 从配置文件中获取redis的IP以及db
	REDIS_HOST = beego.AppConfig.String("redis_host")
	REDIS_PORT = beego.AppConfig.String("redis_port")

	// 建立连接池
	RedisClient = &redis.Pool{
		MaxIdle:  beego.AppConfig.DefaultInt("redis.maxidle",1),
		MaxActive:  beego.AppConfig.DefaultInt("redis.maxactive",10),
		IdleTimeout:  180*time.Second,
		Dial:func() (redis.Conn,error){
			rs,err := redis.Dial("tcp",REDIS_HOST+":"+REDIS_PORT)
			if err != nil {
				return nil,err
			}
			rs.Do("SELECT",0)
			return rs,nil
		},
	}

}
