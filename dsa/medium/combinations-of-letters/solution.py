class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if not digits:
            return []

        letters = {
            '2': "abc", '3': "def", '4': "ghi",
            '5': "jkl", '6': "mno", '7': "pqrs",
            '8': "tuv", '9': "wxyz"
        }

        def generate(idx, temp):
            if idx == len(digits):
                result.append(temp)
                return

            for ch in letters[digits[idx]]:
                generate(idx + 1, temp + ch)

        result = []
        generate(0, "")
        return result