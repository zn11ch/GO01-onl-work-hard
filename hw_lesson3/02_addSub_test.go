package main

import (
	"fmt"
	"math"
	"testing"
)

func TestSub64(t *testing.T) {
	var tests = []struct {
		a    int64
		b    int64
		want int64
	}{
		{0, 0, 0},
		{math.MaxInt64, 1, 9223372036854775806},
		{math.MaxInt64, math.MaxInt64, 0},
		{math.MaxInt64, -1, 0},
		{math.MaxInt64, math.MinInt64, 0},
		{math.MinInt64, -1, -9223372036854775807},
		{math.MinInt64, math.MaxInt64, 0},
		{math.MinInt64, 1, 0},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d %d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans, _ := Sub64(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestAdd64(t *testing.T) {
	var tests = []struct {
		a    int64
		b    int64
		want int64
	}{
		{0, 0, 0},
		{math.MaxInt64, 1, 0},
		{math.MaxInt64, math.MaxInt64, 0},
		{math.MaxInt64, -1, 9223372036854775806},
		{math.MaxInt64, math.MinInt64, -1},
		{math.MinInt64, -1, 0},
		{math.MinInt64, math.MaxInt64, -1},
		{math.MinInt64, 1, -9223372036854775807},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d %d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans, _ := Add64(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestAddSub(t *testing.T) {
	var tests = []struct {
		a    int64
		b    int64
		want int64
	}{
		{0, 0, 0},
		{1, 3, 4},
		{2, 4, -2},
		{0, 3, 3},
		{2, 2, 0},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d %d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans, _ := AddSub(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
