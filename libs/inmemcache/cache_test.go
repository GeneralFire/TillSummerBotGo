package inmemcache_test

import (
	"testing"
	"time"

	"route256/libs/inmemcache"

	"github.com/stretchr/testify/assert"
)

type KeyValuePair struct {
	K int
	V int
}

var SampleKey = 3
var SampleValue = 3
var SampleTTL = time.Millisecond * 100

func TestEmptyCacheResponseNOTTL(t *testing.T) {
	c := inmemcache.New[int, int](0)
	_, ok := c.Get(SampleKey)
	assert.Equal(t, false, ok)
}

func TestEmptyCacheResponse(t *testing.T) {
	c := inmemcache.New[int, int](0)
	_, ok := c.Get(SampleKey)
	assert.Equal(t, false, ok)
}

func TestDataTTLExceed(t *testing.T) {
	c := inmemcache.New[int, int](SampleTTL)
	ok := c.Put(SampleKey, SampleValue)

	assert.Equal(t, ok, true)

	time.Sleep(SampleTTL * 2)
	v, ok := c.Get(SampleKey)

	assert.Equal(t, 0, v)
	assert.Equal(t, false, ok)
}

func TestDataTTL(t *testing.T) {
	c := inmemcache.New[int, int](SampleTTL)
	ok := c.Put(SampleKey, SampleValue)

	assert.Equal(t, true, ok)

	v, ok := c.Get(SampleValue)

	assert.Equal(t, SampleValue, v)
	assert.Equal(t, true, ok)
}

func TestMultipleDataNOTTL(t *testing.T) {
	c := inmemcache.New[int, int](0)

	ok := c.Put(SampleKey, SampleValue)
	assert.Equal(t, true, ok)

	ok = c.Put(SampleKey+1, SampleValue+1)
	assert.Equal(t, true, ok)

	v, ok := c.Get(SampleValue)
	assert.Equal(t, SampleValue, v)
	assert.Equal(t, true, ok)

	v, ok = c.Get(SampleValue + 1)
	assert.Equal(t, SampleValue+1, v)
	assert.Equal(t, true, ok)
}

func TestMultipleDataOneTTLExceed(t *testing.T) {
	c := inmemcache.New[int, int](SampleTTL)

	ok := c.Put(SampleKey, SampleValue)
	assert.Equal(t, true, ok)

	ok = c.Put(SampleKey+1, SampleValue+1)
	assert.Equal(t, true, ok)

	v, ok := c.Get(SampleValue)
	assert.Equal(t, v, SampleValue)
	assert.Equal(t, true, ok)

	time.Sleep(SampleTTL * 2)
	v, ok = c.Get(SampleValue + 1)
	assert.Equal(t, v, 0)
	assert.Equal(t, ok, false)
}

func TestKeyAlreadyInTTL(t *testing.T) {
	c := inmemcache.New[int, int](SampleTTL)

	ok := c.Put(SampleKey, SampleValue)
	assert.Equal(t, ok, true)

	ok = c.Put(SampleKey, SampleValue)
	assert.Equal(t, ok, false)
}

func TestDeleteKeyTTL(t *testing.T) {
	c := inmemcache.New[int, int](SampleTTL)

	ok := c.Put(SampleKey, SampleValue)
	assert.Equal(t, true, ok)

	c.Delete(SampleKey)
	v, ok := c.Get(SampleKey)
	assert.Equal(t, v, 0)
	assert.Equal(t, false, ok)
}

func TestPopKeyTTL(t *testing.T) {
	c := inmemcache.New[int, int](SampleTTL)

	ok := c.Put(SampleKey, SampleValue)
	assert.Equal(t, true, ok)

	v, ok := c.Pop(SampleKey)
	assert.Equal(t, v, SampleValue)
	assert.Equal(t, true, ok)
}

func TestMultipleDataPopGetNOTTL(t *testing.T) {
	c := inmemcache.New[int, int](0)

	var IterCount = 20

	for i := 0; i < IterCount; i++ {
		ok := c.Put(i, IterCount-i)
		assert.Equal(t, true, ok)
	}

	for i := 0; i < IterCount; i++ {
		v, ok := c.Pop(i)
		assert.Equal(t, true, ok)
		assert.Equal(t, IterCount-i, v)
	}

	for i := 0; i < IterCount; i++ {
		v, ok := c.Get(i)
		assert.Equal(t, false, ok)
		assert.Equal(t, 0, v)
	}
}
