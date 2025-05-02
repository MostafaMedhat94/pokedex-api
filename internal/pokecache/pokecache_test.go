package pokecache

import (
	"testing"
	"time"
)

type testCase struct {
	input    map[string][]byte
	expected []byte
}

func TestAdd(t *testing.T) {
	cases := []testCase{
		{
			input:    map[string][]byte{"foo": []byte("bar")},
			expected: []byte("bar"),
		},
		{
			input:    map[string][]byte{"ping": []byte("pong")},
			expected: []byte("pong"),
		},
		{
			input:    map[string][]byte{"hakuna": []byte("matata")},
			expected: []byte("matata"),
		},
	}

	cache := NewCache(10 * time.Second)

	for _, c := range cases {
		var key string
		var val []byte

		for key, val = range c.input {
			cache.Add(key, val)
		}

		actual, ok := cache.Get(key)

		if !ok {
			t.Errorf("Expected %v, got %v", c.expected, actual)

			t.Fail()
		}
	}
}
