"""AOC day 19"""
import sys
from functools import cache


def get_input(file_path: str) -> tuple[set[str], list[str]]:
    """Returns content from puzzle string file tailored for today's puzzle."""
    with open(file_path, "r", encoding="utf-8") as reader:
        towels, patterns = reader.read().split("\n\n")
    towels = tuple(towel.strip() for towel in towels.split(","))
    patterns = [line.strip() for line in patterns.split("\n")]
    return towels, patterns


@cache
def find_towel_combinations(pattern: str, towels: tuple[str, ...]) -> int:
    """Find all possible combinations of towels for a pattern."""
    if (not pattern):
        return 1

    return sum(
        find_towel_combinations(pattern[len(towel):], towels)
        for towel in towels
        if (pattern.startswith(towel))
    )


def solve_part1(puzzle: tuple[tuple[str, ...], list[str]]) -> int:
    """Solution part 1."""
    towels, patterns = puzzle
    possible_designs: int = 0
    for pattern in patterns:
        if (find_towel_combinations(pattern, towels)):
            possible_designs += 1
    return possible_designs


def solve_part2(puzzle: tuple[set[str], list[str]]) -> int:
    """Solution part 2."""
    towels, patterns = puzzle
    possible_designs: int = 0
    for pattern in patterns:
        possible_designs += find_towel_combinations(pattern, towels)
    return possible_designs


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    puzzle_input = get_input(input_file)
    print(solve_part1(puzzle_input))
    print(solve_part2(puzzle_input))


if (__name__ == "__main__"):
    main()
