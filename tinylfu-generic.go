package ristretto

import (
	"github.com/aryehlev/go-tinylfu"
	"sync"
)

type TinyLFUGeneric struct {
	mu sync.Mutex
	v  *T[string]
}

func NewTinyLFUGeneric(size int) go_tinylfu.Cache {
	return &TinyLFUGeneric{
		v: New[string](size, size*10, true),
	}
}

func (c *TinyLFUGeneric) Name() string {
	return "tinylfu-generic"
}

func (c *TinyLFUGeneric) AllKeys() int {
	return c.v.AllKeys()
}

func (c *TinyLFUGeneric) AllSizes() int {
	return c.v.AllSizes()
}

func (c *TinyLFUGeneric) AllCaps() int {
	return c.v.AllCaps()
}

//func (c *TinyLFUGeneric) IsSameIsSame() bool {
//	return c.v.IsSame()
//}

func (c *TinyLFUGeneric) Set(key string) {
	c.mu.Lock()
	c.v.Add(key, key)
	c.mu.Unlock()
}

func (c *TinyLFUGeneric) Get(key string) bool {
	c.mu.Lock()
	_, ok := c.v.Get(key)
	c.mu.Unlock()
	return ok
}

func (c *TinyLFUGeneric) Close() {}
