package master

// SubstringIndexes returns the set of all character indexes in text that fall
// inside any occurrence of substring. The returned map is keyed by index;
// a missing key means that character is not part of any match.
//
// Example:
//
//	SubstringIndexes("a king kitten have kit", "kit")
//	→ {7:true, 8:true, 9:true, 19:true, 20:true, 21:true}
func SubstringIndexes(text, substring string) map[int]bool {
	result := make(map[int]bool)
	subLen := len(substring)
	if subLen == 0 {
		return result
	}
	for i := 0; i <= len(text)-subLen; i++ {
		if text[i:i+subLen] == substring {
			for j := 0; j < subLen; j++ {
				result[i+j] = true
			}
		}
	}
	return result
}

// AllIndexes returns a set containing every index in [0, length).
// Used when the whole string should be colored.
func AllIndexes(length int) map[int]bool {
	result := make(map[int]bool, length)
	for i := range length {
		result[i] = true
	}
	return result
}




