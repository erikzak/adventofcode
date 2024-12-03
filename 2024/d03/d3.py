"""AOC day 3"""
import re
import sys


def get_input(file_path: str) -> str:
    """Returns content from puzzle string file tailored for today's puzzle."""
    with open(file_path, "r", encoding="utf-8") as reader:
        return reader.read()


def solve_part1(puzzle: str) -> int:
    """Solution part 1."""
    return sum_mults(puzzle)


def solve_part2(puzzle: str) -> int:
    """Solution part 2."""
    total = 0
    while ((dont_idx := puzzle.find("don't()")) != -1):
        do_idx = puzzle.find("do()", dont_idx)
        if (do_idx != -1):
            puzzle = puzzle[:dont_idx] + puzzle[do_idx + 4:]
        else:
            puzzle = puzzle[:dont_idx]
            break
    total += sum_mults(puzzle)
    return total


def sum_mults(text: str) -> int:
    """Returns sum of mul(x,y) in text."""
    total = 0
    for mul in re.findall(r"mul\([0-9]+,[0-9]+\)", text):
        split = mul.split(",")
        total += int(split[0].split("(")[-1]) * int(split[1][:-1])
    return total


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    puzzle_input = get_input(input_file)
    print(solve_part1(puzzle_input))
    print(solve_part2(puzzle_input))


if (__name__ == "__main__"):
    main()
