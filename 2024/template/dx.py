"""AOC day #"""
import sys


def solve_part1(puzzle: str) -> int:
    """Solution part 1."""


def solve_part2(puzzle: str) -> int:
    """Solution part 2."""


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    with open(input_file, "r", encoding="utf-8") as reader:
        puzzle_input = reader.readlines()
    print(solve_part1(puzzle_input))
    print(solve_part2(puzzle_input))


if (__name__ == "__main__"):
    main()
