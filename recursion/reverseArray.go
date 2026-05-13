// reverse an array
package main

func reverseArray(a []int, l, r int) []int {
	if l >= r {
		return a
	}
	a[l], a[r] = a[r], a[l]
	return reverseArray(a, l+1, r-1)
}

/* EXPLANATION:
Reverse Array using Recursion (Two-Pointer Technique)

Goal: Reverse an array in place using recursion

Example: reverseArray([1, 2, 3, 4], 0, 3) → [4, 3, 2, 1]

Algorithm - Two Pointer Approach:
1. Maintain two pointers: left (l) and right (r)
2. Swap elements at positions l and r
3. Move pointers inward: l++, r--
4. Recurse until pointers meet or cross

Base Case:
- if l >= r, return array (pointers have met/crossed, reversal complete)

Recursive Case:
1. Swap elements: a[l], a[r] = a[r], a[l]
2. Recurse with: reverseArray(a, l+1, r-1)

Execution Example: reverseArray([90, 2877, 67, 92], 0, 3)

Initial: [90, 2877, 67, 92]
         l=0                r=3

Step 1: Swap a[0] and a[3]
  90 ↔ 92
  Array: [92, 2877, 67, 90]
  Recurse with l=1, r=2

Step 2: Swap a[1] and a[2]
  2877 ↔ 67
  Array: [92, 67, 2877, 90]
  Recurse with l=2, r=1

Step 3: Base case l >= r
  l=2, r=1 (l > r)
  Return [92, 67, 2877, 90]

Final Result: [92, 67, 2877, 90]

Why This Works:
- Swap opposite ends: first with last, second with second-last, etc.
- After each swap, reduce array window by moving pointers inward
- Continue until middle is reached

In-Place Reversal:
- No extra array needed
- Only swaps elements directly
- Space efficient (not counting recursion stack)

Time Complexity: O(n/2) = O(n) where n is array length
Space Complexity: O(n/2) = O(n) for recursion depth = n/2

Key Insight:
- Simultaneous operations from both ends
- Efficient way to reverse without explicit loops
*/
