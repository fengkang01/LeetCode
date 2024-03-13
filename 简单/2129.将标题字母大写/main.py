class Solution:
    def capitalizeTitle(self, title: str) -> str:
        newWords = []
        for word in title.split(" "):
            newWord = word.lower()
            if len(word) > 2:
                newWords.append(newWord[0].upper() + newWord[1:])
            else:
                newWords.append(word.lower())
        words = " ".join(newWords)
        return words


if __name__ == '__main__':
    s = Solution()
    tests = [
        {
            "input": "capiTalIze tHe titLe",
            "output": "",
            "want": "Capitalize The Title",
            "passed": False,
        },
        {
            "input": "First leTTeR of EACH Word",
            "output": "",
            "want": "First Letter of Each Word",
            "passed": False,
        },
        {
            "input": "i lOve leetcode",
            "output": "",
            "want": "i Love Leetcode",
            "passed": False,
        }
    ]
    for test in tests:
        got = s.capitalizeTitle(test["input"])
        test["output"] = got
        if got == test["want"]:
            test["passed"] = True
    print(tests)


