package strings

// check if a string is a palindrome and A != B
func ContainsABBA(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if IsPalindrome(s[i : i+4]) && s[i] != s[i+1] {
			return true
		}
	}

	return false
}

func IsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-(1+i)] {
			return false
		}
	}
	return true
}