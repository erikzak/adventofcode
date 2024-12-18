"""AOC day 18"""
from __future__ import annotations
import sys
from queue import Queue


Coords = tuple[int, int]


class Node:
    """Map node."""
    def __init__(self, coords: Coords, value: str) -> None:
        self.coords = coords
        self.x, self.y = coords
        self.value = value

    def __str__(self) -> str:
        return f"Node({self.x}, {self.y})"

    def __repr__(self) -> str:
        return f"Node(({self.x}, {self.y}), {self.value})"

    def __eq__(self, other: Node) -> bool:
        return self.coords == other.coords and self.value == other.value


class Map:
    """Map stuff."""
    def __init__(self) -> None:
        self.index: dict[Coords, Node] = {}
        self.dim = (7, 7) if (any(arg in ["-t", "--test"] for arg in sys.argv)) else (71, 71)
        for y in range(self.dim[1]):
            for x in range(self.dim[0]):
                node = Node((x, y), ".")
                self.index[(x, y)] = node

    def __str__(self) -> str:
        output = ""
        for y in range(self.dim[1] - 1, -1, -1):
            for x in range(self.dim[0]):
                node = self.get_node((x, y))
                output += node.value
            output += "\n"
        return output

    def get_node(self, coords: Coords) -> Node | None:
        """Returns node at given coordinates, or None if out of bounds."""
        return self.index.get(coords)

    def draw(self) -> None:
        """Outputs the map to command line."""
        print(str(self))

    def corrupt_memory(self, byte_list: list[str], ticks: int = None) -> None:
        """Corrupts memory with input byte list."""
        for tick, line in enumerate(byte_list, start=1):
            x, y = list(map(int, line.split(",")))
            y = self.dim[1] - 1 - y  # Positive y-axis is up
            node = self.get_node((x, y))
            node.value = "#"
            if (ticks and tick == ticks):
                break

    def get_manhattan_neighbors(self, node: Node) -> list[Node | None]:
        """Returns list of neighboring nodes inside map boundary in XY direction only."""
        neighbors = [
            self.get_node((node.x - 1, node.y)),
            self.get_node((node.x + 1, node.y)),
            self.get_node((node.x, node.y - 1)),
            self.get_node((node.x, node.y + 1)),
        ]
        return [n for n in neighbors if n]

    def breadth_first_manhattan_search(self, start: Node, end: Node) -> int | None:
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
                return distance
            for neighbor in self.get_manhattan_neighbors(node):
                if (neighbor.value == "#" or neighbor.coords in explored):
                    continue
                explored.add(neighbor.coords)
                queue.put((neighbor, distance + 1))
        return None


def get_input(file_path: str) -> list[str]:
    """Returns content from puzzle string file tailored for today's puzzle."""
    with open(file_path, "r", encoding="utf-8") as reader:
        return reader.readlines()


def solve_part1(byte_list: list[str]) -> int:
    """Solution part 1."""
    memory_space = Map()
    ticks = 12 if (any(arg in ["-t", "--test"] for arg in sys.argv)) else 1024
    memory_space.corrupt_memory(byte_list, ticks)
    start = memory_space.get_node((0, memory_space.dim[1] - 1))
    end = memory_space.get_node((memory_space.dim[0] - 1, 0))
    return memory_space.breadth_first_manhattan_search(start, end)


def solve_part2(byte_list: list[str]) -> int:
    """Solution part 2."""
    memory_space = Map()
    start = memory_space.get_node((0, memory_space.dim[1] - 1))
    end = memory_space.get_node((memory_space.dim[0] - 1, 0))
    byte: str
    for byte in byte_list:
        memory_space.corrupt_memory([byte])
        if (memory_space.breadth_first_manhattan_search(start, end) is None):
            break
    return byte


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    puzzle_input = get_input(input_file)
    print(solve_part1(puzzle_input))
    print(solve_part2(puzzle_input))


if (__name__ == "__main__"):
    main()
