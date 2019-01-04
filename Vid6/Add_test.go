package main

import (
	"testing"
	"fmt"
)

func TestAdd(t *testing.T) {

	var testCases = []struct {
		input1   int
		input2   int
		expected int
	}{
		{2, 3, 5},
		{-2, 3, 1},
		{0, 3, 3},
		{2333, 3, 2336},
	}

	for _, v := range testCases {
		total := Add(v.input1, v.input2)
		if total != v.expected {
			t.Errorf("Sum was incorrect , got %v , want %v", total, v.expected)
		}
	}

}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(Add(2, 3))
	}
}