package syncx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnce_Do(t *testing.T) {
	var (
		v    int
		once Once
	)

	once.Do(func() {
		v++
	})
	assert.Equal(t, v, 1)

	once.Do(func() {
		v++
	})
	once.Do(func() {
		v++
	})
	assert.Equal(t, v, 1)
}

func TestOnce_Reset(t *testing.T) {
	var (
		v    int
		once Once
	)

	once.Do(func() {
		v++
	})
	once.Do(func() {
		v++
	})
	once.Do(func() {
		v++
	})
	assert.Equal(t, v, 1)

	once.Reset()
	once.Do(func() {
		v++
	})
	once.Do(func() {
		v++
	})
	once.Do(func() {
		v++
	})
	assert.Equal(t, v, 2)
}
