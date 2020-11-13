package storage

import (
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"

	"workWechat_api/pkg/config"
	"workWechat_api/pkg/logging"
)

var (
	pool *redisPoll
	once sync.Once
)

type redisPoll struct {
	pool *redis.Pool
}


func newPool()  *redisPoll {
	logging.Info("creating redis pool")
	return &redisPoll{&redis.Pool{
		MaxIdle:     5,
		MaxActive:   0,
		IdleTimeout: time.Duration(120),
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				config.Redis.Network,
				config.Redis.Server,
				redis.DialReadTimeout(time.Duration(5000)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(5000)*time.Millisecond),
				redis.DialConnectTimeout(time.Duration(5000)*time.Millisecond),
				redis.DialDatabase(config.Redis.DB),
				redis.DialPassword(config.Redis.Password),
			)
		},
	},
	}
}

func (rp *redisPoll) Get(key string)  (string,error){
	conn := rp.pool.Get()
	defer conn.Close()
	logging.Debug("get %s from redis",key)
	return redis.String(conn.Do("GET" ,key))
}

func (rp *redisPoll) Set(key string,value string)  bool{
	conn :=rp.pool.Get()
	defer conn.Close()
	_,err := conn.Do("SET" ,key,value)
	logging.Debug("set %s with %s to redis",key, value)
	if err != nil {
		logging.Error("set %s error: %T",key,err)
		return false
	}
	logging.Debug("set expire of 7200 second on %s",key)
	_,err =redis.Bool(conn.Do("EXPIRE", key, 2*3600))  //两小时过期
	if err != nil {
		logging.Error("set expire error: %T",err)
		return false
	}
	return true

}

func RedisPool() *redisPoll {
	once.Do(func() {
		pool = newPool()
	})
	return pool

}
