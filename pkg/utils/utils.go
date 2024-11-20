package utils

import nanoid "github.com/matoous/go-nanoid/v2"

func Ternary[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

func GenID() string {
	return nanoid.MustGenerate("1234567890abcdefghijklmnopqrstuvwxyz", 12)
}
