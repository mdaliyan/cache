package cache_test

import (
	"testing"
	"time"

	"github.com/akyoto/cache"
)

func Test(t *testing.T) {
	cycle := 20 * time.Millisecond
	c := cache.New(cycle)
	defer c.Close()

	c.Set("Hello", "World", cycle/2)
	hello, found := c.Get("Hello")

	if !found {
		t.FailNow()
	}

	if hello.(string) != "World" {
		t.FailNow()
	}

	time.Sleep(cycle / 2)

	_, found = c.Get("Hello")

	if found {
		t.FailNow()
	}

	time.Sleep(cycle)

	_, found = c.Get("404")

	if found {
		t.FailNow()
	}
}

func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache.New(5 * time.Second).Close()
		}
	})
}

func BenchmarkGet(b *testing.B) {
	c := cache.New(5 * time.Second)
	defer c.Close()
	c.Set("Hello", "World", 0)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Get("Hello")
		}
	})
}

func BenchmarkSet(b *testing.B) {
	c := cache.New(5 * time.Second)
	defer c.Close()

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Set("Hello", "World", 0)
		}
	})
}
