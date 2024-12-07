"""AOC day 7"""
import sys
from itertools import product


class Test:
    """Test class, please ignore."""
    def __init__(self, test: int, equation: list[int]) -> None:
        self.test = test
        self.equation = equation

    def test_operators(self, operators: list[str]) -> int:
        """Returns the result of the equation with the given operators."""
        result = self.equation[0]
        for i, operator in enumerate(operators):
            if (operator == "+"):
                result += self.equation[i + 1]
            elif (operator == "*"):
                result *= self.equation[i + 1]
            elif (operator == "||"):
                result = int(str(result) + str(self.equation[i + 1]))
        return result

    def passes(self, permutations: list[list[str]]) -> bool:
        """
        Returns True if the equation can be made to match test with
        combinations of + and * operators.
        """
        for operators in permutations:
            if (self.test_operators(operators.copy()) == self.test):
                return True
        return False


def get_input(file_path: str) -> list[Test]:
    """Returns content from puzzle string file tailored for today's puzzle."""
    with open(file_path, "r", encoding="utf-8") as reader:
        puzzle = reader.readlines()
    tests = []
    for line in puzzle:
        split = line.split(":")
        tests.append(
            Test(
                int(split[0]),
                list(map(int, split[1].strip().split(" ")))
            )
        )
    return tests


def solve_part1(tests: list[Test]) -> int:
    """Solution part 1."""
    result = 0
    for test in tests:
        permutations = [list(p) for p in product(["+", "*"], repeat=len(test.equation) - 1)]
        if (test.passes(permutations)):
            result += test.test
    return result


def solve_part2(tests: list[Test]) -> int:
    """Solution part 2."""
    result = 0
    for test in tests:
        permutations = [list(p) for p in product(["+", "*", "||"], repeat=len(test.equation) - 1)]
        if (test.passes(permutations)):
            result += test.test
    return result


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    tests = get_input(input_file)
    print(solve_part1(tests))
    print(solve_part2(tests))


if (__name__ == "__main__"):
    main()
