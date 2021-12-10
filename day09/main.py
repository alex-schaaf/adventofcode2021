import numpy as np

with open("./input", "r") as file:
    lines = [l.strip() for l in file.readlines()]

data = []
for line in lines:
    data.append([int(l) for l in line])
data = np.array(data)

rows = data.shape[0]
cols = data.shape[1]
minima = []
minima_xy = []

for x in range(0, rows):
    for y in range(0, cols):
        n = data[x, y]
        isMinima = True

        if x > 0:
            isMinima &= n < data[x - 1, y]
        if x < rows - 1:
            isMinima &= n < data[x + 1, y]
        if y > 0:
            isMinima &= n < data[x, y - 1]
        if y < cols - 1:
            isMinima &= n < data[x, y + 1]
        
        if isMinima:
            minima.append(n)
            minima_xy.append((x, y))

print(f"Part 1: {sum([m + 1 for m in minima])}")