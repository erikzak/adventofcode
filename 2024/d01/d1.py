"""AOC day 1"""
import sys


def solve_part1(puzzle: str) -> int:
    """Solution part 1."""
    left, right = [], []
    for line in puzzle:
        split = line.strip().split()
        left.append(int(split[0]))
        right.append(int(split[1]))
    distance = sum(abs(l-r) for l, r in zip(sorted(left), sorted(right)))
    return distance


def solve_part2(puzzle: str) -> int:
    """Solution part 2."""
    numbers, counts = {}, {}
    for line in puzzle:
        left, right = (int(v) for v in line.strip().split())
        numbers[left] = numbers.get(left, 0) + 1
        counts[right] = counts.get(right, 0) + 1
    similarity = sum(number * counts.get(number, 0) * numbers[number] for number in numbers)
    return similarity


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
