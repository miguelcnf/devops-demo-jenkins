package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Run("should return Hello World", func(t *testing.T) {
		res := HelloWorld()
		if res != "Hello World" {
			t.Errorf("expected string to be Hello World but was: %s", res)
		}
	})
}
