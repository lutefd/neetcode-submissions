class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        subpointer = 0
        for i in range(len(nums)):
            if nums[i] != val:
                nums[subpointer] = nums[i]
                subpointer += 1
        return subpointer