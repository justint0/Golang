/*
A Go string is a read-only slice of bytes. The language and the standard library treats strings specially - as containers of text encoded in UTF-8. In other languages, strings are made of "characters". In Go, the concept of a character is called a rune - it's an integer that represents a Unicode code point.
*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

		// This line defines a string s containing the Thai word "สวัสดี" (hello). Each character might be represented by multiple bytes.
    const s = "สวัสดี"

		// This prints the length of the string s. Since strings are treated as byte slices ([]byte), len(s) returns the number of bytes in the string, not the number of characters.
    fmt.Println("Len:", len(s))

		// This loop iterates over the string byte by byte, printing the hexadecimal representation of each byte. This shows that the Thai characters are composed of multiple bytes.
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Println()

		// This uses the utf8.RuneCountInString function to count the actual number of Unicode characters (runes) in the string.
    fmt.Println("Rune count:", utf8.RuneCountInString(s))

		// This loop demonstrates the idiomatic way to iterate over runes in a string. The range keyword, when used with a string, automatically decodes the UTF-8 encoded string into runes. For each rune, it prints the Unicode code point (%#U) and the starting byte index (%d) within the string.
    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }

		// This loop shows how to manually decode runes from a string using utf8.DecodeRuneInString. This function takes a string and returns the first rune and its width in bytes. The loop then uses the width to advance to the next rune.
    fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width

        examineRune(runeValue)
    }
}

// This function takes a rune as input and checks if it's either the English letter 't' or the Thai character 'ส'. This demonstrates how you can work with individual runes.
func examineRune(r rune) {
    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}