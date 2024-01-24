package arand

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRand(t *testing.T) {
	assert.True(t, len(Rand()) > 0)
	assert.True(t, len(RandLetterN(1)) > 0)
	assert.True(t, len(RandLetterDigitN(1)) > 0)
	assert.True(t, len(RandDigitN(1)) > 0)
	assert.True(t, len(RandDigitNoZeroN(1)) > 0)

	const size = 10
	assert.True(t, len(RandLetterN(size)) == size)
	assert.True(t, len(RandLetterDigitN(size)) == size)
	assert.True(t, len(RandDigitN(size)) == size)
	assert.True(t, len(RandDigitNoZeroN(size)) == size)
	assert.True(t, RandWithPoolN(size, "a") == "aaaaaaaaaa")

	assert.True(t, regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(RandLetterN(size)))
	assert.True(t, regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(RandLetterDigitN(size)))
	assert.True(t, regexp.MustCompile(`^[0-9]+$`).MatchString(RandDigitN(size)))
	assert.True(t, regexp.MustCompile(`^[1-9]+$`).MatchString(RandDigitNoZeroN(size)))
}

func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandLetterN(1)
		_ = RandLetterN(10)
	}
}
