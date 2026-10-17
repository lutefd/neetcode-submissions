class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        count = 0
        curr = 0
        for num in nums:
            if num == 0:
                count = max(count, curr)
                curr = 0
            else:
                curr += 1
            count = max(count,curr)
        return count

            