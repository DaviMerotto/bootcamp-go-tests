package gotest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Fibonacci 0",
			n:    0,
			want: 0,
		},
		{
			name: "Fibonacci 1",
			n:    1,
			want: 1,
		},
		{
			name: "Fibonacci 2",
			n:    2,
			want: 1,
		},
		{
			name: "Fibonacci 3",
			n:    3,
			want: 2,
		},
		{
			name: "Fibonacci 4",
			n:    4,
			want: 3,
		},
		{
			name: "Fibonacci 5",
			n:    5,
			want: 5,
		},
		{
			name: "Fibonacci 6",
			n:    6,
			want: 8,
		},
		{
			name: "Fibonacci 7",
			n:    7,
			want: 13,
		},
		{
			name: "Fibonacci 8",
			n:    8,
			want: 21,
		},
		{
			name: "Fibonacci 9",
			n:    9,
			want: 34,
		},
		{
			name: "Fibonacci 10",
			n:    10,
			want: 55,
		},
		{
			name: "Fibonacci 11",
			n:    11,
			want: 89,
		},
		{
			name: "Fibonacci 12",
			n:    12,
			want: 144,
		},
		{
			name: "Fibonacci 13",
			n:    13,
			want: 233,
		},
		{
			name: "Fibonacci 14",
			n:    14,
			want: 377,
		},
		{
			name: "Fibonacci 15",
			n:    15,
			want: 610,
		},
		{
			name: "Fibonacci 16",
			n:    16,
			want: 987,
		},
	}
	for _, tt := range tests {
		fib := FibonacciIterative(tt.n)
		assert.Equal(t, tt.want, fib, tt.name)
	}
}
