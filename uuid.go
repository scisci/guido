package guido

import (
	"crypto/rand"
	"fmt"
	"strings"
)

func GenerateUUID4() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	b[6] = (b[6] & 0x0f) | 0x40 // Enforce part 3 starts with '4'
	b[8] = (b[8] & 0x3f) | 0x80 // Enforce part 4 starts with '8','9','a', or 'b'
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func NormalizeUUID4(uuid string) string {
	return strings.ToLower(uuid)
}
