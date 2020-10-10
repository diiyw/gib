package gache

import (
	"sync"
	"time"
)

// 简单的内存缓存实现
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
	return &Cache{&sync.Map{}}
}

// 设置值
func (c *Cache) Set(key string, value interface{}) {
	c.keyValue.Store(key, &data{time.Now(), -1, value})
}

func (c *Cache) SetEx(key string, value interface{}, expire time.Duration) {
	c.keyValue.Store(key, &data{time.Now(), expire, value})
}

// 取值
func (c *Cache) Get(key string) interface{} {

	if v, ok := c.keyValue.Load(key); ok {
		if v.(*data).created.Add(v.(*data).expired).Before(time.Now()) && v.(*data).expired != -1 {
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
		if v.(*data).expired == -1 {
			return true
		}
		if v.(*data).created.Add(v.(*data).expired).Before(time.Now()) {
			// 过期了
			c.keyValue.Delete(key)
			return false
		}
		return v.(*data).value != nil
	}
	return false
}

// 获取键值的过期时间
func (c *Cache) GetTTL(key string) int {
	if v, ok := c.keyValue.Load(key); ok {

		left := v.(*data).created.Add(v.(*data).expired).Sub(time.Now()).Seconds()
		return int(left)
	}
	return 0
}
