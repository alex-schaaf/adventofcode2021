import numpy as np
import sys
from copy import copy

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
    # game2 = game[:]

    winner = False
    winner_board = None
    winner_number = None
    boards_that_won = []
    all_won = False
    print(numbers)
    while numbers and not all_won:
        number = numbers.pop()

        game[1, game[0, :] == number] = 1


        for i, board in enumerate(game[1]):
            if i in boards_that_won:
                continue
            if 5 in board.sum(axis=0):
                col = np.where(board.sum(axis=0) == 5)[0][0]
                print(f'WINNER BOARD {i + 1} COL {col} with NUMBER {number}')
                winner = True
                winner_number = number
                winner_board = i
                boards_that_won.append(i)
            
            elif 5 in board.sum(axis=1):
                row = np.where(board.sum(axis=1) == 5)[0][0]
                print(f'WINNER BOARD {i + 1} ROW {row} with NUMBER {number}')
                winner = True
                winner_number = number
                winner_board = i
                boards_that_won.append(i)
            
            if game[1].shape[0] == len(boards_that_won):
                print(board)
                print(game[0, i, game[1, i, :] == 0].sum() * number)
                print(number)

    score = game[0, winner_board, game[1, winner_board, :] == 0].sum() * winner_number
    print(score)
