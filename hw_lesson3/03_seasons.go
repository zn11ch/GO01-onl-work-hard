package main

import (
	"errors"
)

var Seasons = map[int]string{
	1: "Зима",
	2: "Весна",
	3: "Лето",
	4: "Осень",
}

func GetSeason(num int) (string, error) {
	if val, ok := Seasons[num]; ok {
		return val, nil
	}
	return "", errors.New("unknown season")
}
