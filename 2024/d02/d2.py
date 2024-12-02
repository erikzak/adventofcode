"""AOC day #"""
import sys


def solve_part1(puzzle: str) -> int:
    """Solution part 1."""
    safe_reports = 0
    for line in puzzle:
        report = [int(v) for v in line.strip().split()]
        if (check_report(report)):
            safe_reports += 1
    return safe_reports


def solve_part2(puzzle: str) -> int:
    """Solution part 2."""
    safe_reports = 0
    for line in puzzle:
        report = [int(v) for v in line.strip().split()]
        if (check_report(report, dampener=True)):
            safe_reports += 1
    return safe_reports


def check_report(report: list[int], dampener: bool = False) -> bool:
    """Returns True if report is safe. Dampener flag allows one level to be removed."""
    increasing = report[1] - report[0] > 0
    for i, v1 in enumerate(report[:-1]):
        v2 = report[i+1]
        if (
                not (1 <= abs(v2 - v1) <= 3)
                or (increasing and v2 < v1)
                or (not increasing and v2 > v1)
        ):
            if (dampener):
                # Edge cases for swapping increment direction
                if (check_report(report[1:]) or check_report(report[:1] + report[2:])):
                    return True
                return check_report(report[:i+1] + report[i+2:], False)
            return False
    return True


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
