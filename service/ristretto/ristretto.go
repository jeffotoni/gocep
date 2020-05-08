package ristretto

import (
	"github.com/dgraph-io/ristretto"
	"github.com/jeffotoni/gocep/config"
	"log"
	"sync"
	"time"
)

var (
	once      sync.Once
	cacheOnce *ristretto.Cache
	err       error
)

func Run() *ristretto.Cache {
	once.Do(func() {
		if cacheOnce != nil {
			return
		}
		cacheOnce, err = ristretto.NewCache(&ristretto.Config{
			NumCounters: config.NumCounters, // Num keys to track frequency of (30M).
			MaxCost:     config.MaxCost,     // Maximum cost of cache (2GB).
			BufferItems: config.BufferItems, // Number of keys per Get buffer.
		})
		if err != nil {
			log.Println(err)
			return
		}
	})
	return cacheOnce
}

func SetTTL(key, value string, ttl time.Duration) bool {
	if len(key) < 0 || len(value) < 0 {
		return false
	}
	cache := Run()
	if cache.SetWithTTL(key, value, 1, ttl) {
		time.Sleep(10 * time.Millisecond)
		return true
	}
	return false
}

func Set(key, value string) bool {
	if len(key) < 0 || len(value) < 0 {
		return false
	}
	cache := Run()
	if cache.Set(key, value, 1) {
		time.Sleep(10 * time.Millisecond)
		return true
	}
	return false
}

func Get(key string) string {
	if len(key) <= 0 {
		return ""
	}
	cache := Run()
	value, found := cache.Get(key)
	if !found {
		return ""
	}
	return value.(string)
}
