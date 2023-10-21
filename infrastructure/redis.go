package infrastructure

import (
	"context"
	"log"

	"github.com/centraldigital/cfw-core-lib/pkg/util/redisutil"
)

func NewRedis(ctx context.Context) *redisutil.RedisTemplate {

	redisTemplate, err := redisutil.NewRedisTemplate(ctx)
	if err != nil {
		log.Panicf("failed to initial redis-pool: %+v", err)
	}

	_, err = redisTemplate.WithContext(ctx).Ping()
	if err != nil {
		log.Panicf("unable to connect to redis: %+v", err)
	}

	return redisTemplate
}
