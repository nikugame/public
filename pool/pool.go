// Copyright 2016 Nikugame. All Rights Reserved

package pool

import "github.com/garyburd/redigo/redis"

//Client reids连接池客户端
var Client redis.Pool

func init() {

}

type redisConf struct {
	host string
	port string
	auth string
}
