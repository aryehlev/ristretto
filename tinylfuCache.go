package ristretto

//
//import (
//	"strconv"
//	"sync"
//)
//
//type TinyLFUChat struct {
//	mu sync.Mutex
//	v  *CacheChatGpt
//}
//
//func NewTinyLFUChat(size int) Cache {
//	return &TinyLFUChat{
//		v: NewCacheChatGpt(size),
//	}
//}
//
//func (c *TinyLFUChat) Name() string {
//	return "tinylfu-chat"
//}
//
////func (c *TinyLFUGeneric) IsSameIsSame() bool {
////	return c.v.IsSame()
////}
//
//func (c *TinyLFUChat) Set(key string) {
//	c.mu.Lock()
//	c.v.Set(key, strconv.Atoi(key))
//	c.mu.Unlock()
//}
//
//func (c *TinyLFUChat) Get(key string) bool {
//	c.mu.Lock()
//	_, ok := c.v.Get(key)
//	c.mu.Unlock()
//	return ok
//}
//
//func (c *TinyLFUChat) Close() {}
