"""AOC day 10"""
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
        self.value = int(value)
        self.parent: Node = None

    def __str__(self) -> str:
        return f"Node({self.x}, {self.y})"

    def __repr__(self) -> str:
        return f"Node(({self.x}, {self.y}), {self.value})"

    def __eq__(self, other: Node) -> bool:
        return self.coords == other.coords and self.value == other.value

    def copy(self) -> Node:
        """Returns a copy of the node."""
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
        path.reverse()
        return path


class Map:
    """Map stuff."""
    def __init__(self, data: list[str]) -> None:
        self.index: dict[Coords, Node] = {}
        self.trailheads: list[Node] = []
        self.tops: list[Node] = []
        x, y = None, None
        for y, line in enumerate(reversed(data)):
            for x, char in enumerate(line):
                node = Node((x, y), char)
                self.index[(x, y)] = node
                if (node.value == 0):
                    self.trailheads.append(node)
                elif (node.value == 9):
                    self.tops.append(node)
        self.dim = [len(data[0].strip()), len(data)]

    def __str__(self) -> str:
        output = ""
        for y in range(self.dim[1] - 1, -1, -1):
            for x in range(self.dim[0]):
                node = self.get_node((x, y))
                output += str(node.value)
            output += "\n"
        return output

    @property
    def nodes(self) -> Generator[Node]:
        """Returns a generator of map nodes."""
        yield from self.index.values()

    def get_node(self, coords: Coords) -> Node:
        """Returns node at given coordinates, or None if out of bounds."""
        return self.index.get(coords)

    def draw(self) -> None:
        """Outputs the map to command line."""
        print(str(self))

    def get_manhattan_neighbors(self, node: Node) -> list[Node]:
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
                if (neighbor.coords not in explored and neighbor.value == node.value + 1):
                    neighbor = neighbor.copy()
                    explored.add(neighbor.coords)
                    neighbor.parent = node
                    queue.put((neighbor, distance + 1))
        return None

    def get_paths(self, start: Node, end: Node) -> list[list[Coords]]:
        """
        Breadth-first search implementation for all manhattan distance paths from
        start to end along incrementing elevation.

        Returns empty list if no path found.
        """
        paths: list[list[Coords]] = []
        explored_paths: set[tuple[Coords]] = set()
        queue = Queue()
        path: list[Coords] = [start.coords]
        queue.put(path)
        while (not queue.empty()):
            path = queue.get()
            if (tuple(path) in explored_paths):
                continue
            coords = path[-1]
            if (coords == end.coords):
                paths.append(path)
                continue
            node = self.get_node(coords)
            for neighbor in self.get_manhattan_neighbors(node):
                if (neighbor.value != node.value + 1):
                    continue
                new_path: list[Coords] = path.copy()
                new_path.append(neighbor.coords)
                queue.put(new_path)
            explored_paths.add(tuple(path))
        return paths


def get_input(file_path: str) -> Map:
    """Returns content from puzzle string file tailored for today's puzzle."""
    with open(file_path, "r", encoding="utf-8") as reader:
        return Map([line.strip() for line in reader.readlines()])


def solve_part1(topo_map: Map) -> int:
    """Solution part 2."""
    sum_scores = 0
    for trailhead in topo_map.trailheads:
        trailhead_score = 0
        for top in topo_map.tops:
            distance = topo_map.breadth_first_manhattan_search(trailhead, top)
            if (distance):
                trailhead_score += 1
        sum_scores += trailhead_score
    return sum_scores


def solve_part2(topo_map: Map) -> int:
    """Solution part 2."""
    sum_scores = 0
    for trailhead in topo_map.trailheads:
        trailhead_score = 0
        for top in topo_map.tops:
            paths = topo_map.get_paths(trailhead, top)
            if (paths):
                trailhead_score += len(paths)
        sum_scores += trailhead_score
    return sum_scores


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    puzzle_input = get_input(input_file)
    print(solve_part1(puzzle_input))
    print(solve_part2(puzzle_input))


if (__name__ == "__main__"):
    main()
