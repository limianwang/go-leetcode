package lrucache_test

import (
	"testing"

	"github.com/limianwang/goexamples/lrucache"
)

func TestNewCache(t *testing.T) {
	c := lrucache.NewCache(2)

	c.Set("a", 10)
	val := c.Get("a")
	if val != 10 {
		t.Fatalf("failed, expected %d, got %v", 10, val)
	}

	c.Set("b", 11)
	val2 := c.Get("b")
	val1 := c.Get("a")
	if val1 != 10 {
		t.Fatalf("failed, expected %d, got %v", 10, val1)
	}
	if val2 != 11 {
		t.Fatalf("failed, expected %d, got %v", 11, val2)
	}

	c.Set("c", 12)
	val3 := c.Get("c")
	val1 = c.Get("b")

	if val1 != -1 {
		t.Fatalf("expected %d but got %v", -1, val1)
	}
	if val3 != 12 {
		t.Fatalf("failed, expected %d, got %v", 12, val3)
	}
}
