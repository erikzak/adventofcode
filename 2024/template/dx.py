"""AOC day #"""
import sys


def get_input(file_path: str) -> str:
    """Returns content from puzzle string file tailored for today's puzzle."""
    with open(file_path, "r", encoding="utf-8") as reader:
        return reader.readlines()


def solve_part1(puzzle: str) -> int:
    """Solution part 1."""


def solve_part2(puzzle: str) -> int:
    """Solution part 2."""


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    puzzle_input = get_input(input_file)
    print(solve_part1(puzzle_input))
    print(solve_part2(puzzle_input))


if (__name__ == "__main__"):
    main()
