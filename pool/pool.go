// Copyright 2016 Nikugame. All Rights Reserved

package pool

import (
	"fmt"
	"os"
	"time"

	"github.com/nikgame/public/config"
	"github.com/nikgame/public/tools"

	"github.com/garyburd/redigo/redis"
)

//Client reids连接池客户端
var Client *redis.Pool

func init() {
	var conf redisConf
	conf.init()
	Client = &redis.Pool{
		MaxIdle:     conf.idle,
		MaxActive:   conf.active,
		IdleTimeout: conf.Timeout(),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", conf.connPath())
			if err != nil {
				fmt.Printf("[%s] Redis Connect ERROR: %s", tools.TimeStampString(), err)
				os.Exit(1)
			}
			if conf.mode {
				if _, err := c.Do("AUTH", conf.auth); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, nil
		},
	}
}

func (r *redisConf) init() {
	conf, _ := config.NewConfig("ini", "conf/settings.conf")
	r.host = conf.DefaultString("Redis::host", "127.0.0.1")
	r.port = conf.DefaultString("Redis::port", "6379")
	r.mode = conf.DefaultBool("Redis::mode", false)
	if r.mode {
		r.auth = conf.DefaultString("Redis::password", "")
	}
	r.wait = conf.DefaultInt("Redis::timeout", 180)
	r.idle = conf.DefaultInt("Redis::idle", 5)
	r.active = conf.DefaultInt("Redis::active", 500)
}

func (r *redisConf) connPath() string {
	return tools.StringJoin(":", r.host, r.port)
}

func (r *redisConf) Timeout() time.Duration {
	return time.Duration(r.wait) * time.Second
}

type redisConf struct {
	host   string
	port   string
	auth   string
	mode   bool
	wait   int
	active int
	idle   int
}
