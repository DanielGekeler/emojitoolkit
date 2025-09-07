package main

import (
	emojitoolkit "emoji-toolkit"
	"fmt"
)

// https://www.unicode.org/reports/tr42/
// https://www.unicode.org/Public/16.0.0/ucdxml/ucd.nounihan.flat.zip

func main() {
	fmt.Println(emojitoolkit.IsSingleCharacterEmoji('🥇'))
}
