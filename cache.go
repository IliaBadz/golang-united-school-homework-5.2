package cache

import (
	"time"
)

type Value struct {
	val      string
	deadline time.Time
}

type Cache struct {
	pair map[string]Value
}

func NewCache() Cache {
	return Cache{make(map[string]Value)}
}

func (c Cache) Get(key string) (string, bool) {

	if res, ok := c.pair[key]; ok {
		if time.Now().Before(res.deadline) || res.deadline.IsZero() {
			return res.val, true
		} else {
			delete(c.pair, key)
		}
	}

	return "", false
}

func (c Cache) Put(key, value string) {
	c.pair[key] = Value{val: value}
}

func (c Cache) Keys() []string {
	res := make([]string, 0, len(c.pair))

	for k := range c.pair {
		res = append(res, k)
	}

	return res
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.pair[key] = Value{val: value, deadline: deadline}
}
