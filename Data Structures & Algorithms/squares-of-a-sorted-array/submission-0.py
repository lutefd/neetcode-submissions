class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        for n in nums:
            nums[nums.index(n)] = n * n
        nums.sort()
        return nums