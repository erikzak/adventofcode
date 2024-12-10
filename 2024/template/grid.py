"""Grid node/coordinate classes and helper functions."""
from __future__ import annotations
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
        x, y = None, None
        for y, line in enumerate(reversed(data)):
            for x, char in enumerate(line):
                node = Node((x, y), char)
                self.index[(x, y)] = node
        self.dim = [len(data[0].strip()), len(data)]

    def __str__(self) -> str:
        output = ""
        for y in range(self.dim[1] - 1, -1, -1):
            for x in range(self.dim[0]):
                node = self.get_node((x, y))
                output += node.value
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
                if (neighbor.coords not in explored):
                    neighbor = neighbor.copy()
                    explored.add(neighbor.coords)
                    neighbor.parent = node
                    queue.put((neighbor, distance + 1))
        return None
