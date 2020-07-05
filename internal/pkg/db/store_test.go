package db

import (
	"testing"
)

func TestEmbeddedKeyValueStore(t *testing.T) {
	keyValue, err := EmbeddedKeyValueStore(nil)
	if err != nil {
		t.Error(err)
	}

	key := []byte("hello")
	value := []byte("world")

	err = keyValue.Set(key, value)
	if err != nil {
		t.Error(err)
	}

	got, err := keyValue.Get(key)
	if err != nil {
		t.Error(err)
	}

	gotValue := string(got)

	if gotValue != "world" {
		t.Errorf("got %s, want %s", got, string(value))
	}

	keyValue.Close()
}
