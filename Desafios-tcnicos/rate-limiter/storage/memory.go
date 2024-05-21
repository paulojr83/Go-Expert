package storage

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

type MemoryStorage struct {
	mu    sync.Mutex
	store map[string]*item
}

type item struct {
	value      interface{}
	expiration int64
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		store: make(map[string]*item),
	}
}

func (m *MemoryStorage) Get(ctx context.Context, key string) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	it, found := m.store[key]
	if !found || (it.expiration > 0 && it.expiration < time.Now().UnixNano()) {
		return "", redis.Nil
	}

	return it.value.(string), nil
}

func (m *MemoryStorage) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	var exp int64
	if expiration > 0 {
		exp = time.Now().Add(expiration).UnixNano()
	}

	m.store[key] = &item{
		value:      value,
		expiration: exp,
	}

	return nil
}

func (m *MemoryStorage) Incr(ctx context.Context, key string) (int64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	it, found := m.store[key]
	if !found || (it.expiration > 0 && it.expiration < time.Now().UnixNano()) {
		m.store[key] = &item{value: "1"}
		return 1, nil
	}

	val, _ := it.value.(string)
	num := 0
	fmt.Sscanf(val, "%d", &num)
	num++
	it.value = fmt.Sprintf("%d", num)

	return int64(num), nil
}

func (m *MemoryStorage) Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	it, found := m.store[key]
	if !found {
		return false, nil
	}

	if expiration > 0 {
		it.expiration = time.Now().Add(expiration).UnixNano()
	} else {
		it.expiration = 0
	}

	return true, nil
}
