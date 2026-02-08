package main

import (
	"errors"
	"strconv"
	"strings"
)

const (
	coefFizz = 3
	coefBuzz = 5

	msgFizz = "Fizz"
	msgBuzz = "Buzz"
)

var (
	ErrNegativeNumber = errors.New("negative number")
)

func fizzBuzz(n int) (string, error) {
	if n < 0 {
		return "", ErrNegativeNumber
	}

	isFizz := n%coefFizz == 0
	isBuzz := n%coefBuzz == 0

	if !(isFizz || isBuzz) {
		return strconv.Itoa(n), nil
	}

	var sb strings.Builder
	sb.Grow(len(msgFizz) + len(msgBuzz))
	if isFizz {
		sb.WriteString(msgFizz)
	}
	if isBuzz {
		sb.WriteString(msgBuzz)
	}
	return sb.String(), nil
}
