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
	"github.com/HJH0924/GenericGo/slice"
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
				_ = redisClient.Set(ctx, "name", "Tvux", time.Minute)
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
				_ = redisClient.Set(ctx, "name", "Tvux", time.Minute)
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
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
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
	type args struct {
		key string
		val any
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
			name: "GETSET name Jerry",
			args: args{
				key: "name",
				val: "Jerry",
			},
			wantVal: "Tom",
			before: func(ctx context.Context, t *testing.T) {
				_ = redisClient.Set(ctx, "name", "Tom", time.Minute)
			},
			after: func(ctx context.Context, t *testing.T) {
				res, err := redisClient.Get(ctx, "name").Result()
				assert.NoError(t, err)
				assert.Equal(t, "Jerry", res)
			},
			clean: func(ctx context.Context, t *testing.T) {
				_, err := redisClient.Del(ctx, "name").Result()
				assert.NoError(t, err)
			},
		},
		{
			name: "GETSET key not exist",
			args: args{
				key: "name",
				val: "Jerry",
			},
			wantErr: cache.NewErrKeyNotExist,
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			clean:   func(ctx context.Context, t *testing.T) {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redisCache := NewCache(redisClient)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			tt.before(ctx, t)
			oldVal, err := redisCache.GetSet(ctx, tt.args.key, tt.args.val)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantVal, oldVal)
				tt.after(ctx, t)
			}
			tt.clean(ctx, t)
		})
	}
}

func TestCache_Delete(t *testing.T) {
	type args struct {
		keys []string
	}
	tests := []struct {
		name    string
		args    args
		wantVal int64
		wantErr error
		before  func(ctx context.Context, t *testing.T)
		after   func(ctx context.Context, t *testing.T)
		clean   func(ctx context.Context, t *testing.T)
	}{
		{
			name: "DEL name",
			args: args{
				keys: []string{"name"},
			},
			wantVal: 1,
			before: func(ctx context.Context, t *testing.T) {
				_ = redisClient.Set(ctx, "name", "Tom", time.Minute)
			},
			after: func(ctx context.Context, t *testing.T) {
				_, err := redisClient.Get(ctx, "name").Result()
				assert.Equal(t, redis.Nil, err)
			},
			clean: func(ctx context.Context, t *testing.T) {},
		},
		{
			name: "DEL name age",
			args: args{
				keys: []string{"name", "age"},
			},
			wantVal: 2,
			before: func(ctx context.Context, t *testing.T) {
				_ = redisClient.Set(ctx, "name", "Tom", time.Minute)
				_ = redisClient.Set(ctx, "age", "18", time.Minute)
			},
			after: func(ctx context.Context, t *testing.T) {
				_, err := redisClient.Get(ctx, "name").Result()
				assert.Equal(t, redis.Nil, err)
				_, err = redisClient.Get(ctx, "age").Result()
				assert.Equal(t, redis.Nil, err)
			},
			clean: func(ctx context.Context, t *testing.T) {},
		},
		{
			name: "DEL name age 'not exist key'",
			args: args{
				keys: []string{"name", "age", "phone"},
			},
			wantVal: 2,
			before: func(ctx context.Context, t *testing.T) {
				_ = redisClient.Set(ctx, "name", "Tom", time.Minute)
				_ = redisClient.Set(ctx, "age", "18", time.Minute)
			},
			after: func(ctx context.Context, t *testing.T) {
				_, err := redisClient.Get(ctx, "name").Result()
				assert.Equal(t, redis.Nil, err)
				_, err = redisClient.Get(ctx, "age").Result()
				assert.Equal(t, redis.Nil, err)
			},
			clean: func(ctx context.Context, t *testing.T) {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redisCache := NewCache(redisClient)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			tt.before(ctx, t)
			delNum, err := redisCache.Delete(ctx, tt.args.keys...)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantVal, delNum)
				tt.after(ctx, t)
			}
			tt.clean(ctx, t)
		})
	}
}

func TestCache_LPush(t *testing.T) {
	type args struct {
		key  string
		vals []any
	}
	tests := []struct {
		name    string
		args    args
		wantVal int64
		wantErr error
		before  func(ctx context.Context, t *testing.T)
		after   func(ctx context.Context, t *testing.T)
		clean   func(ctx context.Context, t *testing.T)
	}{
		{
			name: "LPUSH name 'Tvux' 'Tom'",
			args: args{
				key:  "name",
				vals: []any{"Tvux", "Tom"},
			},
			wantVal: 2,
			before:  func(ctx context.Context, t *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				vals := []string{"Tvux", "Tom"}
				for _, val := range vals {
					res, err := redisClient.RPop(ctx, "name").Result()
					assert.NoError(t, err)
					assert.Equal(t, val, res)
				}
			},
			clean: func(ctx context.Context, t *testing.T) {
				err := redisClient.Del(ctx, []string{"name"}...).Err()
				assert.NoError(t, err)
			},
		},
		{
			name: "LPUSH name 'Tom' 'Jerry' after name is exist",
			args: args{
				key:  "name",
				vals: []any{"Tom", "Jerry"},
			},
			wantVal: 3,
			before: func(ctx context.Context, t *testing.T) {
				err := redisClient.LPush(ctx, "name", "Tvux").Err()
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				vals := []string{"Tvux", "Tom", "Jerry"}
				for _, val := range vals {
					res, err := redisClient.RPop(ctx, "name").Result()
					assert.NoError(t, err)
					assert.Equal(t, val, res)
				}
			},
			clean: func(ctx context.Context, t *testing.T) {
				err := redisClient.Del(ctx, "name").Err()
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
			llen, err := redisCache.LPush(ctx, tt.args.key, tt.args.vals...)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantVal, llen)
				tt.after(ctx, t)
			}
			tt.clean(ctx, t)
		})
	}
}

func TestCache_LPop(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		wantVal any
		wantErr error
		before  func(ctx context.Context, t *testing.T)
		after   func(ctx context.Context, t *testing.T)
		clean   func(ctx context.Context, t *testing.T)
	}{
		{
			name: "LPOP name",
			args: args{
				key: "name",
			},
			wantVal: "Jerry",
			before: func(ctx context.Context, t *testing.T) {
				err := redisClient.LPush(ctx, "name", "Tom", "Jerry").Err()
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {},
			clean: func(ctx context.Context, t *testing.T) {
				err := redisClient.Del(ctx, "name").Err()
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
			val, err := redisCache.LPop(ctx, tt.args.key)
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

func TestCache_SAdd(t *testing.T) {
	type args struct {
		key     string
		members []any
	}
	tests := []struct {
		name    string
		args    args
		wantVal int64
		wantErr error
		before  func(ctx context.Context, t *testing.T)
		after   func(ctx context.Context, t *testing.T)
		clean   func(ctx context.Context, t *testing.T)
	}{
		{
			name: "SADD name 'Tom' 'Jerry'",
			args: args{
				key:     "name",
				members: []any{"Tom", "Jerry"},
			},
			wantVal: 2,
			before:  func(ctx context.Context, t *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				res, err := redisClient.SMembers(ctx, "name").Result()
				assert.NoError(t, err)
				assert.True(t, slice.ContainsAll[string](res, []string{"Tom", "Jerry"}))
			},
			clean: func(ctx context.Context, t *testing.T) {
				err := redisClient.Del(ctx, "name").Err()
				assert.NoError(t, err)
			},
		},
		{
			name: "SADD name 'Tom' 'Jerry' 'Tom'",
			args: args{
				key:     "name",
				members: []any{"Tom", "Jerry", "Tom"},
			},
			wantVal: 2,
			before:  func(ctx context.Context, t *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				res, err := redisClient.SMembers(ctx, "name").Result()
				assert.NoError(t, err)
				assert.True(t, slice.ContainsAll[string](res, []string{"Tom", "Jerry"}))
			},
			clean: func(ctx context.Context, t *testing.T) {
				err := redisClient.Del(ctx, "name").Err()
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
			addNum, err := redisCache.SAdd(ctx, tt.args.key, tt.args.members...)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantVal, addNum)
				tt.after(ctx, t)
			}
			tt.clean(ctx, t)
		})
	}
}

func TestCache_SRem(t *testing.T) {
	type args struct {
		key     string
		members []any
	}
	tests := []struct {
		name    string
		args    args
		wantVal int64
		wantErr error
		before  func(ctx context.Context, t *testing.T)
		after   func(ctx context.Context, t *testing.T)
		clean   func(ctx context.Context, t *testing.T)
	}{
		{
			name: "SREM name 'Tom' 'Jerry'",
			args: args{
				key:     "name",
				members: []any{"Tom", "Jerry"},
			},
			wantVal: 2,
			before: func(ctx context.Context, t *testing.T) {
				err := redisClient.SAdd(ctx, "name", "Tom", "Jerry", "Tvux").Err()
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				res, err := redisClient.SMembers(ctx, "name").Result()
				assert.NoError(t, err)
				assert.True(t, slice.ContainsAll[string](res, []string{"Tvux"}))
			},
			clean: func(ctx context.Context, t *testing.T) {
				err := redisClient.Del(ctx, "name").Err()
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
			delNum, err := redisCache.SRem(ctx, tt.args.key, tt.args.members...)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantVal, delNum)
				tt.after(ctx, t)
			}
			tt.clean(ctx, t)
		})
	}
}

func TestCache_IncrBy(t *testing.T) {
	type args struct {
		key   string
		value int64
	}
	tests := []struct {
		name    string
		args    args
		wantVal int64
		wantErr error
		before  func(ctx context.Context, t *testing.T)
		after   func(ctx context.Context, t *testing.T)
		clean   func(ctx context.Context, t *testing.T)
	}{
		{
			name: "INCRBY age 2",
			args: args{
				key:   "age",
				value: 2,
			},
			wantVal: 20,
			before: func(ctx context.Context, t *testing.T) {
				err := redisClient.Set(ctx, "age", 18, time.Minute).Err()
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {},
			clean: func(ctx context.Context, t *testing.T) {
				err := redisClient.Del(ctx, "age").Err()
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
			res, err := redisCache.IncrBy(ctx, tt.args.key, tt.args.value)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantVal, res)
				tt.after(ctx, t)
			}
			tt.clean(ctx, t)
		})
	}
}

func TestCache_DecrBy(t *testing.T) {
	type args struct {
		key       string
		decrement int64
	}
	tests := []struct {
		name    string
		args    args
		wantVal int64
		wantErr error
		before  func(ctx context.Context, t *testing.T)
		after   func(ctx context.Context, t *testing.T)
		clean   func(ctx context.Context, t *testing.T)
	}{
		{
			name: "DECRBY age 2",
			args: args{
				key:       "age",
				decrement: 2,
			},
			wantVal: 18,
			before: func(ctx context.Context, t *testing.T) {
				err := redisClient.Set(ctx, "age", 20, time.Minute).Err()
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {},
			clean: func(ctx context.Context, t *testing.T) {
				err := redisClient.Del(ctx, "age").Err()
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
			res, err := redisCache.DecrBy(ctx, tt.args.key, tt.args.decrement)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantVal, res)
				tt.after(ctx, t)
			}
			tt.clean(ctx, t)
		})
	}
}

func TestCache_IncrByFloat(t *testing.T) {
	type args struct {
		key   string
		value float64
	}
	tests := []struct {
		name    string
		args    args
		wantVal float64
		wantErr error
		before  func(ctx context.Context, t *testing.T)
		after   func(ctx context.Context, t *testing.T)
		clean   func(ctx context.Context, t *testing.T)
	}{
		{
			name: "INCRBYFLOAT score 0.5",
			args: args{
				key:   "score",
				value: 0.5,
			},
			wantVal: 100,
			before: func(ctx context.Context, t *testing.T) {
				err := redisClient.Set(ctx, "score", 99.5, time.Minute).Err()
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {},
			clean: func(ctx context.Context, t *testing.T) {
				err := redisClient.Del(ctx, "score").Err()
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
			res, err := redisCache.IncrByFloat(ctx, tt.args.key, tt.args.value)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantVal, res)
				tt.after(ctx, t)
			}
			tt.clean(ctx, t)
		})
	}
}

func newRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
