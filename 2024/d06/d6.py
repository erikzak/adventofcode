"""AOC day 6"""
from __future__ import annotations
import sys


Coord = tuple[int, int]


class Guard:
    """Keeps track of guard state."""
    def __init__(self, coords: Coord, facing: str):
        self._initial_coords = coords
        self._initial_facing = facing
        self.coords = coords
        self.rotation = ["N", "E", "S", "W"]
        if (facing not in self.rotation):
            raise ValueError(f"Invalid facing direction: {facing} not in {self.rotation}")
        self.facing = facing
        self.seen: set[Coord] = {coords}

    def copy(self) -> Guard:
        """Returns a copy of the guard object."""
        return Guard(self._initial_coords, self._initial_facing)

    def get_next_tile(self) -> Coord:
        """Returns the coordinates of the next tile."""
        if (self.facing == "N"):
            return (self.coords[0], self.coords[1] + 1)
        if (self.facing == "E"):
            return (self.coords[0] + 1, self.coords[1])
        if (self.facing == "S"):
            return (self.coords[0], self.coords[1] - 1)
        if (self.facing == "W"):
            return (self.coords[0] - 1, self.coords[1])
        raise ValueError(f"Invalid facing direction: {self.facing}")

    def move(self, guard_map: dict[Coord, str]) -> Coord:
        """
        Moves one step and checks if turn.

        Returns the new coordinates or None if out of map.
        """
        next_tile = self.get_next_tile()
        if (next_tile not in guard_map):
            return None
        self.coords = next_tile
        self.seen.add(self.coords)
        while guard_map.get(self.get_next_tile()) in ("#", "O"):
            self.turn()
        return self.coords

    def turn(self):
        """Turns guard according to 1518 logic."""
        self.facing = self.rotation[(self.rotation.index(self.facing) + 1) % len(self.rotation)]

    def patrol(self, guard_map: dict[Coord, str]) -> set[Coord]:
        """Moves guard until looping or map exit. Returns seen tiles."""
        states: set[tuple[Coord, str]] = {self.state}
        while (self.move(guard_map) in guard_map):
            if (self.state in states):
                break
            states.add(self.state)
        return self.seen

    @property
    def state(self) -> tuple[Coord, str]:
        """Returns the guard's current spatial state."""
        return (self.coords, self.facing)


def get_input(file_path: str) -> tuple[Guard, dict[Coord, str]]:
    """Returns content from puzzle string file tailored for today's puzzle."""
    with open(file_path, "r", encoding="utf-8") as reader:
        puzzle = [line.strip() for line in reader.readlines() if line.strip()]
    guard: Guard = None
    guard_map: dict[Coord, str] = {}
    for y, line in enumerate(reversed(puzzle)):
        for x, char in enumerate(line):
            guard_map[(x, y)] = char
            if (char == "^"):
                guard = Guard((x, y), "N")
    if (guard is None):
        raise ValueError("No guard found in map.")
    return guard, guard_map


def solve_part1(guard: Guard, guard_map: dict[Coord, str]) -> int:
    """Solution part 1."""
    return len(guard.patrol(guard_map))


def solve_part2(guard: Guard, guard_map: dict[Coord, str]) -> int:
    """Solution part 2."""
    combos = 0
    # Get initial guard path, and only test obstacles where guard has been
    path = guard.copy().patrol(guard_map)
    for coord, char in guard_map.items():
        if (char != "." or coord not in path):
            continue
        test_guard = guard.copy()
        test_map = guard_map.copy()
        test_map[coord] = "O"
        if (check_map_for_loop(test_guard, test_map)):
            combos += 1
    return combos


def check_map_for_loop(guard: Guard, test_map: dict[Coord, str]):
    """Checks if map makes guard get stuck in loop."""
    states: set[tuple[Coord, str]] = {guard.state}
    while (guard.move(test_map) in test_map):
        # If we've seen this state before, we're stuck in a loop
        if (guard.state in states):
            return True
        states.add(guard.state)
    return False


def print_map(guard_map: dict[Coord, str]):
    """Prints the guard map."""
    x_max, y_max = 0, 0
    for x, y in guard_map:
        x_max = max(x, x_max)
        y_max = max(y, y_max)
    for y in range(y_max, -1, -1):
        for x in range(x_max + 1):
            print(guard_map.get((x, y), " "), end="")
        print()
    print()


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    guard, guard_map = get_input(input_file)
    print(solve_part1(guard.copy(), guard_map))
    print(solve_part2(guard.copy(), guard_map))


if (__name__ == "__main__"):
    main()
