// Package redis
/**
* @Project : GenericGo
* @File    : cache_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/10/8 17:05
**/

package redis

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/HJH0924/GenericGo/cache"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var redisClient *redis.Client

func TestMain(m *testing.M) {
	// 初始化 Redis 客户端
	redisClient = newRedisClient()
	// 检查 Redis 是否可以 ping 通
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		fmt.Printf("Failed to connect to Redis: %v\n", err)
		os.Exit(1)
	}
	// 运行所有测试用例
	exitCode := m.Run()
	// 退出前清理资源
	redisClient.Close()
	// 退出程序
	os.Exit(exitCode)
}

func TestCache_Set(t *testing.T) {
	type args struct {
		key        string
		val        any
		expiration time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
		after   func(ctx context.Context, t *testing.T)
		clean   func(ctx context.Context, t *testing.T)
	}{
		{
			name: "SET name Tvux EX 60",
			args: args{
				key:        "name",
				val:        "Tvux",
				expiration: time.Minute,
			},
			after: func(ctx context.Context, t *testing.T) {
				res, err := redisClient.Get(ctx, "name").Result()
				assert.NoError(t, err)
				assert.Equal(t, "Tvux", res)
			},
			clean: func(ctx context.Context, t *testing.T) {
				_, err := redisClient.Del(ctx, "name").Result()
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redisCache := NewCache(redisClient)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			err := redisCache.Set(ctx, tt.args.key, tt.args.val, tt.args.expiration)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
				tt.after(ctx, t)
			}
			tt.clean(ctx, t)
		})
	}
}

func TestCache_SetNX(t *testing.T) {
	type args struct {
		key        string
		val        any
		expiration time.Duration
	}
	tests := []struct {
		name     string
		args     args
		wantBool bool
		wantErr  error
		before   func(ctx context.Context, t *testing.T)
		after    func(ctx context.Context, t *testing.T)
		clean    func(ctx context.Context, t *testing.T)
	}{
		{
			name: "SETNX name Tvux EX 60",
			args: args{
				key:        "name",
				val:        "Tvux",
				expiration: time.Minute,
			},
			wantBool: true,
			before:   func(ctx context.Context, t *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				res, err := redisClient.Get(ctx, "name").Result()
				assert.NoError(t, err)
				assert.Equal(t, "Tvux", res)
			},
			clean: func(ctx context.Context, t *testing.T) {
				_, err := redisClient.Del(ctx, "name").Result()
				assert.NoError(t, err)
			},
		},
		{
			name: "SETNX name Tvux EX 60 return false",
			args: args{
				key:        "name",
				val:        "Tvux",
				expiration: time.Minute,
			},
			wantBool: false,
			before: func(ctx context.Context, t *testing.T) {
				_ = redisClient.Set(ctx, "name", "Tvux", time.Minute).Err()
			},
			after: func(ctx context.Context, t *testing.T) {
				res, err := redisClient.Get(ctx, "name").Result()
				assert.NoError(t, err)
				assert.Equal(t, "Tvux", res)
			},
			clean: func(ctx context.Context, t *testing.T) {
				_, err := redisClient.Del(ctx, "name").Result()
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redisCache := NewCache(redisClient)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			tt.before(ctx, t)
			res, err := redisCache.SetNX(ctx, tt.args.key, tt.args.val, tt.args.expiration)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantBool, res)
				tt.after(ctx, t)
			}
			tt.clean(ctx, t)
		})
	}
}

func TestCache_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		wantVal string
		wantErr error
		before  func(ctx context.Context, t *testing.T)
		after   func(ctx context.Context, t *testing.T)
		clean   func(ctx context.Context, t *testing.T)
	}{
		{
			name: "GET name",
			args: args{
				key: "name",
			},
			wantVal: "Tvux",
			before: func(ctx context.Context, t *testing.T) {
				_ = redisClient.Set(ctx, "name", "Tvux", time.Minute).Err()
			},
			after: func(ctx context.Context, t *testing.T) {},
			clean: func(ctx context.Context, t *testing.T) {
				_, err := redisClient.Del(ctx, "name").Result()
				assert.NoError(t, err)
			},
		},
		{
			name: "GET not exist key",
			args: args{
				key: "age",
			},
			wantErr: cache.NewErrKeyNotExist,
			before: func(ctx context.Context, t *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {},
			clean: func(ctx context.Context, t *testing.T) {
				_, err := redisClient.Del(ctx, "name").Result()
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redisCache := NewCache(redisClient)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			tt.before(ctx, t)
			val, err := redisCache.Get(ctx, tt.args.key)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantVal, val)
				tt.after(ctx, t)
			}
			tt.clean(ctx, t)
		})
	}
}

func TestCache_GetSet(t *testing.T) {

}

func TestCache_Delete(t *testing.T) {

}

func TestCache_LPush(t *testing.T) {

}

func TestCache_LPop(t *testing.T) {

}

func TestCache_SAdd(t *testing.T) {

}

func TestCache_SRem(t *testing.T) {

}

func TestCache_IncrBy(t *testing.T) {

}

func TestCache_DecrBy(t *testing.T) {

}

func TestCache_IncrByFloat(t *testing.T) {

}

func newRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
