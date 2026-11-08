class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        unique_subsets = []
        path = []
        def dfs(start_index:int) -> None:
            unique_subsets.append(path[:])

            for i in range(start_index, len(nums)):
                num = nums[i]
                path.append(num)
                dfs(i + 1)
                path.pop()
        
        dfs(0)
        return unique_subsets