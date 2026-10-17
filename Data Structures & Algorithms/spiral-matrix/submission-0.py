class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        result = []
        while matrix:
            result.extend(matrix.pop(0))

            if matrix and matrix[0]:
                for sublist in matrix:
                    result.append(sublist.pop())

            if matrix:
                result.extend(matrix.pop()[::-1])
            print(matrix)
            if matrix and matrix[0]:
                for sublist in reversed(matrix):
                    result.append(sublist.pop(0))

        return result