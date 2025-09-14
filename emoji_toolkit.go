package emojitoolkit

import (
	_ "embed"
	"encoding/binary"
)

// Supported Unicode version
const Version = "16.0.0"

//go:generate go run generator/main.go

//go:embed emoji_ranges.bin
var emoji_ranges []byte

func IsSingleCharacterEmoji(r rune) bool {
	return isInRange(r, emoji_ranges)
}

//go:embed emoji_ranges2.bin
var emoji_ranges2 []byte

// Check if a string contains a basic emoji defined by [ED-20]
//
// [ED-20]: https://www.unicode.org/reports/tr51/#def_basic_emoji_set
func ContainsEmoji(s string) bool {
	runes := []rune(s)
	l := len(runes)
	for i, r := range runes {
		if IsSingleCharacterEmoji(r) {
			return true
		}

		if i+1 < l && isInRange(r, emoji_ranges2) && runes[i+1] == '\uFE0F' {
			return true
		}
	}
	return false
}

// Check if a rune is within any range of a set of ranges.
// These sets are an array of int32 with each range being defined
// as two consecutive int32 defining the first and last element
// of a range respectively.
func isInRange(r rune, ranges []byte) bool {
	for i := 0; i < len(ranges); i += 8 {
		min := binary.LittleEndian.Uint32(ranges[i:])
		max := binary.LittleEndian.Uint32(ranges[i+4:])
		if r >= rune(min) && r <= rune(max) {
			return true
		}
	}
	return false
}
