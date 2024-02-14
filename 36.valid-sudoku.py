import collections

class Solution:
    def isValidSudoku(self, board: list[list[str]]) -> bool:
        for row in board:
            dup = [k for k, v in collections.Counter(row).items() if v > 1 and k != '.']
            if len(dup) > 0:
                return False

        for i in range(9):
            col = [board[j][i] for j in range(9)]
            dup = [k for k, v in collections.Counter(col).items() if v > 1 and k != '.']
            if len(dup) > 0:
                return False

        for i in range(3):
            for j in range(3):
                sub_matrix = [board[i*3][j*3], board[i*3][j*3+1], board[i*3][j*3+2], board[i*3+1][j*3], board[i*3+1][j*3+1], board[i*3+1][j*3+2], board[i*3+2][j*3], board[i*3+2][j*3+1], board[i*3+2][j*3+2]]
                dup = [k for k, v in collections.Counter(sub_matrix).items() if v > 1 and k != '.']
                if len(dup) > 0:
                    return False

        return True

if __name__ == '__main__':
    board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
    print(Solution().isValidSudoku(board))
