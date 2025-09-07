package emojitoolkit

import (
	_ "embed"
	"encoding/binary"
)

// Supported Unicode version
const Version = "16.0.0"

/*type Emoji struct {
	Rune         rune
	Name         string
	IntroducedIn string
}*/

//go:embed emoji_ranges.bin
var emoji_ranges []byte

func IsSingleCharacterEmoji(r rune) bool {
	for i := 0; i < len(emoji_ranges); i += 8 {
		min := binary.LittleEndian.Uint32(emoji_ranges[i:])
		max := binary.LittleEndian.Uint32(emoji_ranges[i+4:])
		if r >= rune(min) && r <= rune(max) {
			return true
		}
	}
	return false
}
