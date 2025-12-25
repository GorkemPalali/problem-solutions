# Longest Balanced Parentheses

**Difficulty:** Medium  
**Category:** Stack  
**Language:** Go


## Problem Summary
Given a string consisting only of `'('` and `')'`, find the length of the longest contiguous substring that forms a valid (balanced) parentheses sequence.


## Key Insight
A valid parentheses substring must:
- Never close before opening
- End with all opened parentheses closed

A stack can be used to track:
- Indices of unmatched `'('`
- The last position where a valid substring cannot start


## Algorithm
1. Initialize a stack with `-1` as a sentinel value.
2. Traverse the string:
   - On `'('`: push its index onto the stack.
   - On `')'`: pop the stack.
     - If the stack becomes empty, push the current index as a new base.
     - Otherwise, calculate the length using `currentIndex - stackTop`.
3. Track the maximum valid length encountered.


## Complexity Analysis
- **Time Complexity:** `O(n)`
- **Space Complexity:** `O(n)` (stack)

