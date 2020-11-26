package fib

import (
	"testing"
)

func TestFib(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{
			name: "Check negative",
			num:  -1,
			want: 0,
		},
		{
			name: "Check zero",
			num:  0,
			want: 0,
		},
		{
			name: "Check one",
			num:  1,
			want: 1,
		},
		{
			name: "Check 6 value",
			num:  6,
			want: 8,
		},
		{
			name: "Check 10 value",
			num:  10,
			want: 55,
		},
		{
			name: "Check 20 value",
			num:  20,
			want: 6765,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Fib(tt.num)

			if got != tt.want {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
