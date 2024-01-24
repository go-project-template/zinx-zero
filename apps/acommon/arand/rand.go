package arand

import (
	"math/rand"
	"time"
)

var _rand = rand.New(rand.NewSource(time.Now().UnixNano()))

const (
	letterBytes        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterBytesNoDigit = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitBytes         = "0123456789"
	digitBytesNoZero   = "123456789"
	letterIdxBits      = 6 // 6 bits to represent a letter index
	idLen              = 8
	defaultRandLen     = 8
)

func Intn(n int) int {
	return _rand.Intn(n)
}

func Int31n(n int32) int32 {
	return _rand.Int31n(n)
}

func Int63n(n int64) int64 {
	return _rand.Int63n(n)
}

// Rand returns a random string.
func Rand() string {
	return RandLetterN(defaultRandLen)
}

// RandLetterN returns a random string with length n.
func RandLetterN(n int) string {
	return RandWithPoolN(n, letterBytesNoDigit)
}

// RandLetterDigitN returns a random string with length n.
func RandLetterDigitN(n int) string {
	return RandWithPoolN(n, letterBytes)
}

// RandDigitN returns a random string with length n.
func RandDigitN(n int) string {
	return RandWithPoolN(n, digitBytes)
}

// RandDigitNoZeroN returns a random string with length n.
func RandDigitNoZeroN(n int) string {
	return RandWithPoolN(n, digitBytesNoZero)
}

// RandWithPoolN returns a random string with length n.
func RandWithPoolN(n int, _pool string) string {
	pool := []rune(_pool)
	res := make([]rune, n)
	for i := 0; i < n; i++ {
		res[i] = pool[_rand.Intn(len(pool))]
	}
	return string(res)
}
