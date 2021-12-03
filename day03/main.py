from typing import Literal
from collections import defaultdict


def thin_input(lines: list[str], pos: int, mode: Literal["oxygen", "co2"]) -> list[str]:
    count: dict[str, int] = defaultdict(int)

    for line in lines:
        count[line[pos]] += 1

    print(count)

    if mode == "co2":

        if count["0"] > count["1"]:
            return [line for line in lines if line[pos] == "1"]
        else:
            return [line for line in lines if line[pos] == "0"]
    else:
        if count["1"] < count["0"]:
            return [line for line in lines if line[pos] == "0"]
        else:
            return [line for line in lines if line[pos] == "1"]


if __name__ == "__main__":
    with open("./input", "r") as file:
        input_ = [l.strip() for l in file.readlines()]

    lines = input_[:]
    pos = 0
    while len(lines) > 1:
        lines = thin_input(lines, pos, "oxygen")
        pos += 1

    oxygen = eval("0b" + lines.pop())

    lines = input_[:]
    pos = 0
    while len(lines) > 1:
        lines = thin_input(lines, pos, "co2")
        pos += 1

    co2 = eval("0b" + lines.pop())

    print(oxygen * co2)
