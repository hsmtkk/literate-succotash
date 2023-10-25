package main_test

import (
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
}
