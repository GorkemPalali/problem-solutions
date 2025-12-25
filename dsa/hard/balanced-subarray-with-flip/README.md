# Balanced Subarray with One Flip

**Difficulty:** Hard  
**Category:** Prefix Sum, Hashing  
**Language:** Go


## Problem Summary
Given a binary array, find the length of the longest contiguous subarray that can be made balanced
(equal number of `0`s and `1`s) by flipping **at most one element**.


## Key Insight
- Convert values:
  - `0 → -1`
  - `1 → +1`
- A balanced subarray has a total sum of `0`.
- Flipping one element changes the sum by `±2`.


## Algorithm
1. Use a prefix sum array while traversing the input.
2. Store the **earliest index** at which each prefix sum occurs.
3. At each position, check:
   - Same prefix sum → no flip needed
   - Prefix sum ± 2 → one flip can balance
4. Track the maximum subarray length.


## Complexity Analysis
- **Time Complexity:** `O(n)`
- **Space Complexity:** `O(n)`
