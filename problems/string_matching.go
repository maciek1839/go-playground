package problems

import (
	"math"
)

/**

 */

// BruteForce Search
// The Brute Force algorithm is a straightforward string matching algorithm that compares the pattern
// against every substring of the text of the same length as the pattern. If a match is found, the
// match counter is incremented. This algorithm is simple but inefficient for larger texts due to its
// O(n*m) time complexity.
func BruteForce(text, pattern string) (int, int) {
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
// The Morris-Pratt algorithm is an optimized version of the Brute Force search that preprocesses the pattern
// by computing an LPS (Longest Prefix Suffix) array. It improves performance by avoiding redundant comparisons.
// This algorithm is more efficient than Brute Force in terms of time complexity, especially for long patterns.
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
				j = lps[j-1]
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
// The Knuth-Morris-Pratt (KMP) algorithm is an efficient string matching algorithm that preprocesses the pattern
// to create an LPS (Longest Prefix Suffix) array and avoids unnecessary comparisons during the search. It has
// a linear time complexity of O(n + m) where n is the length of the text and m is the length of the pattern.
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
				j = lps[j-1]
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

// Karp-Rabin Search with rolling hash
// The Karp-Rabin algorithm uses a rolling hash to efficiently compare substrings in the text with the pattern.
// It avoids performing character-by-character comparison until a potential match is found based on hash equality.
// This algorithm has a time complexity of O(n + m), but it can be affected by hash collisions.
func KarpRabin(text, pattern string) (int, int) {
	m := len(pattern)
	n := len(text)
	matches := 0
	comparisons := 0
	prime := 101 // A prime number for modulus to reduce collisions

	// Compute the hash of the pattern and the first window of the text
	patternHash := hash(pattern, m, prime)
	textHash := hash(text[:m], m, prime)

	// Rolling hash: slide over the text
	for i := 0; i <= n-m; i++ {
		// Compare the hashes first (quick check)
		comparisons++
		if patternHash == textHash && text[i:i+m] == pattern {
			matches++
		}

		// Rehash the text if not at the end
		if i < n-m {
			textHash = rehash(textHash, text[i], text[i+m], m, prime)
		}
	}
	return matches, comparisons
}

// Boyer-Moore Search
// The Boyer-Moore algorithm is a highly efficient string matching algorithm that preprocesses the pattern
// by creating a bad character heuristic. It skips over large portions of the text by comparing from the
// rightmost character of the pattern. This algorithm works efficiently in practice with a time complexity
// of O(n/m) in the best case and O(n*m) in the worst case.
func BoyerMoore(text, pattern string) (int, int) {
	m := len(pattern)
	n := len(text)
	matches := 0
	comparisons := 0

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

// Helper functions
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

func hash(str string, length, prime int) int {
	var h int
	for i := 0; i < length; i++ {
		h = (h*256 + int(str[i])) % prime
	}
	return h
}

func rehash(prevHash int, oldChar, newChar byte, patternLen, prime int) int {
	h := prevHash - int(oldChar)*(int(math.Pow(256, float64(patternLen-1))))%prime
	h = (h*256 + int(newChar)) % prime
	return h
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
