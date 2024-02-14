class Solution:
    def get_skelton(self, s: str) -> list:
        d = {}
        i = 0
        for c in s:
            if not c in d:
                d[c] = i
                i += 1

        skelton = [d[c] for c in s]
        return skelton

    def isIsomorphic(self, s: str, t: str) -> bool:
        return self.get_skelton(s) == self.get_skelton(t)

if __name__ == '__main__':
    s = 'egg'
    t = 'add'
    print(Solution().isIsomorphic(s, t))
