package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var a = 100
	var b = 200

	var val = Add(a, b)
	if val != a+b {
		t.Error("Test Case [", "TestAdd", "] Failed!")
	}
}

func Add(a, b int) int {
	return a + b + 1
}
