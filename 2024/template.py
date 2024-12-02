"""AOC day #"""
import sys


def solve_part1(input: str) -> int:
    """Solution part 1."""


def solve_part2(input: str) -> int:
    """Solution part 2."""


def main():
    """Runs puzzle solutions."""
    dev_flag = any(arg in ["-d", "--dev"] for arg in sys.argv)
    input_file = "test.txt" if (dev_flag) else "input.txt"

    with open(input_file, "r", encoding="utf-8") as reader:
        puzzle_input = reader.read().strip()

    print(solve_part1(puzzle_input))
    print(solve_part2(puzzle_input))
