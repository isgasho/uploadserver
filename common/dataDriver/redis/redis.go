/**
 * redis 连接池封装
 */
package redis

import(
	"github.com/gomodule/redigo/redis"
	"time"
	"strconv"
)

type RedisPool struct {
	redisPool *redis.Pool
}

func New(addr string, maxidle, maxactive, idleTimeout, connTimeout, readTimeout, writeTimeout int) (*RedisPool, error) {
	redisPool := &redis.Pool{
		MaxIdle: maxidle,
		MaxActive: maxactive,
		IdleTimeout: time.Duration( idleTimeout ) * time.Millisecond,
		Dial: func() (redis.Conn, error) {
			c, errDail := redis.DialTimeout("tcp", addr, time.Duration(connTimeout) * time.Millisecond, time.Duration(readTimeout) * time.Millisecond, time.Duration(writeTimeout) * time.Millisecond)
			if errDail != nil {
				return nil, errDail
			}
			return c, nil
		},
	}

	libRedisPool := &RedisPool{redisPool}
	return libRedisPool, nil
}

func (this *RedisPool) Get() (conn redis.Conn, err error) {
	conn = this.redisPool.Get()
	return
}

func String(buf interface{}, err error) (string, error) {
	return redis.String(buf, err)
}

func Int(buf interface{}, err error) (int, error) {
	str, _ := redis.String(buf, err)
	tmpInt64, _ := strconv.ParseInt(str, 10, 64)
	tmpInt := int(tmpInt64)
	return tmpInt, nil
}

func Float32(buf interface{}, err error) (float32, error) {
	str, _ := redis.String(buf, err)
	tmpFloat64, _ := strconv.ParseFloat(str, 64)
	tmpFloat32 := float32(tmpFloat64)
	return tmpFloat32, nil
}

func Values(reply interface{}, err error) ([]interface{}, error) {
	return redis.Values(reply, err)
}