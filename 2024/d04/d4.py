"""AOC day 4"""
import sys


def get_input(file_path: str) -> str:
    """Returns content from puzzle string file tailored for today's puzzle."""
    with open(file_path, "r", encoding="utf-8") as reader:
        return [line.strip() for line in reader.readlines() if (line.strip())]


def solve_part1(puzzle: str) -> int:
    """Solution part 1."""
    count = 0
    vectors = ((1, 0), (1, -1), (0, -1), (-1, -1), (-1, 0), (-1, 1), (0, 1), (1, 1))
    for y, line in enumerate(puzzle):
        for x in range(len(line)):
            count += look_for_string("XMAS", x, y, puzzle, vectors)
    return count


def solve_part2(puzzle: str) -> int:
    """Solution part 2."""
    count = 0
    for y in range(1, len(puzzle) - 1):
        line = puzzle[y]
        for x in range(1, len(line) - 1):
            char = line[x]
            if (char != "A"):
                continue
            if (
                    (
                        look_for_string("MAS", x - 1, y - 1, puzzle, [(1, 1)])
                        or look_for_string("MAS", x + 1, y + 1, puzzle, [(-1, -1)])
                    ) and (
                        look_for_string("MAS", x - 1, y + 1, puzzle, [(1, -1)])
                        or look_for_string("MAS", x + 1, y - 1, puzzle, [(-1, 1)])
                    )
            ):
                count += 1
    return count


def look_for_string(
        string: str, x: int, y: int,
        lines: list[str], vectors: tuple[tuple[int]]
) -> int:
    """Looks for string in given directions from given (X, Y) index."""
    if (lines[y][x] != string[0]):
        return 0
    return sum(look_for_string_in_direction(string[1:], x, y, lines, vector) for vector in vectors)


def look_for_string_in_direction(
        string: str, x: int, y: int,
        lines: list[str], vector: tuple[int]
) -> bool:
    """Looks for string in a given direction from given (X, Y) index."""
    dx, dy = vector
    word_length = len(string)
    # Check for out of bounds coordinates
    if (
            any(coord < 0 for coord in (x + dx * word_length, y + dy * word_length))
            or any(coord >= len(lines) for coord in (y + dy * word_length, x + dx * word_length))
    ):
        return False
    factor = 1
    while (string):
        if (lines[y + dy * factor][x + dx * factor] != string[0]):
            return False
        string = string[1:]
        factor += 1
    return True


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    puzzle_input = get_input(input_file)
    print(solve_part1(puzzle_input))
    print(solve_part2(puzzle_input))


if (__name__ == "__main__"):
    main()
