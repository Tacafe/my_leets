class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        find_idx = -1
        matching = False

        i = 0
        j = 0
        while i < len(needle) and j < len(haystack):
            # print(j, i, haystack[j], needle[i])
            if matching and needle[i] == haystack[j]:
                i += 1
            elif not matching and needle[i] == haystack[j]:
                matching = True
                find_idx = j
                i += 1
            elif matching and needle[i] != haystack[j]:
                matching = False
                i = 0
                j = find_idx
                find_idx = -1
            else:
                matching = False
                find_idx = -1
                i = 0
            j += 1

        if i < len(needle):
            find_idx = -1

        return find_idx

        #########################
        # simple solution
        #########################
        # length_needle=len(needle)
        # for i in range(len(haystack)):
        #     test=haystack[i:length_needle+i]
        #     if test==needle:
        #         return i
        # return -1


if __name__ == '__main__':
    print(Solution().strStr("mississippi", "pi"))
