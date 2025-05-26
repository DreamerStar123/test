import tkinter as tk
import random
import time

def print_grid(grid):
    for row in grid:
        print(" ".join(str(num) for num in row))

def is_valid(grid, row, col, num):
    for x in range(9):
        if grid[row][x] == num or grid[x][col] == num:
            return False
    start_row, start_col = 3 * (row // 3), 3 * (col // 3)
    for i in range(3):
        for j in range(3):
            if grid[start_row + i][start_col + j] == num:
                return False
    return True

def solve_sudoku(grid):
    for row in range(9):
        for col in range(9):
            if grid[row][col] == 0:
                for num in range(1, 10):
                    if is_valid(grid, row, col, num):
                        grid[row][col] = num
                        if solve_sudoku(grid):
                            return True
                        grid[row][col] = 0
                return False
    return True

def fill_diagonal_box(grid, row, col):
    nums = list(range(1, 10))
    random.shuffle(nums)
    for i in range(3):
        for j in range(3):
            grid[row + i][col + j] = nums.pop()

def generate_sudoku():
    grid = [[0 for _ in range(9)] for _ in range(9)]
    for i in range(3):
        fill_diagonal_box(grid, i * 3, i * 3)
    solve_sudoku(grid)
    return grid

def display_sudoku():
    start_time = time.time()  # Start timing
    grid = generate_sudoku()
    end_time = time.time()  # End timing
    execution_time = end_time - start_time

    for row in range(9):
        for col in range(9):
            entry = tk.Entry(root, width=2, font=('Arial', 20), justify='center')
            entry.grid(row=row, column=col)
            entry.insert(0, grid[row][col] if grid[row][col] != 0 else '')

    time_label.config(text=f"Generation Time: {execution_time:.6f} seconds")

# Create the main window
root = tk.Tk()
root.title("Sudoku Generator")

# Create a button to generate Sudoku
button = tk.Button(root, text="Generate Sudoku", command=display_sudoku, font=('Arial', 14))
button.grid(row=9, column=0, columnspan=9)

# Create a label to display generation time
time_label = tk.Label(root, text="", font=('Arial', 14))
time_label.grid(row=10, column=0, columnspan=9)

# Start the main event loop
root.mainloop()