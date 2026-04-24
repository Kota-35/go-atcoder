package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		l    int
		r    int
		want int
	}{
		{3, 5, 3},
		{1, 7, 7},
		{14, 79, 66},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, solve(tt.l, tt.r))
	}
}
