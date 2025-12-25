# Communication Decryptor

**Difficulty:** Easy  
**Category:** HashMap, String Processing  
**Language:** Go


## Problem Summary
Given an encrypted message and a decryption map that defines substitutions for lowercase letters (`a-z`), decrypt the message.

Rules:
- Lowercase letters are replaced using the given mapping.
- Uppercase letters use the same mapping but remain uppercase.
- Non-alphabetic characters (spaces, punctuation, digits) remain unchanged.


## Approach
- Iterate through the message character by character.
- Use a hash map (`map[rune]rune`) for constant-time lookups.
- For uppercase letters:
  - Convert to lowercase
  - Apply the mapping
  - Convert the result back to uppercase
- Append characters efficiently using `strings.Builder`.


## Complexity Analysis
- **Time Complexity:** `O(n)`
- **Space Complexity:** `O(1)` (excluding output)
