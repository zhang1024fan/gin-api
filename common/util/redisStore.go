package util

import (
	"context"
	"gin-api/common/constant"
	"gin-api/pkg/redis"
	"log"
	"time"
)

var ctx = context.Background()

type RedisStore struct{}

// Set 实现 base64Captcha.Store 接口方法，必须返回 error 类型
func (r RedisStore) Set(id string, value string) error {
	key := constant.LOGIN_CODE + id
	err := redis.RedisDb.Set(ctx, key, value, time.Minute*5).Err()
	if err != nil {
		log.Println("Redis Set Error:", err)
		return err
	}
	return nil
}

// Get 获取验证码
func (r RedisStore) Get(id string, clear bool) string {
	key := constant.LOGIN_CODE + id
	val, err := redis.RedisDb.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return val
}

// Verify 验证码校验
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := r.Get(id, clear)
	return v == answer
}
