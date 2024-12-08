"""AOC day 8"""
from __future__ import annotations
import sys
from typing import Generator


Coords = tuple[int, int]


class Node:
    """Map node."""
    def __init__(self, coords: Coords, value: str) -> None:
        self.coords = coords
        self.x, self.y = coords
        self.value = value
        self.antinodes = set()

    def __str__(self) -> str:
        return f"Node({self.x}, {self.y})"

    def __repr__(self) -> str:
        return f"Node(({self.x}, {self.y}), {self.value})"

    def is_collinear(self, n1: Node, n2: Node) -> bool:
        """Returns True if node is collinear with two others."""
        return self.x * (n1.y - n2.y) + n1.x * (n2.y - self.y) + n2.x * (self.y - n1.y) == 0

    def distance_to(self, other: Node) -> int:
        """Returns manhattan distance to another node."""
        return abs(self.x - other.x) + abs(self.y - other.y)

    def is_antinode(self, antennas: list[Node], any_distance: bool = False) -> bool:
        """Returns True if node is an antinode of any of the given antennas."""
        for i, a1 in enumerate(antennas[:-1]):
            for a2 in antennas[i + 1:]:
                d1 = self.distance_to(a1)
                d2 = self.distance_to(a2)
                if (
                        self.is_collinear(a1, a2)
                        and (any_distance or (d1 == d2 * 2 or d2 == d1 * 2))
                ):
                    return True
        return False


class Map:
    """Map class."""
    def __init__(self, data: list[str]) -> None:
        self.index: dict[Coords, Node] = {}
        self.antennas: dict[str, list[Node]] = {}
        x, y = None, None
        for y, line in enumerate(reversed(data)):
            for x, char in enumerate(line):
                node = Node((x, y), char)
                self.index[(x, y)] = node
                if (char != "."):
                    self.antennas.setdefault(char, []).append(node)
        self.dim = [len(data[0].strip()), len(data)]

    def __str__(self) -> str:
        output = ""
        for y in range(self.dim[1] - 1, -1, -1):
            for x in range(self.dim[0]):
                node = self.get_node((x, y))
                output += "#" if (node.antinodes and node.value == ".") else node.value
            output += "\n"
        return output

    @property
    def nodes(self) -> Generator[Node]:
        """Returns a generator of map nodes."""
        yield from self.index.values()

    def get_node(self, coords: Coords) -> Node:
        """Returns node at given coordinates, or None if out of bounds."""
        return self.index.get(coords)

    def get_antenna_nodes(self, signal: str) -> list[Node]:
        """
        Returns a list of antenna nodes sending the given signal or None
        if no antennas.
        """
        if (signal not in self.antennas):
            return
        return self.antennas[signal]

    def draw(self) -> None:
        """Outputs the map to command line."""
        print(str(self))


def get_input(file_path: str) -> Map:
    """Returns content from puzzle string file tailored for today's puzzle."""
    with open(file_path, "r", encoding="utf-8") as reader:
        return Map([line.strip() for line in reader.readlines() if line.strip()])


def solve_part1(antenna_map: Map) -> int:
    """Solution part 1."""
    for node in antenna_map.nodes:
        for signal, antennas in antenna_map.antennas.items():
            if (node.value == signal or signal in node.antinodes):
                continue
            if (node.is_antinode(antennas)):
                node.antinodes.add(signal)
                continue
    return len([node for node in antenna_map.nodes if node.antinodes])


def solve_part2(antenna_map: Map) -> int:
    """Solution part 2."""
    for node in antenna_map.nodes:
        for signal, antennas in antenna_map.antennas.items():
            if (signal in node.antinodes):
                continue
            if (node.is_antinode(antennas, any_distance=True)):
                node.antinodes.add(signal)
                continue
    return len([node for node in antenna_map.nodes if node.antinodes])


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    antenna_map = get_input(input_file)
    print(solve_part1(antenna_map))
    antenna_map = get_input(input_file)
    print(solve_part2(antenna_map))


if (__name__ == "__main__"):
    main()
