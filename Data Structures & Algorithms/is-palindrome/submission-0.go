func isPalindrome(s string) bool {
	l := 0
	r := len(s)-1

	runes := []rune(strings.ToLower(s))

	for l < r {
		leftLetter := runes[l]
		rightLetter := runes[r]

		if !unicode.IsLetter(leftLetter) && !unicode.IsNumber(leftLetter) {
			l++
			continue
		}

		if !unicode.IsLetter(rightLetter) && !unicode.IsNumber(rightLetter) {
			r--
			continue
		}

		if leftLetter != rightLetter {
			return false
		}

		l++
		r--
	}

	return true
}
