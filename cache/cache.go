package cache

import (
	"sync"
	"time"
)

type data struct {
	// 过期时间
	created time.Time
	expired time.Duration
	value   interface{}
}

type Cache struct {
	// 缓存的键值
	keyValue *sync.Map
}

func New() *Cache {
	return &Cache{&sync.Map{},}
}

// 设置值
func (c *Cache) Set(key string, value interface{}) {
	c.keyValue.Store(key, &data{time.Now(), -1, value})
}

// 取值
func (c *Cache) Get(key string) interface{} {

	if v, ok := c.keyValue.Load(key); ok {
		if v.(*data).created.Add(v.(*data).expired).Before(time.Now()) {
			c.keyValue.Delete(key)
			return nil
		}
		return v.(*data).value
	}
	return nil

}

// 设置过期时间
func (c *Cache) Expired(key string, t time.Duration) {
	if t <= 0 {
		c.keyValue.Delete(key)
	}
	if v, ok := c.keyValue.Load(key); ok {
		v.(*data).expired = t
		c.keyValue.Store(key, v)
	}
}

// 值是否存在
func (c *Cache) Exits(key string) bool {

	if v, ok := c.keyValue.Load(key); ok {
		if v.(*data).created.Add(v.(*data).expired).Before(time.Now()) {
			// 过期了
			c.keyValue.Delete(key)
			return false
		}
		return v.(*data).value == nil
	}
	return false
}