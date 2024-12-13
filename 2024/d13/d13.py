"""AOC day 13"""
import sys


Coords = tuple[int, int]
Vec = tuple[int, int]


class ClawMachine:
    """Machine stuff."""
    def __init__(self, lines: list[str]) -> None:
        self._lines = lines
        self.a: Vec = self._parse_input_line(lines[0])
        self.b: Vec = self._parse_input_line(lines[1])
        self.prize: Coords = self._parse_input_line(lines[2], "=")

    def __str__(self) -> str:
        return f"ClawMachine(A: {self.a}, B: {self.b}, Prize: {self.prize})"

    def __repr__(self) -> str:
        return f"ClawMachine({self._lines})"

    @staticmethod
    def _parse_input_line(line: str, operator: str = "+") -> Vec | Coords:
        """Returns XY button distances."""
        button_part = line[line.index(f"X{operator}") + 2:].replace(f"Y{operator}", "")
        return tuple(map(int, [value.strip() for value in button_part.split(",")]))

    def get_lowest_cost(self) -> int | None:
        """
        Solve for a and b given the following equations:

        Xa * a + Xb * b = Xp
        Ya * a + Yb * b = Yp

        Cramer's rule:
        a = (Xp * Yb - Xb * Yp) / (Xa * Yb - Xb * Ya)
        b = (Xa * Yp - Xp * Ya) / (Xa * Yb - Xb * Ya)
        """
        determinant = self.a[0] * self.b[1] - self.b[0] * self.a[1]
        a = (self.prize[0] * self.b[1] - self.b[0] * self.prize[1]) / determinant
        b = (self.a[0] * self.prize[1] - self.prize[0] * self.a[1]) / determinant
        if (
                a == int(a) and b == int(b)
                and (self.a[0] * a + self.b[0] * b, self.a[1] * a + self.b[1] * b) == self.prize
        ):
            return int(a * 3 + b)
        return None

    def brute_force(self) -> int | None:
        """Brute force."""
        a_hammer_time = None
        b_hammer_time = None
        for a in range(min(self.prize[0] // self.a[0], self.prize[1] // self.a[1]) + 1):
            for b in range(min(self.prize[0] // self.b[0], self.prize[1] // self.b[1]) + 1):
                if (
                        a * self.a[0] + b * self.b[0] == self.prize[0]
                        and a * self.a[1] + b * self.b[1] == self.prize[1]
                ):
                    a_hammer_time = a * 3 + b
                    break
        for b in range(min(self.prize[0] // self.b[0], self.prize[1] // self.b[1]) + 1):
            for a in range(min(self.prize[0] // self.a[0], self.prize[1] // self.a[1]) + 1):
                if (
                        a * self.a[0] + b * self.b[0] == self.prize[0]
                        and a * self.a[1] + b * self.b[1] == self.prize[1]
                ):
                    b_hammer_time = a * 3 + b
                    break
        if (a_hammer_time is not None and b_hammer_time is not None):
            return min(a_hammer_time, b_hammer_time)
        return None


def get_input(file_path: str) -> list[ClawMachine]:
    """Returns content from puzzle string file tailored for today's puzzle."""
    machines: list[ClawMachine] = []
    with open(file_path, "r", encoding="utf-8") as reader:
        for machine in reader.read().split("\n\n"):
            machines.append(ClawMachine([line.strip() for line in machine.split("\n")]))
    return machines


def solve_part1(machines: list[ClawMachine]) -> int:
    """Solution part 1."""
    min_cost = 0
    for machine in machines:
        min_cost += machine.brute_force() or 0
    return min_cost


def solve_part2(machines: list[ClawMachine]) -> int:
    """Solution part 2."""
    min_cost = 0
    for machine in machines:
        machine.prize = (machine.prize[0] + 10000000000000, machine.prize[1] + 10000000000000)
        min_cost += machine.get_lowest_cost() or 0
    return min_cost


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    puzzle_input = get_input(input_file)
    print(solve_part1(puzzle_input))
    print(solve_part2(puzzle_input))


if (__name__ == "__main__"):
    main()
