class Solution:
    def summaryRanges(self, nums: list[int]) -> list[str]:
        ranges = []
        r = []
        for i in range(len(nums)):
            if i == 0:
                r = [nums[0]]
            elif nums[i] - nums[i-1] == 1:
                r.append(nums[i])
            elif nums[i] == nums[i-1]:
                continue
            else:
                ranges.append(r)
                r = [nums[i]]

        if len(r) > 0:
            ranges.append(r)

        result = []
        for r in ranges:
            if len(r) == 1:
                result.append(f"{r[0]}")
            else:
                result.append(f"{r[0]}->{r[-1]}")

        return result

if __name__ == '__main__':
    print(Solution().summaryRanges([0,1,2,4,5,7]))
