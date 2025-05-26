import time

def print_grid(grid):
    i = 0
    for row in grid:
        for col in range(9):
            if col % 3 == 0 and col != 0:
                print("|", end=" ")
            if col == 8:
                print(row[col])
            else:
                print(row[col], end=" ")
        if i % 3 == 2 and i != 8:
            print("- - - + - - - + - - -")
        i += 1

def is_valid(grid, row, col, num):
    # Check if the number is not in the given row and column
    for x in range(9):
        if grid[row][x] == num or grid[x][col] == num:
            return False
    
    # Check the 3x3 box
    start_row, start_col = 3 * (row // 3), 3 * (col // 3)
    for i in range(3):
        for j in range(3):
            if grid[start_row + i][start_col + j] == num:
                return False
    
    return True

def solve_sudoku(grid):
    for row in range(9):
        for col in range(9):
            if grid[row][col] == 0:  # Find an empty cell
                for num in range(1, 10):  # Try numbers 1-9
                    if is_valid(grid, row, col, num):
                        grid[row][col] = num
                        if solve_sudoku(grid):
                            return True
                        grid[row][col] = 0  # Backtrack
                return False
    return True

# Input Sudoku puzzle (0 represents empty cells)
sudoku_grid = [
    [0, 6, 1, 3, 0, 4, 8, 0, 0],
    [0, 0, 0, 0, 6, 0, 5, 1, 4],
    [4, 2, 9, 0, 8, 5, 3, 0, 0],
    [0, 0, 0, 7, 5, 3, 0, 8, 2],
    [0, 0, 0, 0, 0, 8, 0, 0, 3],
    [0, 0, 0, 4, 0, 0, 0, 0, 1],
    [9, 0, 4, 5, 3, 1, 0, 7, 0],
    [2, 5, 0, 0, 0, 9, 0, 0, 0],
    [0, 1, 0, 8, 0, 7, 0, 4, 5]
]

# sudoku_grid = [
#     [0, 0, 0, 0, 0, 0, 0, 0, 0],
#     [0, 0, 0, 0, 0, 0, 0, 0, 0],
#     [0, 0, 0, 0, 0, 0, 0, 0, 0],
#     [0, 0, 0, 0, 0, 0, 0, 0, 0],
#     [0, 0, 0, 0, 0, 0, 0, 0, 0],
#     [0, 0, 0, 0, 0, 0, 0, 0, 0],
#     [0, 0, 0, 0, 0, 0, 0, 0, 0],
#     [0, 0, 0, 0, 0, 0, 0, 0, 0],
#     [0, 0, 0, 0, 0, 0, 0, 0, 0]
# ]

start_time = time.time()  # Start timing
if solve_sudoku(sudoku_grid):
    execution_time = time.time() - start_time
    print(f"Solved Sudoku: {execution_time:.6f} seconds\n")
    print_grid(sudoku_grid)
else:
    print("No solution exists.")