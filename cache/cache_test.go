package pokecache

import (
	"testing"
	"time"
)

func TestAddAndGet(t *testing.T) {
	cache := NewCache(5 * time.Second) // reap interval irrelevant for this test

	key := "pikachu"
	value := []byte("electric-mouse")

	cache.Add(key, value)

	got, ok := cache.Get(key)
	if !ok {
		t.Errorf("Expected key %q to be in cache, but it was not found", key)
	}
	if string(got) != string(value) {
		t.Errorf("Expected value %q, but got %q", string(value), string(got))
	}
}

func TestGetNonexistentKey(t *testing.T) {
	cache := NewCache(5 * time.Second)

	_, ok := cache.Get("missing")
	if ok {
		t.Error("Expected Get() to return false for missing key, but got true")
	}
}