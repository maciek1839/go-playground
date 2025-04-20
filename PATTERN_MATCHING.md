# Pattern matching algorithms

Pattern matching algorithms are fundamental in computer science, used to find occurrences of a pattern (usually a substring or sequence) within a text. They're widely used in areas like:
- Text searching (e.g., finding a word in a document)
- DNA sequence analysis
- Syntax highlighting
- Spam detection
- Search engines
- Compilers (for lexical analysis)

Pattern matching is the act of checking whether a particular pattern (a sequence of characters, symbols, or tokens) exists within a given larger structure (typically a string or array), and if so, identifying where.

1. Brute Force Algorithm
   - Idea: Check for the pattern at every possible position in the text.
   - Time complexity: O(n * m), where n is the text length, and m is the pattern length.
   - Use: Simple to implement, not efficient for large texts.
2. Knuth-Morris-Pratt (KMP) Algorithm
   - Idea: Use information about the pattern itself to avoid unnecessary comparisons.
   - Time complexity: O(n + m)
   - Key concept: Preprocesses the pattern into a "failure function" (or prefix table).
3. Boyer-Moore Algorithm
   - Idea: Start matching from the end of the pattern and use two heuristics (bad character and good suffix) to skip sections of text.
   - Time complexity: Best case: sub-linear; Worst case: O(n * m)
   - Use: One of the fastest in practice for text searching.
4. Rabin-Karp Algorithm
   - Idea: Use hashing to find any one of a set of pattern strings in a text.
   - Time complexity: Average O(n + m), Worst O(n * m)
   - Key concept: Rolling hash function.
5. Aho-Corasick Algorithm
   - Idea: Search for multiple patterns simultaneously using a trie and failure links.
   - Time complexity: O(n + m + z), where z is the number of matches.
   - Use: Good for virus scanning, text filtering.
6. Suffix Tree / Suffix Array
   - Idea: Build a data structure representing all suffixes of the text or pattern.
   - Time complexity: O(n) for construction, fast for querying.
   - Use: Complex but very efficient for repeated queries.


| Use Case                                  | Suggested Algorithm | Time Complexity                   |
|-------------------------------------------|---------------------|-----------------------------------|
| Simple one-time search                    | Brute-Force or KMP  | Brute-Force: O(n × m), KMP: O(n + m)    |
| Fast large-scale search                   | Boyer-Moore         | Best: sub-linear, Worst: O(n × m) |
| Searching multiple patterns               | Aho-Corasick        | O(n + m + z)                      |
| Pattern search with wildcards             | Suffix Tree / Trie  | O(n) preprocessing, O(m) query    |
| Search for multiple patterns with hashing | Rabin-Karp          | Avg: O(n + m), Worst: O(n × m)    |

- n = length of the text
- m = length of the pattern
- z = total number of pattern matches found

## Brute-Force / Naive Algorithm

The Brute Force algorithm is a straightforward string matching algorithm that compares the pattern 
against every substring of the text of the same length as the pattern.
It works by checking for the pattern at every possible position in the text. 
For each position, it compares the substring with the pattern character by character. 

If a match is found, the match counter is incremented. This algorithm is simple but inefficient for larger texts due to its O(n*m) time complexity.

```text
function naiveSearch(text, pattern):
    n = length(text)
    m = length(pattern)
    
    for i from 0 to n - m:
        match = true
        for j from 0 to m - 1:
            if text[i + j] != pattern[j]:
                match = false
                break
        if match:
            print "Pattern found at index", i
```

## Knuth-Morris-Pratt (KMP)

KMP improves over the brute-force approach by avoiding unnecessary re-checks. It pre-processes the pattern to create a longest prefix-suffix (LPS) table, which is used to shift the pattern efficiently when a mismatch occurs.


The **Morris-Pratt** algorithm is the foundational version of pattern matching that first introduced the idea of preprocessing the pattern using a _Longest Prefix Suffix (LPS)_ table. It allows the search algorithm to skip redundant comparisons by leveraging repeated prefix/suffix structures in the pattern. The **Knuth-Morris-Pratt (KMP)** algorithm builds upon this idea with a subtle but powerful optimization: during LPS construction, if a mismatch occurs, KMP attempts to reuse previously computed LPS values to avoid unnecessary resets. In practice, this means that KMP may "retry" earlier parts of the pattern before giving up on a match, while Morris-Pratt always advances forward. Despite this difference, their core loop logic and performance are nearly identical in modern implementations. Most real-world "KMP" implementations actually resemble Morris-Pratt with this optimization folded into the LPS computation step.


```text
function computeLPS(pattern):
    lps = array of zeros with size length(pattern)
    length = 0
    i = 1
    while i < length(pattern):
        if pattern[i] == pattern[length]:
            length += 1
            lps[i] = length
            i += 1
        else if length != 0:
            length = lps[length - 1]
        else:
            lps[i] = 0
            i += 1
    return lps

function KMPSearch(text, pattern):
    lps = computeLPS(pattern)
    i = 0  // index for text
    j = 0  // index for pattern
    while i < length(text):
        if pattern[j] == text[i]:
            i += 1
            j += 1
        if j == length(pattern):
            print "Pattern found at index", i - j
            j = lps[j - 1]
        else if i < length(text) and pattern[j] != text[i]:
            if j != 0:
                j = lps[j - 1]
            else:
                i += 1
```

### Examples

- 🧵 Text: "a b b c e f d e f g"
- 🧩 Pattern: "e f g"
- 📊 LPS: [0, 0, 0]


| Step | `i` (Text) | `j` (Pattern) | `text[i]` | `pattern[j]` | Match? | Action                                      |
|------|------------|---------------|-----------|----------------|--------|---------------------------------------------|
| 1    | 0          | 0             | `a`       | `e`            | ❌     | `j == 0` → `i++`                             |
| 2    | 1          | 0             | `b`       | `e`            | ❌     | `j == 0` → `i++`                             |
| 3    | 2          | 0             | `b`       | `e`            | ❌     | `j == 0` → `i++`                             |
| 4    | 3          | 0             | `c`       | `e`            | ❌     | `j == 0` → `i++`                             |
| 5    | 4          | 0             | `e`       | `e`            | ✅     | `i++`, `j++ → j = 1`                         |
| 6    | 5          | 1             | `f`       | `f`            | ✅     | `i++`, `j++ → j = 2`                         |
| 7    | 6          | 2             | `d`       | `g`            | ❌     | `j != 0` → `j = lps[1] = 0` (use LPS!)      |
| 8    | 6          | 0             | `d`       | `e`            | ❌     | `j == 0` → `i++`                             |
| 9    | 7          | 0             | `e`       | `e`            | ✅     | `i++`, `j++ → j = 1`                         |
| 10   | 8          | 1             | `f`       | `f`            | ✅     | `i++`, `j++ → j = 2`                         |
| 11   | 9          | 2             | `g`       | `g`            | ✅     | `i++`, `j++ → j = 3 == len(pattern)` → MATCH |
| ✅   | —          | —             | —         | —              | ✅     | Pattern found at `i - j = 9 - 3 = 6`         |

- 🧵 Text: "a b a b a b c a b a b d"
- 🧩 Pattern: "a b a b d"
- 📊 LPS: [0, 0, 1, 2, 0]

Step 1: Compute LPS Table

| i (index) | pattern[i] | length | LPS[i] | Explanation                         |
|-----------|------------|--------|--------|-------------------------------------|
| 0         | a          | —      | 0      | Always 0                            |
| 1         | b          | 0      | 0      | No match with pattern[0]            |
| 2         | a          | 0 → 1  | 1      | `a == a`, increment length          |
| 3         | b          | 1 → 2  | 2      | `b == b`, increment length          |
| 4         | d          | ✖️     | 0      | `d != a`, fall back to LPS[1] = 0   |

✅ Final **LPS table**: `[0, 0, 1, 2, 0]`

---

| Step | `i` (Text) | `j` (Pattern) | `text[i]` | `pattern[j]` | Match? | Action / Explanation                          |
|------|------------|---------------|-----------|----------------|--------|-----------------------------------------------|
| 1    | 0          | 0             | `a`       | `a`            | ✅     | `i++, j++`                                     |
| 2    | 1          | 1             | `b`       | `b`            | ✅     | `i++, j++`                                     |
| 3    | 2          | 2             | `a`       | `a`            | ✅     | `i++, j++`                                     |
| 4    | 3          | 3             | `b`       | `b`            | ✅     | `i++, j++`                                     |
| 5    | 4          | 4             | `a`       | `d`            | ❌     | Mismatch → `j = lps[3] = 2`                    |
| 6    | 4          | 2             | `a`       | `a`            | ✅     | `i++, j++`                                     |
| 7    | 5          | 3             | `b`       | `b`            | ✅     | `i++, j++`                                     |
| 8    | 6          | 4             | `c`       | `d`            | ❌     | Mismatch → `j = lps[3] = 2`                    |
| 9    | 6          | 2             | `c`       | `a`            | ❌     | Mismatch → `j = lps[1] = 0`                    |
| 10   | 6          | 0             | `c`       | `a`            | ❌     | `i++`                                          |
| 11   | 7          | 0             | `a`       | `a`            | ✅     | `i++, j++`                                     |
| 12   | 8          | 1             | `b`       | `b`            | ✅     | `i++, j++`                                     |
| 13   | 9          | 2             | `a`       | `a`            | ✅     | `i++, j++`                                     |
| 14   | 10         | 3             | `b`       | `b`            | ✅     | `i++, j++`                                     |
| 15   | 11         | 4             | `d`       | `d`            | ✅     | 🎉 Pattern match at index `i - j = 6`          |

- 🧵 Text: "a b a b a b a b a c"
- 🧩 Pattern: "a b a b a c"
- 📊 LPS: [0, 0, 1, 2, 3, 0]

Step 1: Compute LPS Table

| i | pattern[i] | Compare with `pattern[length]` | length before | Match? | length after | LPS[i] |
|---|------------|----------------------------------|---------------|--------|---------------|--------|
| 0 | a          | —                                | —             | —      | —             | 0      |
| 1 | b          | a                                | 0             | ❌     | 0             | 0      |
| 2 | a          | a                                | 0             | ✅     | 1             | 1      |
| 3 | b          | b                                | 1             | ✅     | 2             | 2      |
| 4 | a          | a                                | 2             | ✅     | 3             | 3      |
| 5 | c          | b                                | 3             | ❌     | fallback to 1 (then 0) | 0 |

**✅ Final LPS Table**: `[0, 0, 1, 2, 3, 0]`

---

| Step | Text Index `i` | Pattern Index `j` | `text[i]` | `pattern[j]` | Match? | Action / Explanation                          |
|------|----------------|-------------------|-----------|--------------|--------|-----------------------------------------------|
| 1    | 0              | 0                 | a         | a            | ✅     | `i++, j++`                                     |
| 2    | 1              | 1                 | b         | b            | ✅     | `i++, j++`                                     |
| 3    | 2              | 2                 | a         | a            | ✅     | `i++, j++`                                     |
| 4    | 3              | 3                 | b         | b            | ✅     | `i++, j++`                                     |
| 5    | 4              | 4                 | a         | a            | ✅     | `i++, j++`                                     |
| 6    | 5              | 5                 | b         | c            | ❌     | Fallback using `LPS[4] = 3 → j = 3`            |
| 7    | 5              | 3                 | b         | b            | ✅     | `i++, j++`                                     |
| 8    | 6              | 4                 | a         | a            | ✅     | `i++, j++`                                     |
| 9    | 7              | 5                 | b         | c            | ❌     | Fallback `j = LPS[4] = 3`                      |
| 10   | 7              | 3                 | b         | b            | ✅     | `i++, j++`                                     |
| 11   | 8              | 4                 | a         | a            | ✅     | `i++, j++`                                     |
| 12   | 9              | 5                 | c         | c            | ✅     | 🎉 Full match at index `i - j = 9 - 6 = 3`     |

## Rabin-Karp

The Karp-Rabin algorithm uses a rolling hash to efficiently compare substrings in the text with the pattern.
It avoids performing character-by-character comparison until a potential match is found based on hash equality.
This algorithm has a time complexity of O(n + m), but it can be affected by hash collisions.

Instead of comparing the full pattern with each substring of the text, it:
1. Calculates a hash of the pattern.
2. Slides a window over the text, computing a rolling hash of each substring.
3. Only compares characters when the hashes match (likely match).

```text
function RabinKarp(text, pattern, d, q):
    n = length(text)
    m = length(pattern)
    h = pow(d, m-1) % q
    p = 0  // hash for pattern
    t = 0  // hash for text
    for i from 0 to m - 1:
        p = (d * p + ASCII(pattern[i])) % q
        t = (d * t + ASCII(text[i])) % q
    for s from 0 to n - m:
        if p == t:
            if text[s:s + m] == pattern:
                print "Pattern found at index", s
        if s < n - m:
            t = (d * (t - ASCII(text[s]) * h) + ASCII(text[s + m])) % q
            if t < 0:
                t = t + q
```

d: Base (usually 256 for extended ASCII)
q: A large prime number to reduce collisions and avoid overflows

---

Time Complexity:
- Average	O(n + m)
    - Fast because it uses hashing to compare substrings. Only does full character comparison when hashes match.
- Worst	O(n × m)
    - Happens when there are many hash collisions, so it ends up comparing each character manually for many shifts.

Text: "aaaaa"  
Pattern: "aaa"
- All substrings like "aaa" will have the same hash.
- If the hashing isn't perfect, it could falsely match each one and do full O(m) comparison each time.
- That's where the worst-case O(n × m) comes from.

## Boyer-Moore

The Boyer-Moore algorithm is one of the fastest string-search algorithms in practice. It was developed by Robert S. Boyer and J Strother Moore in 1977. 
It starts matching from the end of the pattern and uses two heuristics (bad character and good suffix) to skip large parts of the text. 
It's very efficient in practice for large texts with a time complexity of O(n/m) in the best case and O(n*m) in the worst case.

---

- Bad Character Rule
    - When a mismatch occurs, shift the pattern so that the bad character in the text aligns with the last occurrence of that character in the pattern.
    - If the character doesn't exist in the pattern, the pattern can be shifted past it entirely.

```text
Text:     A B C F A B C D A B C F
Pattern:  A B C F A B C F
Mismatch at position 7 ('D' vs 'F')

'D' is not found in the pattern, so the Bad Character Heuristic tells us to shift the pattern entirely to the right, skipping the character 'D'.

The pattern shifts 8 positions (since 'D' is not in the pattern).

Text:     A B C F A B C D A B C F
                          ↑  
Pattern:                  A B C F A B C F
```

- Good Suffix Rule
    - If a suffix of the pattern matches but the preceding character mismatches, shift the pattern to the next occurrence of this suffix or a prefix that matches.

```text
Text:     A B C F A B C D A B C F
Pattern:  A B C F A B C F
Mismatch at position 7 ('D' vs 'F')

The pattern shifts by 4 positions, aligning the known suffix 'A B C F' with the same occurrence in the text.

The pattern shifts 4 positions to align the matching suffix 'A B C F'.

Text:     A B C F A B C D A B C F
                  ↑  
Pattern:          A B C F A B C F
```

A heuristic is a practical method or shortcut used to solve a problem faster when classic methods are too slow or complex. It doesn't always guarantee the perfect or optimal solution, but it's good enough, fast, and efficient in most cases. Heuristics examples:
- Bad Character Rule
- Good Suffix Rule

 Heuristic explained simply:
- Think of it like "rule of thumb".
- Instead of checking everything, you use experience, patterns, or clever shortcuts to make a pretty good guess.

| Aspect            | Heuristic                                              | Algorithm                                              |
|-------------------|--------------------------------------------------------|--------------------------------------------------------|
| Definition         | A practical approach or rule of thumb to solve a problem faster or more easily | A step-by-step, well-defined procedure to solve a problem |
| Goal              | Find a good enough solution quickly                    | Find the correct/optimal solution                      |
| Basis             | Intuition, experience, or educated guess               | Logic and mathematics                                  |
| Performance       | Typically faster, not always accurate                  | Deterministic and reliable in accuracy                 |
| Consistency       | May produce different results on different runs        | Always produces the same output for the same input     |
| Examples          | Boyer-Moore heuristics, A* search estimations          | KMP, Dijkstra’s algorithm, Binary Search               |
| Guarantees        | No guarantee of optimality or even success             | Guarantees a correct result if implemented correctly   |
| Use Cases         | Complex problems with large search spaces              | Problems where precision is critical                   |


```text
function badCharHeuristic(pattern):
    badChar = array of size 256 initialized to -1
    for i from 0 to length(pattern) - 1:
        //  the ASCII code of the character at index i in the pattern.
        badChar[ASCII(pattern[i])] = i
    return badChar

function goodSuffixHeuristic(pattern):
    m = length(pattern)
    shift = array of size m + 1 initialized to 0
    border = array of size m + 1 initialized to 0

    i = m
    j = m + 1
    border[i] = j

    while i > 0:
        while j <= m and pattern[i - 1] ≠ pattern[j - 1]:
            if shift[j] == 0:
                shift[j] = j - i
            j = border[j]
        i = i - 1
        j = j - 1
        border[i] = j

    j = border[0]
    for i from 0 to m:
        if shift[i] == 0:
            shift[i] = j
        if i == j:
            j = border[j]

    return shift

function BoyerMooreSearch(text, pattern):
    m = length(pattern)
    n = length(text)
    if m == 0 or n == 0 or m > n:
        return

    badChar = badCharHeuristic(pattern)
    goodSuffix = goodSuffixHeuristic(pattern)
    s = 0

    while s <= n - m:
        j = m - 1
        while j ≥ 0 and pattern[j] == text[s + j]:
            j = j - 1

        if j < 0:
            print "Pattern found at index", s
            s = s + goodSuffix[0]
        else:
            badCharShift = j - badChar[ASCII(text[s + j])]
            goodSuffixShift = goodSuffix[j + 1]
            s = s + max(1, max(badCharShift, goodSuffixShift))

```


✅ Best Case: O(n/m) | Bad Character Rule
- This happens when the characters in the text don’t occur in the pattern, or mismatches happen near the end of the pattern.
- Boyer-Moore skips most of the text, potentially shifting the pattern m positions at a time.
- o, in the best case, we only do n/m comparisons.
- Example: Searching for "ABCD" in "ZZZZZZZZZZZZZZZZZZ"
    - Each 'Z' mismatch causes a big jump — super fast!
    - No match found - only 4 comparisons were made: at i = 0, 4, 8, 12

🔄 Step-by-Step Matching

Step 1: Align pattern at i = 0

```text
Text:    Z Z Z Z Z Z Z Z Z Z Z Z Z Z Z Z Z Z
Pattern: A B C D
          ↑
```

- Compare from right: D vs Z → ❌ mismatch
- 'Z' not in pattern → shift = 3 - (-1) = 4
- → Shift pattern right by 4 → i = 4

Step 2: i = 4

```text
Text:        Z Z Z Z Z Z Z Z Z Z Z Z Z Z Z Z
Pattern:         A B C D
                  ↑
```
- D vs Z → ❌ mismatch
- Shift = 4 → i = 8

Step 3: i = 8

```text
Text:            Z Z Z Z Z Z Z Z Z Z Z Z Z Z
Pattern:             A B C D
                      ↑
```

- D vs Z → ❌ mismatch
- Shift = 4 → i = 12

Step 4: i = 12

```text
Text:                Z Z Z Z Z Z Z Z Z Z Z Z
Pattern:                 A B C D
                          ↑
```
    
- D vs Z → ❌ mismatch
- Shift = 4 → i = 16

Step 5: i = 16

```text
Text:                    Z Z Z Z Z Z Z Z Z Z
Pattern:                     A B C D
                              ↑
```

- Not enough text left to match → Done.

❌ Worst Case: O(n·m) | Bad Character Rule
- Happens when the text and pattern have many overlapping characters and the shifts are small.
- Each character of the text is compared with almost all of the pattern (like in degenerate cases).
- Example of worst case: Text like "aaaaaaaaaaab" and pattern like "aaaab"
    - The text and pattern are very repetitive, and mismatches happen late in the pattern...
    - The shifts become small, and the pattern is moved only one character forward each time, resulting in lots of redundant comparisons.

🔄 Step-by-Step Matching

Step 1: Align pattern at i = 0

```text
Text:     a a a a a a a a a a a b
Pattern:  a a a a b
          ↑ ↑ ↑ ↑ ↑
          0 1 2 3 4
```

Compare from the right:
- b vs a → ❌ mismatch
- But the character a does exist in the pattern → bad character heuristic gives small shift
    - Mismatch char is 'a' in text.
    - Mismatched index in pattern = 4
    - Last occurrence of 'a' in the pattern is index 3
    - shift = j - badChar[text[s + j]] = 4 - 3 = 1
- Good suffix doesn’t help much either — we only shift 1 position
- We compare right to left:


Step 2: Align at i = 1

```text
Text:       a a a a a a a a a a b
Pattern:    a a a a b
             ↑
```

- Same thing: compare all the way to b → mismatch again
- → Shift by 1 again
- This happens repeatedly...

Until we get to:  
Final Step: Align at i = 7

```text
Text:             a a a a b
Pattern:          a a a a b
                   ↑
```

Final Result:
- m = 5, n = 12
- Nearly (n - m) * m = (7 * 5) = 35 comparisons
- That's O(n·m) in the worst case

---

Optimized Example (Good Suffix Rule)

```text
Text: A B C F A B C D A B C F
Pattern: A B C F A B C F
```


🔢 Step 1: Align at index 0

```text
Text:     A B C F A B C D A B C F
Pattern:  A B C F A B C F
                        ↑
```

- Mismatch at position 7 of pattern (F) vs text[7] = D → ❌
- Matched suffix: "A B C F" (last 4 chars matched before the mismatch)
- 🔁 Bad Character Heuristic would only shift by 1
- 🔥 But Good Suffix says:
    - "Hey, I’ve seen this same suffix A B C F earlier in the pattern — at index 0. So let’s shift the pattern so that the previous occurrence aligns with it."

🎯 Shift the pattern by 4 instead of 1!

```text
Text:         A B C F A B C D A B C F
Pattern:              A B C F A B C F
                              ↑
```

Now we can immediately recheck the match without redundant comparisons.

## Aho-Corasick

The Aho-Corasick algorithm constructs a trie of all the patterns and uses failure links (like in KMP) to perform multi-pattern matching in linear time. It’s especially effective for dictionary matching problems.

```text
function buildTrie(patterns):
    root = new TrieNode()
    for pattern in patterns:
        insert pattern into trie rooted at root

function buildFailureLinks(root):
    queue = empty queue
    root.fail = root
    for child in root.children:
        child.fail = root
        queue.enqueue(child)
    while not queue.empty():
        current = queue.dequeue()
        for char, child in current.children:
            fail = current.fail
            while fail != root and char not in fail.children:
                fail = fail.fail
            child.fail = fail.children.get(char, root)
            queue.enqueue(child)

function searchAC(text, root):
    node = root
    for i in range(length(text)):
        while node != root and text[i] not in node.children:
            node = node.fail
        if text[i] in node.children:
            node = node.children[text[i]]
        if node.output:
            for pattern in node.output:
                print "Pattern", pattern, "found at", i - len(pattern) + 1
```

## Suffix Tree / Suffix Array

Suffix Trees and Arrays represent all possible suffixes of a text and enable very fast substring searches. 
Suffix Trees are more space-heavy but allow more advanced operations, while Suffix Arrays are compact and efficient with binary search.

```text
function buildSuffixArray(text):
    suffixes = [(text[i:], i) for i in range(len(text))]
    suffixes.sort()  // lexicographically
    return [index for (suffix, index) in suffixes]

function searchPattern(text, pattern, suffixArray):
    l = 0
    r = length(suffixArray) - 1
    while l <= r:
        mid = (l + r) // 2
        start = suffixArray[mid]
        substr = text[start:start + length(pattern)]
        if substr == pattern:
            return start  // or collect all matches
        else if pattern < substr:
            r = mid - 1
        else:
            l = mid + 1
    return -1  // not found
```
