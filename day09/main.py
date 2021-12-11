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

directions = [
        (-1, 0),
        (1, 0),
        # (-1, -1),
        # (1, 1),
        # (-1, 1),
        # (1, -1),
        (0, -1),
        (0, 1),
    ]

def bfs(grid: np.ndarray, i: int, j: int):
    visited = set()
    visit = {(i, j)}
    basin = set()

    while visit:
        ci, cj = visit.pop()
        height = grid[ci, cj]
        # print(f"{(ci, cj)} -> {height}")
        if height < 9:
            basin.add((ci, cj))
        visited.add((ci, cj))

        for di, dj in directions:
            if height == 9:
                continue
            ni, nj = ((ci + di, cj + dj))
            if ni < 0 or nj < 0 or ni > grid.shape[0] - 1 or nj > grid.shape[1] - 1:
                pass
            elif (ni, nj) not in visited:
                visit.add((ni, nj))

    return len(basin)

sizes = []
for x, y in minima_xy[:]:
    size = bfs(data, x, y)
    sizes.append(size)
    print(f"({x}, {y}) {size=}")

product = 1
for size in sorted(sizes)[-3:]:
    product *= size

print(f"Part 2: {product}")


