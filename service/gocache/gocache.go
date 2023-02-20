package gocache

import (
	"sync"
	"time"

	gcache "github.com/patrickmn/go-cache"
)

type Gocache struct {
	Once       sync.Once
	GcacheOnce *gcache.Cache
	Err        error
}

var gcs = new(Gocache)

func Run() *gcache.Cache {
	gcs.Once.Do(func() {
		if gcs.GcacheOnce != nil {
			return
		}
		gcs.GcacheOnce = gcache.New(24*time.Hour, 24*time.Hour)
	})
	return gcs.GcacheOnce
}

func SetTTL(key, value string, ttl time.Duration) bool {
	if len(key) < 0 || len(value) < 0 {
		return false
	}
	g := Run()
	g.Set(key, value, ttl)
	return true
}

func Get(key string) string {
	if len(key) <= 0 {
		return ""
	}
	g := Run()
	value, found := g.Get(key)
	if !found {
		return ""
	}
	return value.(string)
}
