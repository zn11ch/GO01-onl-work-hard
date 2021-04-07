package main

import (
	"errors"
)

func hourQuarter(minutes uint8) (uint8, error) {
	if minutes < 0 {
		return 0, errors.New("Minutes should be above zero")
	}
	if minutes > 59 {
		return 0, errors.New("Minutes should be less than 59")
	}
	return minutes/15 + 1, nil
}
