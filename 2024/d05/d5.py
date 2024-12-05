"""AOC day 5"""
import sys


def get_input(file_path: str) -> str:
    """Returns content from puzzle string file tailored for today's puzzle."""
    with open(file_path, "r", encoding="utf-8") as reader:
        puzzle = reader.read()
    puzzle_parts = puzzle.split("\n\n")
    rules = {}
    for rule in puzzle_parts[0].split("\n"):
        before, after = rule.split("|")
        rules.setdefault(int(after), set()).add(int(before))
    updates = [[int(p) for p in line.split(",")] for line in puzzle_parts[1].split("\n")]
    return rules, updates


def solve_part1(rules: dict[int, set[int]], updates: list[list[int]]) -> int:
    """Solution part 1."""
    result = 0
    for update in updates:
        if (not check_update(update, rules)):
            continue
        result += update[len(update) // 2]
    return result


def check_update(update: list[int], rules: dict[int, set[int]]) -> bool:
    """Returns True if update is correctly ordered."""
    before = []
    after = list(update)
    for page in update:
        if (page in rules and any(previous_page in after for previous_page in rules[page])):
            return False
        before.append(after.pop(0))
    return True


def solve_part2(rules: dict[int, set[int]], updates: list[list[int]]) -> int:
    """Solution part 2."""
    result = 0
    for update in updates:
        if (check_update(update, rules)):
            continue
        update = fix_update(update, rules)
        result += update[len(update) // 2]
    return result


def fix_update(update: list[int], rules: dict[int, set[int]]) -> list[int]:
    """Returns a correctly ordered update."""
    return update


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    rules, updates = get_input(input_file)
    print(solve_part1(rules, updates))
    print(solve_part2(rules, updates))


if (__name__ == "__main__"):
    main()
