# Emoji toolkit
A package to check if runes are Emojis or strings contain Emojis and flags.

The standard [Unicode](https://pkg.go.dev/unicode) package does not
provide any functionality to work with Emojis as of go 1.25.

## How Emojis work
Emojis are a complicated feature of the Unicode Standard.
A simple Emoji like ✨ (U+2728 SPARKLES) is repressented as a single codepoint or rune in go.
This character has the properties `Emoji=Yes` and `Emoji_Presentation=Yes` according to the Unicode Character Database (UCD).
`Emoji_Presentation=Yes` means that this character will appear as a Emoji by default.
See section [Emoji Character Properties of UTS #51](https://www.unicode.org/reports/tr51/#Emoji_Properties) for further information.

Some characters can appear as Emoji but appear as text by default.
A example of such a character is ⛓ (U+26D3 CHAINS) because it has the properties `Emoji=Yes` and `Emoji_Presentation=No`.
Putting U+FE0F VARIATION SELECTOR-16 after the character forces it to appear as an Emoji
which looks like this: ⛓️ (U+26D3 CHAINS U+FE0F VARIATION SELECTOR-16).
*Note: This requires two codepoints or runes to store the Emoji.*

Emoji Keycap Sequences like #️⃣*️⃣0️⃣1️⃣2️⃣3️⃣ are not seperate characters.
They are made using the ASCII # (U+0023), \* (U+002A) and digits (U+0030 .. U+0039)
followed by U+FE0F VARIATION SELECTOR-16 and U+20E3 COMBINING ENCLOSING KEYCAP.
A spacial case is 🔟 (U+1F51F KEYCAP TEN) which is a independent character with
the properties `Emoji=Yes` and `Emoji_Presentation=Yes` making it appear as an Emoji by default.

Flags 🇺🇳 🇪🇺 🇩🇪 also require two codepoints. This is because they are repressented by a second alphabet,
the regional indicator symbols, that is used to write a two letter country code.
Example: 🇩🇪 (U+1F1E9 REGIONAL INDICATOR SYMBOL LETTER A U+1F1EA REGIONAL INDICATOR SYMBOL LETTER E).

## Features
- Unicode Standard 1.16
- Detects all Emojis listed in emoji-sequences.txt.
- Detect Emojis in a single rune (only default emoji presentation character)

## Development
Download [ucd.nounihan.flat.zip](https://www.unicode.org/Public/16.0.0/ucdxml/) and place `ucd.nounihan.flat.xml` in the repository root.

## References
- [Unicode Character Database in XML (UTS #42)](https://www.unicode.org/reports/tr42/)
- [Unicode Character Database (UTS #44)](https://www.unicode.org/reports/tr44/)
- [Unicode Emoji (UTS #51)](https://www.unicode.org/reports/tr51/)
- [Glossary of Unicode Terms](https://www.unicode.org/glossary/)
- [emoji-sequences.txt](https://www.unicode.org/Public/emoji/latest/emoji-sequences.txt)

## License
Copyright 2025 Daniel Gekeler

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
