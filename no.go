package ristretto

import (
	"github.com/aryehlev/go-tinylfu"
	"sync"
)

type TinyLFUGeneric1 struct {
	mu sync.Mutex
	v  *T[string]
}

func NewTinyLFUGeneric1(size int) go_tinylfu.Cache {
	return &TinyLFUGeneric1{
		v: New[string](size, size*10, false),
	}
}

func (c *TinyLFUGeneric1) Name() string {
	return "tinylfu-no-resize"
}

func (c *TinyLFUGeneric1) AllKeys() int {
	return c.v.AllKeys()
}

func (c *TinyLFUGeneric1) AllSizes() int {
	return c.v.AllSizes()
}

func (c *TinyLFUGeneric1) AllCaps() int {
	return c.v.AllCaps()
}

//func (c *TinyLFUGeneric1) IsSameIsSame() bool {
//	return c.v.IsSame()
//}

func (c *TinyLFUGeneric1) Set(key string) {
	c.mu.Lock()
	c.v.Add(key, key)
	c.mu.Unlock()
}

func (c *TinyLFUGeneric1) Get(key string) bool {
	c.mu.Lock()
	_, ok := c.v.Get(key)
	c.mu.Unlock()
	return ok
}

func (c *TinyLFUGeneric1) Close() {}
