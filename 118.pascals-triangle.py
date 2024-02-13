class Solution:
    def generate(self, numRows: int) -> list[list[int]]:
        if numRows == 1: return [[1]]
        if numRows == 2: return [[1], [1,1]]
        pas = [[1]]
        for i in range(1, numRows):
            new_row = [1]
            for j in range(len(pas[i-1])-1):
                new_row.append(pas[i-1][j] + pas[i-1][j+1])
            new_row.append(1)
            pas.append(new_row)
        return pas





if __name__ == '__main__':
    num_rows = 5
    print(Solution().generate(num_rows))
