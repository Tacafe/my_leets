from bisect import bisect_left

class Solution:
    def lengthOfLIS(self, nums: list[int]) -> int:
        piles = []

        for num in nums:
            # Binary Search for a number >= num
            insertion_point = bisect_left(piles, num)
            if insertion_point == len(piles):
                piles.append(num)
            else:
                piles[insertion_point] = num
            print(piles)

        return len(piles)



if __name__ == '__main__':
    print(Solution().lengthOfLIS([7,7,7,7]))
