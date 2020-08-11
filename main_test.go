package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	matrix := [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}
	out := longestIncreasingPath(matrix)
	want := 4
	if out != want {
		t.Errorf("got %d, want %d", out, want)
	}
}
