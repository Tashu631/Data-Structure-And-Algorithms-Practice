// check if the string is palindrome or not
package main

func isPalindrome(s string, l, r int) bool {
	if l >= r {
		return true
	}

	if s[l] != s[r] {
		return false
	}
	return isPalindrome(s, l+1, r-1)
}

/* EXPLANATION:
String Palindrome Check using Recursion

Palindrome: A string that reads same forwards and backwards
Examples:
- "racecar" → palindrome
- "hello" → NOT palindrome
- "madam" → palindrome
- "noon" → palindrome

Algorithm - Two Pointer Approach:
1. Use two pointers: left (l) and right (r)
2. Compare characters at l and r
3. Move pointers inward: l++, r--
4. Repeat until pointers meet or characters don't match

Base Cases:
1. if l >= r: return true (all characters matched, it's palindrome)
2. if s[l] != s[r]: return false (mismatch found, NOT palindrome)

Recursive Case:
- If characters match, recursively check inner substring
- isPalindrome(s, l+1, r-1)

Execution Example: isPalindrome("racecar", 0, 6)
  l=0:'r', r=6:'r' → match, recurse(1,5)
  l=1:'a', r=5:'a' → match, recurse(2,4)
  l=2:'c', r=4:'c' → match, recurse(3,3)
  l=3, r=3 → l >= r → return true ✓

Execution Example: isPalindrome("hello", 0, 4)
  l=0:'h', r=4:'o' → 'h' ≠ 'o' → return false ✗

Time Complexity: O(n/2) = O(n) where n is string length
Space Complexity: O(n) - due to recursion depth

Key Insight:
- Only need to check first half against second half
- Stops as soon as mismatch is found
*/
