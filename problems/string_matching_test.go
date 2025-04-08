package problems

import (
	"testing"
)

type testCase struct {
	text                                string
	pattern                             string
	expectedMatches                     int
	expectedComparisonsBruteForce       int
	expectedComparisonsMorrisPratt      int
	expectedComparisonsKnuthMorrisPratt int
	expectedComparisonsKarpRabin        int
	expectedComparisonsBoyerMoore       int
}

func TestStringMatchingAlgorithms(t *testing.T) {
	// Define a list of test cases
	tests := []testCase{
		{
			text:                                "this is a simple example text for testing",
			pattern:                             "example",
			expectedMatches:                     1,
			expectedComparisonsBruteForce:       45,
			expectedComparisonsMorrisPratt:      45,
			expectedComparisonsKnuthMorrisPratt: 45,
			expectedComparisonsKarpRabin:        35,
			expectedComparisonsBoyerMoore:       20,
		},
		{
			text:                                "this is a simple example text for testing",
			pattern:                             "simple",
			expectedMatches:                     1,
			expectedComparisonsBruteForce:       43,
			expectedComparisonsMorrisPratt:      44,
			expectedComparisonsKnuthMorrisPratt: 44,
			expectedComparisonsKarpRabin:        20,
			expectedComparisonsBoyerMoore:       20,
		},
		{
			text:                                "a quick brown fox jumps over the lazy dog",
			pattern:                             "fox",
			expectedMatches:                     1,
			expectedComparisonsBruteForce:       41,
			expectedComparisonsMorrisPratt:      41,
			expectedComparisonsKnuthMorrisPratt: 41,
			expectedComparisonsKarpRabin:        24,
			expectedComparisonsBoyerMoore:       24,
		},
	}

	// Iterate over the test cases and run each algorithm
	for _, tt := range tests {
		t.Run(tt.pattern, func(t *testing.T) {
			// Test BruteForce algorithm
			matches, comparisons := BruteForce(tt.text, tt.pattern)
			if matches != tt.expectedMatches || comparisons != tt.expectedComparisonsBruteForce {
				t.Errorf("BruteForce: Expected %d matches and %d comparisons, got %d matches and %d comparisons", tt.expectedMatches, tt.expectedComparisonsBruteForce, matches, comparisons)
			}

			// Test Morris-Pratt algorithm
			matches, comparisons = MorrisPratt(tt.text, tt.pattern)
			if matches != tt.expectedMatches || comparisons != tt.expectedComparisonsMorrisPratt {
				t.Errorf("MorrisPratt: Expected %d matches and %d comparisons, got %d matches and %d comparisons", tt.expectedMatches, tt.expectedComparisonsMorrisPratt, matches, comparisons)
			}

			// Test Knuth-Morris-Pratt algorithm
			matches, comparisons = KnuthMorrisPratt(tt.text, tt.pattern)
			if matches != tt.expectedMatches || comparisons != tt.expectedComparisonsKnuthMorrisPratt {
				t.Errorf("KnuthMorrisPratt: Expected %d matches and %d comparisons, got %d matches and %d comparisons", tt.expectedMatches, tt.expectedComparisonsKnuthMorrisPratt, matches, comparisons)
			}

			// Test Karp-Rabin algorithm
			// todo: fix
			//matches, comparisons = KarpRabin(tt.text, tt.pattern)
			//if matches != tt.expectedMatches || comparisons != tt.expectedComparisonsKarpRabin {
			//	t.Errorf("KarpRabin: Expected %d matches and %d comparisons, got %d matches and %d comparisons", tt.expectedMatches, tt.expectedComparisonsKarpRabin, matches, comparisons)
			//}

			// Test Boyer-Moore algorithm
			matches, comparisons = BoyerMoore(tt.text, tt.pattern)
			if matches != tt.expectedMatches || comparisons != tt.expectedComparisonsBoyerMoore {
				t.Errorf("BoyerMoore: Expected %d matches and %d comparisons, got %d matches and %d comparisons", tt.expectedMatches, tt.expectedComparisonsBoyerMoore, matches, comparisons)
			}
		})
	}
}
