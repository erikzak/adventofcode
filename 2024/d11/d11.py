"""AOC day 11"""
from __future__ import annotations
import sys
from functools import cache


def get_input(file_path: str) -> str:
    """Returns content from puzzle string file tailored for today's puzzle."""
    with open(file_path, "r", encoding="utf-8") as reader:
        return reader.read().strip()


@cache
def blink(value: int, blinks: int) -> int:
    """Returns number of stones after n blinks."""
    if (blinks == 0):
        return 1
    if (value == 0):
        return blink(1, blinks - 1)
    value_str = str(value)
    str_len = len(value_str)
    if (str_len % 2 == 0):
        left = int(value_str[:str_len // 2])
        right = int(value_str[str_len // 2:])
        return blink(left, blinks - 1) + blink(right, blinks - 1)
    return blink(value * 2024, blinks - 1)


def solve_part1(puzzle: str) -> int:
    """Solution part 1."""
    return sum(blink(int(value), 25) for value in puzzle.split())


def solve_part2(puzzle: str) -> int:
    """Solution part 2."""
    return sum(blink(int(value), 75) for value in puzzle.split())


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    puzzle_input = get_input(input_file)
    print(solve_part1(puzzle_input))
    print(solve_part2(puzzle_input))


if (__name__ == "__main__"):
    main()
