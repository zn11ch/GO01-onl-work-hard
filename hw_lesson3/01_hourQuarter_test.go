package main

import (
	"fmt"
	"testing"
)

func TestHourQuarter(t *testing.T) {
	var tests = []struct {
		minutes uint8
		want    uint8
	}{
		{0, 1},
		{1, 1},
		{15, 2},
		{16, 2},
		{30, 3},
		{31, 3},
		{45, 4},
		{59, 4},
		{60, 0},
		{255, 0},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.minutes)
		t.Run(testname, func(t *testing.T) {
			ans, _ := hourQuarter(tt.minutes)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
