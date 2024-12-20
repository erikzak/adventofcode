"""AOC day 20"""
from __future__ import annotations
import sys
from queue import Queue
from typing import Generator


Coords = tuple[int, int]


class Node:
    """Map node."""
    def __init__(self, coords: Coords, value: str) -> None:
        self.coords = coords
        self.x, self.y = coords
        self.value = value
        self.parent: Node = None  # For BFS/DFS

    def __str__(self) -> str:
        return f"Node({self.x}, {self.y})"

    def __repr__(self) -> str:
        return f"Node(({self.x}, {self.y}), {self.value})"

    def __eq__(self, other: Node) -> bool:
        return self.coords == other.coords and self.value == other.value

    def copy(self) -> Node:
        """Returns a copy of the node without parent."""
        return Node(self.coords, self.value)

    def manhattan_distance_to(self, other: Node) -> int:
        """Returns manhattan distance to another node."""
        return abs(self.x - other.x) + abs(self.y - other.y)

    def get_path(self) -> list[Coords]:
        """Returns a list of nodes from start to self."""
        path = [self.coords]
        node = self
        while (node.parent):
            path.append(node.parent.coords)
            node = node.parent
        return reversed(path)


class Map:
    """Map stuff."""
    def __init__(self, data: list[str]) -> None:
        self.index: dict[Coords, Node] = {}
        self.start: Coords = None
        self.end: Coords = None
        for y, line in enumerate(reversed(data)):
            for x, char in enumerate(line):
                if (char == "S"):
                    self.start = (x, y)
                    char = "."
                elif (char == "E"):
                    self.end = (x, y)
                    char = "."
                node = Node((x, y), char)
                self.index[(x, y)] = node
        self.dim = [len(data[0].strip()), len(data)]
        self._path: set[Coords] = {}

    def __str__(self) -> str:
        output = ""
        for y in range(self.dim[1] - 1, -1, -1):
            for x in range(self.dim[0]):
                node = self.get_node((x, y))
                if (node.coords in self._path):
                    output += "O" if (node.value != "#") else "1"
                else:
                    output += node.value
            output += "\n"
        return output

    @property
    def nodes(self) -> Generator[Node]:
        """Returns a generator of map nodes."""
        yield from self.index.values()

    def get_node(self, coords: Coords) -> Node | None:
        """Returns node at given coordinates, or None if out of bounds."""
        return self.index.get(coords)

    def draw(self) -> None:
        """Outputs the map to command line."""
        print(str(self))

    def draw_path(self, path: list[Coords]) -> None:
        self._path = set(path)
        self.draw()
        self._path = None

    def get_manhattan_neighbors(self, node: Node) -> list[Node | None]:
        """Returns list of neighboring nodes inside map boundary in XY direction only."""
        neighbors = [
            self.get_node((node.x - 1, node.y)),
            self.get_node((node.x + 1, node.y)),
            self.get_node((node.x, node.y - 1)),
            self.get_node((node.x, node.y + 1)),
        ]
        return [n for n in neighbors if n]

    def bfs(self, start: Node, end: Node) -> int | None:
        """
        Breadth-first search, returns shortest manhattan distance from start to end.

        Returns None if no path found.
        """
        explored: set[Coords] = {start.coords}
        queue = Queue()
        queue.put((start, 0))
        while (not queue.empty()):
            node: Node
            distance: int
            node, distance = queue.get()
            if (node == end):
                return node.get_path()
            for neighbor in self.get_manhattan_neighbors(node):
                if (neighbor.coords in explored or neighbor.value == "#"):
                    continue
                neighbor = neighbor.copy()
                explored.add(neighbor.coords)
                neighbor.parent = node
                queue.put((neighbor, distance + 1))
        return None

    def get_cheat_window(self, coords: Coords, duration: int) -> list[Node]:
        """Returns potential cheat exits based on duration."""
        exits = []
        for dx in range(-duration, duration + 1):
            for dy in range(-duration, duration + 1):
                check = (coords[0] + dx, coords[1] + dy)
                if (abs(coords[0] - check[0]) + abs(coords[1] - check[1]) <= duration):
                    node = self.get_node(check)
                    if (node and node.value != "#"):
                        exits.append(node)
        return exits

    def find_cheats(self, duration: int = 2, threshold: int = 100) -> int:
        """
        Checks potential wall hacks and uses reference race time to find time saved.
        """
        start, end = self.get_node(self.start), self.get_node(self.end)
        reference: dict[Coords, int] = {
            coords: time
            for time, coords in enumerate(self.bfs(start, end))
        }
        cheats_used = set(tuple[Coords, Coords])
        cheats_by_time: dict[int, int] = {}  # {picoseconds: num cheats}
        cheats: int = 0
        for from_coords in reference:
            from_node = self.get_node(from_coords)
            for to_node in self.get_cheat_window(from_node.coords, duration):
                if (
                        to_node.coords not in reference
                        or reference[to_node.coords] < reference[from_node.coords]
                ):
                    continue
                cheat = (from_node.coords, to_node.coords)
                if (cheat in cheats_used):
                    continue
                time_saved = (
                    reference[to_node.coords] - reference[from_node.coords]
                    - from_node.manhattan_distance_to(to_node)
                )
                if (time_saved < threshold):
                    continue
                cheats += 1
                if (time_saved not in cheats_by_time):
                    cheats_by_time[time_saved] = 0
                cheats_by_time[time_saved] += 1
                cheats_used.add(cheat)
        return cheats


def get_input(file_path: str) -> Map:
    """Returns content from puzzle string file tailored for today's puzzle."""
    with open(file_path, "r", encoding="utf-8") as reader:
        return Map([line.strip() for line in reader.readlines()])


def solve_part1(maze: Map) -> int:
    """Solution part 1."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    return maze.find_cheats(duration=2, threshold=1 if (test) else 100)


def solve_part2(maze: Map) -> int:
    """Solution part 2."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    return maze.find_cheats(duration=20, threshold=50 if (test) else 100)


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    puzzle_input = get_input(input_file)
    print(solve_part1(puzzle_input))
    print(solve_part2(puzzle_input))


if (__name__ == "__main__"):
    main()
