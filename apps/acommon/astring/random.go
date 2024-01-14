package astring

import (
	"math/rand"
	"sync"
	"time"
)

const (
	letterBytes        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterBytesNoDigit = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitBytes         = "0123456789"
	digitBytesNoZero   = "123456789"
	letterIdxBits      = 6 // 6 bits to represent a letter index
	idLen              = 8
	defaultRandLen     = 8
	letterIdxMask      = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax       = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = newLockedSource(time.Now().UnixNano())

type lockedSource struct {
	source rand.Source
	lock   sync.Mutex
}

func newLockedSource(seed int64) *lockedSource {
	return &lockedSource{
		source: rand.NewSource(seed),
	}
}

func (ls *lockedSource) Int63() int64 {
	ls.lock.Lock()
	defer ls.lock.Unlock()
	return ls.source.Int63()
}

func (ls *lockedSource) Seed(seed int64) {
	ls.lock.Lock()
	defer ls.lock.Unlock()
	ls.source.Seed(seed)
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
func RandWithPoolN(n int, pool string) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(pool) {
			b[i] = pool[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// Seed sets the seed to seed.
func Seed(seed int64) {
	src.Seed(seed)
}
