from bisect import insort

class Solution:
    def findKthLargest(self, nums: list[int], k: int) -> int:
        l = []
        for num in nums:
            if len(l) < k:
                insort(l, num)
                continue

            if num > l[0]:
                l.pop(0)
                insort(l, num)

        return l[0]


if __name__ == '__main__':
    print(Solution().findKthLargest([3,2,3,1,2,4,5,5,6], 4))
