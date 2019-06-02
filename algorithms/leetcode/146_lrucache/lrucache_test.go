package lrucache

import "testing"

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)

	cache.put(1, 1)
	cache.put(2, 2)
	if value := cache.get(1); value != 1 {
		t.Fatalf("expected 1 but got %d", value)
	}
	cache.put(3, 3) // evicts key 2
	if value := cache.get(2); value != -1 {
		t.Fatalf("expected -1 but got %d", value)
	}
	cache.put(4, 4) // evicts key 1
	if value := cache.get(1); value != -1 {
		t.Fatalf("expected -1 but got %d", value)
	}
	if value := cache.get(3); value != 3 {
		t.Fatalf("expected -1 but got %d", value)
	}
	if value := cache.get(4); value != 4 {
		t.Fatalf("expected 4but got %d", value)
	}
}
