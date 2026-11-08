class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        unique_subsets = []
        def dfs(candidates: list[int], start_index:int, path:list[int]) -> None:
            unique_subsets.append(path[:])

            for i in range(start_index, len(candidates)):
                num = nums[i]
                path.append(num)
                dfs(candidates, i + 1, path)
                path.pop()
        
        dfs(nums, 0, [])
        return unique_subsets