import numpy as np

directions = [
    (-1, 0),
    (1, 0),
    (-1, -1),
    (1, 1),
    (-1, 1),
    (1, -1),
    (0, -1),
    (0, 1),
]

with open("./input", "r") as file:
    lines = [l.strip() for l in file.readlines()]
data = np.array([[int(l) for l in line] for line in lines])




def flash(energy: np.ndarray, flashed: np.ndarray, i: int, j: int) -> tuple[np.ndarray, np.ndarray]:
    for di, dj in directions:
        ni, nj = i + di, j + dj
        if ni < 0 or nj < 0 or ni > energy.shape[0] - 1 or nj > energy.shape[1] - 1:
            continue
        if flashed[ni, nj]:
            continue
        energy[ni, nj] += 1
    energy[i, j] = 0
    flashed[i, j] = 1
    return energy, flashed

energy = data.copy()
flashed = np.zeros_like(energy)
flashes = 0
for n in range(1000):
    if energy.sum() == 0:
        print(n)
        break
    flashes += np.count_nonzero(flashed)
    flashed = np.zeros_like(energy)  # reset flash tracker
    energy += 1   # increase energy level of each octopus

    i_coords, j_coords = np.where(energy > 9)
    coords = set(zip(i_coords, j_coords))
    
    while coords:
        i, j = coords.pop()
        if flashed[i, j]:
            continue
        energy, flashed = flash(energy, flashed, i, j)

        i_coords, j_coords = np.where(energy > 9)
        coords = set(zip(i_coords, j_coords))

print(flashes)