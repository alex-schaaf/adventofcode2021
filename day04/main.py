import numpy as np
import sys

if __name__ == "__main__":    
    with open("./input", "r") as file:
        lines = [l.strip() for l in file.readlines()]

    numbers = [int(n) for n in lines[0].split(',')][::-1]

    boards = []
    marks = []
    board = []
    mark = []
    for line in lines[2:]:
        if line == "":
            boards.append(board[:])
            marks.append(mark)
            board = []
            mark = []
            continue

        board.append([int(n) for n in line.split()])
        mark.append([0 for _ in line.split()])

    game = np.stack([boards, marks], axis=0)

    winner = False
    n_board = None
    while numbers and not winner:
        number = numbers.pop()

        game[1, game[0, :] == number] = 1

        for i, board in enumerate(game[1]):
            if 5 in board.sum(axis=0):
                col = np.where(board.sum(axis=0) == 5)[0][0]
                print(f'WINNER BOARD {i + 1} COL {col}')
                winner = True
                n_board = i
            
            elif 5 in board.sum(axis=1):
                row = np.where(board.sum(axis=1) == 5)[0][0]
                print(f'WINNER BOARD {i + 1} ROW {row}')
                winner = True
                n_board = i

    score = game[0, n_board, game[1, n_board, :] == 0].sum() * number
    print(score)