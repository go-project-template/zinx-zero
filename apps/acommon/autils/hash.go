package autils

import (
	"github.com/zeromicro/go-zero/core/hash"
)

// Md5HexByString returns the md5 hex string of data.
func Md5HexByString(param string) string {
	return hash.Md5Hex([]byte(param))
}
