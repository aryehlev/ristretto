package main

import (
	"fmt"
	"github.com/dgraph-io/ristretto"
	"runtime"
	"sync"
	"time"
)

type Benchmark struct {
	ristretto.Generator
	N int
}

func main() {
	cacheSize := []int{1e3, 10e3, 100e3, 1e6}
	multiplier := []int{10, 100, 1000, 10000, 1000000}
	newCache := []NewCacheFunc{
		//NewHashicorpLRUV2,
		ristretto.NewTinyLFUGeneric,
		ristretto.NewTinyLFUGeneric1,
		//NewCacheChatGpt,
	}
	newGen := []ristretto.NewGeneratorFunc{
		//NewTestAryeh,
		ristretto.NewScrambledZipfian,
		ristretto.NewHotspot,
		ristretto.NewUniform,
	}

	var results []*ristretto.BenchmarkResult

	for _, newGen := range newGen {
		for _, cacheSize := range cacheSize {
			for _, multiplier := range multiplier {
				numKey := cacheSize * multiplier

				if len(results) > 0 {
					ristretto.printResults(results)
					results = results[:0]
				}

				for _, newCache := range newCache {
					result := run(newGen, cacheSize, numKey, newCache)
					results = append(results, result)
				}
			}

		}

	}
}

func run(newGen ristretto.NewGeneratorFunc, cacheSize, numKey int, newCache NewCacheFunc) *ristretto.BenchmarkResult {
	gen := newGen(numKey)
	b := &Benchmark{
		Generator: gen,
		N:         1e6,
	}

	alloc1 := memAlloc()
	cache := newCache(cacheSize)
	defer cache.Close()

	start := time.Now()
	hits, misses := bench(b, cache)
	dur := time.Since(start)

	alloc2 := memAlloc()

	return &ristretto.BenchmarkResult{
		GenName:   gen.Name(),
		CacheName: cache.Name(),
		CacheSize: cacheSize,
		NumKey:    numKey,

		Hits:     hits,
		Misses:   misses,
		Duration: dur,
		Bytes:    int64(alloc2) - int64(alloc1),
	}
}

func bench(b *Benchmark, cache Cache) (hits, misses int) {
	wg := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		value := b.Next()
		//wg.Add(1)
		//go func() {
		//	defer wg.Done()
		if cache.Get(value) {
			hits++
		} else {
			misses++
			cache.Set(value)
		}
		//}()

	}

	tiny, ok := cache.(*ristretto.TinyLFUGeneric)

	if ok {
		fmt.Println(tiny.AllSizes())
		fmt.Println(tiny.AllKeys())
		fmt.Println(tiny.AllCaps())
		//fmt.Println(tiny.IsSameIsSame())
	}

	tiny1, ok := cache.(*ristretto.TinyLFUGeneric1)

	if ok {
		fmt.Println(tiny1.AllSizes())
		fmt.Println(tiny1.AllKeys())
		fmt.Println(tiny1.AllCaps())
		//fmt.Println(tiny.IsSameIsSame())
	}

	wg.Wait()
	return hits, misses
}

func memAlloc() uint64 {
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}
