package main

func BruteForce(text string, pattern string) (int, int) {
	m := len(pattern)
	n := len(text)
	matches := 0
	comparisons := 0

	for i := 0; i <= n-m; i++ {
		match := true
		for j := 0; j < m; j++ {
			comparisons++
			if text[i+j] != pattern[j] {
				match = false
				break
			}
		}
		if match {
			matches++
		}
	}
	return matches, comparisons
}

// Morris-Pratt Search
func MorrisPratt(text, pattern string) (int, int) {
	m := len(pattern)
	n := len(text)
	matches := 0
	comparisons := 0

	lps := computeLPS(pattern)

	i, j := 0, 0
	for i < n {
		comparisons++
		if text[i] == pattern[j] {
			i++
			j++
			if j == m {
				matches++
				j = 0 // Reset instead of using lps[j-1]
			}
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return matches, comparisons
}

// Knuth-Morris-Pratt Search
func KnuthMorrisPratt(text, pattern string) (int, int) {
	m := len(pattern)
	n := len(text)
	matches := 0
	comparisons := 0

	lps := computeLPS(pattern)

	i, j := 0, 0
	for i < n {
		comparisons++
		if text[i] == pattern[j] {
			i++
			j++
			if j == m {
				matches++
				j = lps[j-1] // Use LPS fallback after full match
			}
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return matches, comparisons
}

func computeLPS(pattern string) []int {
	m := len(pattern)
	lps := make([]int, m)
	length := 0
	i := 1

	for i < m {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

// Karp-Rabin Search with rolling hash
func KarpRabin(text, pattern string) (int, int) {
	m := len(pattern)
	n := len(text)
	if m == 0 || n == 0 || m > n {
		return 0, 0
	}

	const d = 256     // Number of characters in the alphabet (ASCII)
	const prime = 101 // A prime number for hashing

	var matches, comparisons int

	h := 1
	for i := 0; i < m-1; i++ {
		h = (h * d) % prime
	}

	patternHash := 0
	textHash := 0

	// Calculate initial hash values
	for i := 0; i < m; i++ {
		patternHash = (d*patternHash + int(pattern[i])) % prime
		textHash = (d*textHash + int(text[i])) % prime
	}

	for i := 0; i <= n-m; i++ {
		comparisons++
		if patternHash == textHash {
			// Check character-by-character only if hashes match
			if text[i:i+m] == pattern {
				matches++
			}
		}
		// Calculate hash for next window
		if i < n-m {
			textHash = (d*(textHash-int(text[i])*h) + int(text[i+m])) % prime
			if textHash < 0 {
				textHash += prime
			}
		}
	}

	return matches, comparisons
}

// Boyer-Moore Search (only the bad character heurstic included)
func BoyerMoore(text, pattern string) (int, int) {
	m := len(pattern)
	n := len(text)
	matches := 0
	comparisons := 0

	// In the Bad Character Heuristic, we want to align the last occurrence of the mismatched character
	// in the pattern with where it appears in the text.
	badChar := make(map[byte]int)
	for i := 0; i < m; i++ {
		badChar[pattern[i]] = i
	}

	i := 0
	for i <= n-m {
		j := m - 1
		comparisons++
		for j >= 0 && text[i+j] == pattern[j] {
			j--
			comparisons++
		}
		if j < 0 {
			matches++
			i += m
		} else {
			i += max(1, j-badChar[text[i+j]])
		}
	}
	return matches, comparisons
}

// Boyer-Moore Search (both heuristics included).
func BoyerMooreOptimized(text, pattern string) (int, int) {
	m := len(pattern)
	n := len(text)
	if m == 0 || n == 0 || m > n {
		return 0, 0
	}

	// Preprocessing
	badChar := buildBadCharTable(pattern)
	goodSuffix := buildGoodSuffixTable(pattern)

	matches := 0
	comparisons := 0
	s := 0

	for s <= n-m {
		j := m - 1

		for j >= 0 && pattern[j] == text[s+j] {
			comparisons++
			j--
		}

		if j < 0 {
			matches++
			s += goodSuffix[0]
		} else {
			comparisons++
			badIdx, ok := badChar[text[s+j]]
			if !ok {
				badIdx = -1
			}
			badCharShift := j - badIdx
			goodSuffixShift := goodSuffix[j+1]
			s += max(1, max(badCharShift, goodSuffixShift))
		}
	}

	return matches, comparisons
}

func buildBadCharTable(pattern string) map[byte]int {
	table := make(map[byte]int)
	for i := 0; i < len(pattern); i++ {
		table[pattern[i]] = i
	}
	return table
}

func buildGoodSuffixTable(pattern string) []int {
	m := len(pattern)
	shift := make([]int, m+1)
	border := make([]int, m+1)

	i := m
	j := m + 1
	border[i] = j

	// First pass: preprocess border positions
	for i > 0 {
		for j <= m && pattern[i-1] != pattern[j-1] {
			if shift[j] == 0 {
				shift[j] = j - i
			}
			j = border[j]
		}
		i--
		j--
		border[i] = j
	}

	// Second pass: fill the shift table
	j = border[0]
	for i := 0; i <= m; i++ {
		if shift[i] == 0 {
			shift[i] = j
		}
		if i == j {
			j = border[j]
		}
	}

	return shift
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
