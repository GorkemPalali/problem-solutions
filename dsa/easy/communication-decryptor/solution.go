package communicationdecryptor

import (
	"strings"
)

func DecryptClientMessage(encryptedMessage string, decryptionMap map[rune]rune) string {
	var b strings.Builder
	b.Grow(len(encryptedMessage))

	for _, ch := range encryptedMessage {
		switch {
		case ch >= 'a' && ch <= 'z':
			if orig, ok := decryptionMap[ch]; ok {
				b.WriteRune(orig)
			} else {
				b.WriteRune(ch)
			}
		case ch >= 'A' && ch <= 'Z':
			lower := ch + ('a' - 'A')
			if orig, ok := decryptionMap[lower]; ok {
				upper := orig - ('a' - 'A')
				b.WriteRune(upper)
			} else {
				b.WriteRune(ch)
			}
		default:
			b.WriteRune(ch)
		}
	}
	return b.String()
}
