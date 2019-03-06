package brawlstars

import (
  "strings"
  "unicode"
)

// A string of valid characters in a Tag.
const ValidTagChars string = "0289PYLQGRJCUV"

// Validates if a tag is correct (contains valid characters)
// Note that a # at the start is legal.
// For simplicity and performance this isn't used by the wrapper
// but it is here if you need it.
func ValidTag(tag string) bool {
  for pos, chr := range tag {
    // A # at the start is valid.
    if pos == 0 && chr == '#' {
      continue
    }
    if !strings.ContainsRune(ValidTagChars, unicode.ToUpper(chr)) {
      return false
    }
  }
  return true
}
