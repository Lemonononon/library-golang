package crypto

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

// SHA1 return the sha1(string), warning: it will transform string to upper
func SHA1(content string) string {
	h := sha1.New()
	h.Write([]byte(content))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
