var letterCombinations = function (digits) {
    if (digits.length === 0) return [];

    const keyboard = {
        "2": "abc",
        "3": "def",
        "4": "ghi",
        "5": "jkl",
        "6": "mno",
        "7": "pqrs",
        "8": "tuv",
        "9": "wxyz"
    };

    const output = [];
    const backtrack = (current, index) => {
        if (index === digits.length) {
            output.push(current);
            return;}
        const letters = keyboard[digits[index]];
        for (let letter of letters) {
            backtrack(current + letter, index + 1);}
    };
    backtrack('', 0);
    return output;
};