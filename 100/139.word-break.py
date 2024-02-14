class Solution:
    def wordBreak(self, s: str, wordDict: list[str]) -> bool:
        dp = [True] + [False] * (len(s))
        for i in range(len(s)):
            for j in range(i, len(s)):
                if s[i:j+1] in wordDict:
                    dp[j+1] = dp[i] or dp[j+1]
        return dp[-1]


if __name__ == '__main__':
    test = "cars"
    d = ["car","ca","rs"]
    print(Solution().wordBreak(test, d))
