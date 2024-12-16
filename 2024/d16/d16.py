"""AOC day 16"""
from __future__ import annotations
import sys
from queue import Queue


Coords = tuple[int, int]
State = tuple[Coords, int]


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
        self.start: Coords = None
        self.end: Coords = None
        self.index: dict[Coords, Node] = {}
        for y, line in enumerate(reversed(data)):
            for x, char in enumerate(line):
                if (char == "S"):
                    self.start = (x, y)
                    char = "."
                if (char == "E"):
                    self.end = (x, y)
                    char = "."
                node = Node((x, y), char)
                self.index[(x, y)] = node
        self.dim = [len(data[0].strip()), len(data)]
        self._highlights = set()

    def __str__(self) -> str:
        output = ""
        for y in range(self.dim[1] - 1, -1, -1):
            for x in range(self.dim[0]):
                node = self.get_node((x, y))
                if (node.coords in self._highlights):
                    output += "O"
                else:
                    output += node.value
            output += "\n"
        return output

    def get_node(self, coords: Coords) -> Node | None:
        """Returns node at given coordinates, or None if out of bounds."""
        return self.index.get(coords)

    def draw(self, highlights: set[Coords]) -> None:
        """Outputs the map to command line."""
        self._highlights = highlights
        print(str(self))
        self._highlights = set()

    def get_reindeer_neighbors(self, node: Node, heading: int) -> list[Node | int | None]:
        """
        Returns list of neighboring nodes inside map boundary in XY direction
        only, based on reindeer logic.
        """
        neighbors = []
        for rotation in [0, 90, 270]:
            new_heading = (heading + rotation) % 360
            if (new_heading == 360):
                new_heading = 0
            neighbor: Node = None
            if (new_heading == 0):
                neighbor = self.get_node((node.x, node.y + 1))
            elif (new_heading == 90):
                neighbor = self.get_node((node.x + 1, node.y))
            elif (new_heading == 180):
                neighbor = self.get_node((node.x, node.y - 1))
            elif (new_heading == 270):
                neighbor = self.get_node((node.x - 1, node.y))
            neighbors.append((neighbor, new_heading))
        return neighbors

    def find_best_paths(self, start: Node, end: Node) -> int | None:
        """
        Breadth-first search, returns lowest reindeer score found using
        manhattan distance from start to end.

        Returns None if no path found.
        """
        best_score: int = float("inf")
        best_paths: set[Coords] = set()
        heading: int = 90  # 0 = N, 90 = E, 180 = S, 270 = W
        explored: dict[State, int] = {(start.coords, heading): 0}
        queue = Queue()
        queue.put((start, heading, 0))
        while (not queue.empty()):
            node: Node
            heading: int
            score: int
            node, heading, score = queue.get()
            if (node == end):
                if (score < best_score):
                    best_paths = set(node.get_path())
                    best_score = score
                elif (score == best_score):
                    best_paths.update(node.get_path())
                continue
            # No-scope 360, but no backtracking
            for neighbor, new_heading in self.get_reindeer_neighbors(node, heading):
                if (not neighbor or neighbor.value == "#"):
                    continue
                new_score = score + 1
                if (new_heading != heading):
                    new_score += 1000
                state = (neighbor.coords, new_heading)
                if (state not in explored or explored[state] >= new_score):
                    neighbor = neighbor.copy()
                    explored[state] = new_score
                    neighbor.parent = node
                    queue.put((neighbor, new_heading, new_score))
        return best_score, best_paths


def get_input(file_path: str) -> Map:
    """Returns content from puzzle string file tailored for today's puzzle."""
    with open(file_path, "r", encoding="utf-8") as reader:
        return Map([line.strip() for line in reader.readlines()])


def solve_part1(maze: Map) -> int:
    """Solution part 1."""
    start_node = maze.get_node(maze.start)
    end_node = maze.get_node(maze.end)
    return maze.find_best_paths(start_node, end_node)[0]


def solve_part2(maze: Map) -> int:
    """Solution part 2."""
    start_node = maze.get_node(maze.start)
    end_node = maze.get_node(maze.end)
    return len(maze.find_best_paths(start_node, end_node)[1])


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    puzzle_input = get_input(input_file)
    print(solve_part1(puzzle_input))
    print(solve_part2(puzzle_input))


if (__name__ == "__main__"):
    main()
