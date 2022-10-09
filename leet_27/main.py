
class Solution:
    def removeElement(self, nums, val):
        i = 0
        while i <= len(nums) - 1:
            if nums[i] == val:
                del nums[i]
            else:
                i += 1
        return len(nums)

s = Solution()
nums = [3, 2, 2, 3]
val = 3
print(s.removeElement(nums, val))
