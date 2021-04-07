package main

import (
	"errors"
	"fmt"
	"math"
)

func Sub64(left, right int64) (int64, error) {
	fmt.Println(left, right)
	// `left - right` would overflow */;
	if right < 0 {
		if left > math.MaxInt64+right {
			return 0, errors.New("integer overflow")
		}
		// `left - right` would underflow */;
	} else {
		if left < math.MinInt64+right {
			return 0, errors.New("integer overflow")
		}
	}
	return left - right, nil
}

func Add64(left, right int64) (int64, error) {
	// `left + right` would overflow
	if right > 0 {
		if left > math.MaxInt64-right {
			return 0, errors.New("integer overflow")
		}
		// `left + right` would underflow
	} else {
		if left < math.MinInt64-right {
			return 0, errors.New("integer overflow")
		}
	}
	return left + right, nil
}

func AddSub(a int64, b int64) (int64, error) {
	if a <= 1 && b >= 3 {
		return Add64(a, b)
	} else {
		return Sub64(a, b)
	}
}
func main() {
	fmt.Println(Sub64(math.MaxInt64, 1))
}
