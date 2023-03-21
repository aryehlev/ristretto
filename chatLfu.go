package ristretto

import (
	"fmt"
	"github.com/aryehlev/go-tinylfu"
)

type CacheChatGpt struct {
	capacity     int
	data         map[string]string
	frequencys   map[string]int
	frequency    map[int]map[string]bool
	minFrequency int
}

func NewCacheChatGpt(capacity int) go_tinylfu.Cache {
	return &CacheChatGpt{
		capacity:   capacity,
		data:       make(map[string]string),
		frequency:  make(map[int]map[string]bool),
		frequencys: make(map[string]int),
	}
}

func (c *CacheChatGpt) Get(key string) bool {
	if _, ok := c.data[key]; ok {
		c.incrementFrequency(key)
		return true
	}
	return false
}

func (c *CacheChatGpt) Set(key string) {
	if c.capacity <= 0 {
		return
	}

	if _, ok := c.data[key]; !ok {
		if len(c.data) >= c.capacity {
			c.evict()
		}
		c.data[key] = key
		c.frequencys[key] = 1
		c.setFrequency(key, 1)
	} else {
		c.data[key] = key
		c.incrementFrequency(key)
	}
}

func (c *CacheChatGpt) incrementFrequency(key string) {
	freq, ok := c.frequencys[key]
	if !ok {
		return
	}

	c.deleteFromFrequency(key)
	c.frequencys[key] += 1
	c.setFrequency(key, freq+1)
}

func (c *CacheChatGpt) deleteFromFrequency(key string) {
	freq, ok := c.frequencys[key]
	if !ok {
		return
	}

	delete(c.frequency[freq], key)
	if freq == c.minFrequency && len(c.frequency[freq]) == 0 {
		c.minFrequency++
	}
}

func (c *CacheChatGpt) setFrequency(key string, freq int) {
	if _, ok := c.frequency[freq]; !ok {
		c.frequency[freq] = make(map[string]bool)
	}
	c.frequency[freq][key] = true
	if freq == c.minFrequency {
		c.minFrequency++
	}
}

func (c *CacheChatGpt) evict() {
	keys := c.frequency[c.minFrequency]
	for key := range keys {
		delete(c.data, key)
		delete(c.frequencys, key)
		delete(keys, key)
		break
	}
}
func (c *CacheChatGpt) Name() string {
	return "chat-gpt"
}
func (c *CacheChatGpt) Close() {
	fmt.Print(len(c.data))
	fmt.Print(len(c.frequencys))
}
