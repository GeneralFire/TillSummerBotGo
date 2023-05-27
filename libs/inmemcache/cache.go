package inmemcache

import (
	"sync"
	"time"
)

type CacheInterface[K comparable, V any] interface {
	Put(K, V) bool
	Get(K) (V, bool)
	Pop(K) (V, bool)
	Delete(K)
}

type valueWithCreationTime[V any] struct {
	creationTime time.Time
	value        V
}

func New[K comparable, V any](TTL time.Duration) CacheInterface[K, V] {

	var useTTL bool

	if TTL == 0 {
		useTTL = false
	} else {
		useTTL = true
	}
	return &cache[K, V]{
		storage: make(map[K]valueWithCreationTime[V]),
		TTL:     TTL,
		lock:    sync.RWMutex{},
		useTTL:  useTTL,
	}
}

type cache[K comparable, V any] struct {
	storage      map[K]valueWithCreationTime[V]
	TTL          time.Duration
	lock         sync.RWMutex
	defaultValue V
	useTTL       bool
}

func (c *cache[K, V]) unsafePut(key K, value V) {
	c.storage[key] = valueWithCreationTime[V]{
		creationTime: time.Now(),
		value:        value,
	}
}

func (c *cache[K, V]) Put(key K, value V) bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	v, ok := c.storage[key]
	if !ok {
		c.unsafePut(key, value)
		return true
	}

	if c.useTTL {
		if v.creationTime.Add(c.TTL).After(time.Now()) {
			return false
		}
	}

	c.unsafePut(key, value)
	return true
}

func (c *cache[K, V]) Get(key K) (V, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	v, ok := c.storage[key]
	if !ok {
		return c.defaultValue, false
	}
	if c.useTTL {
		if v.creationTime.Add(c.TTL).Before(time.Now()) {
			delete(c.storage, key)
			return c.defaultValue, false
		}
	}

	return v.value, true
}

func (c *cache[K, V]) Pop(key K) (V, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	defer delete(c.storage, key)

	v, ok := c.storage[key]
	if !ok {
		return c.defaultValue, false
	}

	if c.useTTL {
		if v.creationTime.Add(c.TTL).Before(time.Now()) {
			delete(c.storage, key)
			return c.defaultValue, false
		}
	}

	return v.value, true
}

func (c *cache[K, V]) Delete(key K) {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.storage, key)
}
