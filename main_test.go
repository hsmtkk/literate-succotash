package main_test

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestAlwaysOK(t *testing.T) {
	// blank
}

func TestWithAssert(t *testing.T) {
	assert.True(t, true)
}

func TestRedis(t *testing.T) {
	clt := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer clt.Close()
	ctx := context.Background()
	err := clt.Set(ctx, "key", "value", 0).Err()
	assert.Nil(t, err)

	val, err := clt.Get(ctx, "key").Result()
	assert.Nil(t, err)
	assert.Equal(t, val, "value")
}
