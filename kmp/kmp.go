package kmp

func KMPSearch(s, substr string) int {
	next := buildNext(substr)

	i := 0
	j := 0
	for i < len(s) && j < len(substr) {
		if j < 0 || s[i] == s[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j < len(substr) {
		return -1
	}
	return i - j
}

func buildNext(s string) []int {
	if len(s) == 0 {
		return nil
	}

	next := make([]int, len(s))
	next[0] = -1
	t := -1

	i := 0
	for i < len(s)-1 {
		if 0 > t || s[i] == s[t] {
			i++
			t++
			next[i] = t
		} else {
			t = next[t]
		}
	}

	return next
}
