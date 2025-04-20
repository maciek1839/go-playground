package main

import (
	"testing"
)

type testCase struct {
	text                                   string
	pattern                                string
	expectedMatches                        int
	expectedComparisonsBruteForce          int
	expectedComparisonsMorrisPratt         int
	expectedComparisonsKnuthMorrisPratt    int
	expectedComparisonsKarpRabin           int
	expectedComparisonsBoyerMoore          int
	expectedComparisonsOptimizedBoyerMoore int
}

func TestStringMatchingAlgorithms(t *testing.T) {
	// Define a list of test cases
	tests := []testCase{
		{
			text:                                   "THIS IS A SIMPLE EXAMPLE",
			pattern:                                "SIMPLE",
			expectedMatches:                        1,
			expectedComparisonsBruteForce:          26, // starts matching mid-way
			expectedComparisonsMorrisPratt:         26,
			expectedComparisonsKnuthMorrisPratt:    26,
			expectedComparisonsKarpRabin:           19, // early hash mismatch
			expectedComparisonsBoyerMoore:          16, // jumps efficiently
			expectedComparisonsOptimizedBoyerMoore: 14, // jumps efficiently
		},
		{
			text:                                   "AAAAAAAAAAH",
			pattern:                                "AAAAH",
			expectedMatches:                        1,
			expectedComparisonsBruteForce:          35, // bad for naive due to repeated A's
			expectedComparisonsMorrisPratt:         17, // good use of LPS
			expectedComparisonsKnuthMorrisPratt:    17,
			expectedComparisonsKarpRabin:           7,
			expectedComparisonsBoyerMoore:          12, // skips most
			expectedComparisonsOptimizedBoyerMoore: 11, // skips most
		},
		{
			text:                                   "THIS IS MY NEW STRING AAAAHHHH",
			pattern:                                "NEW",
			expectedMatches:                        1,
			expectedComparisonsBruteForce:          31,
			expectedComparisonsMorrisPratt:         31,
			expectedComparisonsKnuthMorrisPratt:    31,
			expectedComparisonsKarpRabin:           28,
			expectedComparisonsBoyerMoore:          17,
			expectedComparisonsOptimizedBoyerMoore: 12,
		},
		{
			text:                                   "THIS TEST WILL HAVE MULTIPLE MATCHES, SO THAT WE CAN TEST. ONE MORE TEST.",
			pattern:                                "TEST",
			expectedMatches:                        3,
			expectedComparisonsBruteForce:          86, // needs to rescan
			expectedComparisonsMorrisPratt:         78, // good reuse of LPS
			expectedComparisonsKnuthMorrisPratt:    81,
			expectedComparisonsKarpRabin:           70,
			expectedComparisonsBoyerMoore:          42,
			expectedComparisonsOptimizedBoyerMoore: 33,
		},
		{
			text:                                   "COMPUTER SCIENCE IS NO MORE ABOUT COMPUTERS THAN ASTRONOMY IS ABOUT TELESCOPES",
			pattern:                                "NO",
			expectedMatches:                        2,
			expectedComparisonsBruteForce:          81,
			expectedComparisonsMorrisPratt:         80,
			expectedComparisonsKnuthMorrisPratt:    80,
			expectedComparisonsKarpRabin:           77,
			expectedComparisonsBoyerMoore:          86,
			expectedComparisonsOptimizedBoyerMoore: 45,
		},
		{
			text:                                   "NO COMPUTER IS EVER GOING TO ASK A NEW, REASONABLE QUESTION. IT TAKES TRAINED PEOPLE TO DO THAT.",
			pattern:                                "TRAINED",
			expectedMatches:                        1,
			expectedComparisonsBruteForce:          102,
			expectedComparisonsMorrisPratt:         104,
			expectedComparisonsKnuthMorrisPratt:    104,
			expectedComparisonsKarpRabin:           90,
			expectedComparisonsBoyerMoore:          27,
			expectedComparisonsOptimizedBoyerMoore: 22,
		},
		{
			text:                                   "WE CAN ONLY SEE A SHORT DISTANCE AHEAD, BUT WE CAN SEE PLENTY THERE THAT NEEDS TO BE DONE.",
			pattern:                                "DISTANCE",
			expectedMatches:                        1,
			expectedComparisonsBruteForce:          92,
			expectedComparisonsMorrisPratt:         93,
			expectedComparisonsKnuthMorrisPratt:    93,
			expectedComparisonsKarpRabin:           83,
			expectedComparisonsBoyerMoore:          27,
			expectedComparisonsOptimizedBoyerMoore: 22,
		},
		{
			text:                                   "QAZQAZQAZQAZZQAZZQAZZQAZZZZZZZZQQQQQQQZZZZZQAQAQAQAQAZZZQAAAAAZZZQAZZZZZZZQQQQQAAAZZZ",
			pattern:                                "QAZZZ",
			expectedMatches:                        3,
			expectedComparisonsBruteForce:          137, // lots of backtracking
			expectedComparisonsMorrisPratt:         108, // strong LPS savings
			expectedComparisonsKnuthMorrisPratt:    108,
			expectedComparisonsKarpRabin:           81,
			expectedComparisonsBoyerMoore:          63,
			expectedComparisonsOptimizedBoyerMoore: 55,
		},
		{
			text:                                   "COMPUTER SCIENCE IS THE OPERATING SYSTEM FOR ALL INNOVATION.",
			pattern:                                "NOVA",
			expectedMatches:                        1,
			expectedComparisonsBruteForce:          63,
			expectedComparisonsMorrisPratt:         64,
			expectedComparisonsKnuthMorrisPratt:    64,
			expectedComparisonsKarpRabin:           57,
			expectedComparisonsBoyerMoore:          24,
			expectedComparisonsOptimizedBoyerMoore: 19,
		},
		{
			text:                                   "WWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWKWWWWWKWWWWK",
			pattern:                                "WWWK",
			expectedMatches:                        3,
			expectedComparisonsBruteForce:          216, // painful re-checking
			expectedComparisonsMorrisPratt:         108, // big help from LPS = [0, 1, 2, 0]
			expectedComparisonsKnuthMorrisPratt:    108,
			expectedComparisonsKarpRabin:           57, // some hash collisions
			expectedComparisonsBoyerMoore:          63,
			expectedComparisonsOptimizedBoyerMoore: 60,
		},
		{
			text:                                   "aaaaaaabaaaaaaabaaaaaaabaaaaaaabaaaaaaabmatchhereaaaaaaab",
			pattern:                                "aaaaaaabmatch",
			expectedMatches:                        1,
			expectedComparisonsBruteForce:          195, // nearly full re-check on mismatch
			expectedComparisonsMorrisPratt:         61,  // LPS kicks in
			expectedComparisonsKnuthMorrisPratt:    61,  // same as Morris-Pratt
			expectedComparisonsKarpRabin:           45,  // lots of hash collisions
			expectedComparisonsBoyerMoore:          22,  // skips huge chunks
			expectedComparisonsOptimizedBoyerMoore: 21,
		},
		{
			text:                                   "ababababababababababababababababx",
			pattern:                                "abababababababx",
			expectedMatches:                        1,
			expectedComparisonsBruteForce:          159, // lots of retries
			expectedComparisonsMorrisPratt:         42,  // good fallback
			expectedComparisonsKnuthMorrisPratt:    42,
			expectedComparisonsKarpRabin:           19, // some hash checks needed
			expectedComparisonsBoyerMoore:          25, // some clever skipping
			expectedComparisonsOptimizedBoyerMoore: 24,
		},
		{
			text:                                   "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			pattern:                                "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			expectedMatches:                        1,
			expectedComparisonsBruteForce:          38, // perfect match at the end
			expectedComparisonsMorrisPratt:         38,
			expectedComparisonsKnuthMorrisPratt:    38,
			expectedComparisonsKarpRabin:           1,  // one hash match
			expectedComparisonsBoyerMoore:          39, // one big jump to match
			expectedComparisonsOptimizedBoyerMoore: 38,
		},
		{
			text:                                   "lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod",
			pattern:                                "eiusmod",
			expectedMatches:                        1,
			expectedComparisonsBruteForce:          75,
			expectedComparisonsMorrisPratt:         75,
			expectedComparisonsKnuthMorrisPratt:    75,
			expectedComparisonsKarpRabin:           63,
			expectedComparisonsBoyerMoore:          21,
			expectedComparisonsOptimizedBoyerMoore: 19,
		},
		// {
		// 	text:                                   "ababcabcababcabcababcabcabcabcabcabcabcabcabcabcabcababcabc",
		// 	pattern:                                "abcabcabcabcabc",
		// 	expectedMatches:                        2,
		// 	expectedComparisonsBruteForce:          120, // slow due to overlap
		// 	expectedComparisonsMorrisPratt:         55,  // good use of LPS
		// 	expectedComparisonsKnuthMorrisPratt:    55,
		// 	expectedComparisonsKarpRabin:           35, // OK performance
		// 	expectedComparisonsBoyerMoore:          20, // solid skipping
		// 	expectedComparisonsOptimizedBoyerMoore: 22,
		// },
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
			matches, comparisons = KarpRabin(tt.text, tt.pattern)
			if matches != tt.expectedMatches || comparisons != tt.expectedComparisonsKarpRabin {
				t.Errorf("KarpRabin: Expected %d matches and %d comparisons, got %d matches and %d comparisons", tt.expectedMatches, tt.expectedComparisonsKarpRabin, matches, comparisons)
			}

			// Test Boyer-Moore algorithm
			matches, comparisons = BoyerMoore(tt.text, tt.pattern)
			if matches != tt.expectedMatches || comparisons != tt.expectedComparisonsBoyerMoore {
				t.Errorf("BoyerMoore: Expected %d matches and %d comparisons, got %d matches and %d comparisons", tt.expectedMatches, tt.expectedComparisonsBoyerMoore, matches, comparisons)
			}

			// Test optimized Boyer-Moore algorithm
			matches, comparisons = BoyerMooreOptimized(tt.text, tt.pattern)
			if matches != tt.expectedMatches || comparisons != tt.expectedComparisonsOptimizedBoyerMoore {
				t.Errorf("BoyerMooreOptimized: Expected %d matches and %d comparisons, got %d matches and %d comparisons", tt.expectedMatches, tt.expectedComparisonsBoyerMoore, matches, comparisons)
			}
		})
	}
}
