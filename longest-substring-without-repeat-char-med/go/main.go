package main

func main() {
	word := "abcabcbb"
	result := lengthOfLongestSubstring(word)
	println(result) // Output: 3
}

func lengthOfLongestSubstring(s string) int {
	seen := make(map[rune]int)
	left := 0
	maxLength := 0

	for right, ch := range []rune(s) {
		lastIndex, found := seen[ch]
		if found && lastIndex >= left {
			left = lastIndex + 1
		}

		seen[ch] = right
		windowLength := right - left + 1
		if windowLength > maxLength {
			maxLength = windowLength
		}
	}

	return maxLength

}
