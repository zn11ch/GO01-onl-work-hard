package main

import (
	"fmt"
	"testing"
)

func TestGetSeason(t *testing.T) {
	var tests = []struct {
		num  int
		want string
	}{
		{1, "Зима"},
		{2, "Весна"},
		{3, "Лето"},
		{4, "Осень"},
		{5, ""},
		{0, ""},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d %v", tt.num, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans, _ := GetSeason(tt.num)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
