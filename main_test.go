package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlwaysOK(t *testing.T) {
	// blank
}

func TestWithAssert(t *testing.T) {
	assert.True(t, true)
}
