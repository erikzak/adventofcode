"""AOC day 12"""
from __future__ import annotations
import sys
from queue import Queue
from typing import Generator


Coords = tuple[int, int]


class Node:
    """Map node."""
    def __init__(self, coords: Coords, value: str) -> None:
        self.coords: Coords = coords
        self.x, self.y = coords
        self.value: str = value
        self.region_id: int = None

    def __str__(self) -> str:
        return f"Node({self.x}, {self.y})"

    def __repr__(self) -> str:
        return f"Node(({self.x}, {self.y}), {self.value}, {self.region_id})"


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
        self.regions: list[Region] = []

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

    def tag_regions(self) -> None:
        """Tags regions of nodes."""
        region_id = 0
        for node in self.nodes:
            if (node.region_id is not None):
                continue
            self.regions.append(self.tag_region_nodes(node, region_id))
            region_id += 1

    def tag_region_nodes(self, start: Node, region_id: int) -> Region:
        """
        BFS for tagging neighboring nodes with region id.
        """
        start.region_id = region_id
        explored: set[Coords] = {start.coords}
        nodes: dict[Coords, Node] = {start.coords: start}
        queue = Queue()
        queue.put(start)
        while (not queue.empty()):
            node: Node = queue.get()
            for neighbor in self.get_manhattan_neighbors(node):
                if (
                        neighbor is None
                        or neighbor.region_id is not None
                        or neighbor.coords in explored
                        or neighbor.value != start.value
                ):
                    continue
                neighbor.region_id = region_id
                nodes[neighbor.coords] = neighbor
                explored.add(neighbor.coords)
                queue.put(neighbor)
        return Region(region_id, nodes)


class Region:
    """Defines a region of nodes."""
    def __init__(self, region_id: int, nodes: dict[Coords, Node]):
        self.id = region_id
        self.nodes = nodes
        self.manhattan_vectors = ((1, 0), (0, 1), (-1, 0), (0, -1))
        self.corner_vectors = (
            ((-1, 0), (-1, 1), (0, 1)),  # upper left
            ((0, 1), (1, 1), (1, 0)),  # upper right
            ((1, 0), (1, -1), (0, -1)),  # lower right
            ((0, -1), (-1, -1), (-1, 0)),  # lower left
        )

    def get_manhattan_neighbors(self, coords: Coords) -> list[Coords]:
        """Returns list of neighboring coords in manhattan directions."""
        return [(coords[0] + v[0], coords[1] + v[1]) for v in self.manhattan_vectors]

    def count_fences(self) -> int:
        """Returns number of fences in region."""
        fences: int = 0
        for coords in self.nodes.keys():
            for neighbor in self.get_manhattan_neighbors(coords):
                if (neighbor not in self.nodes):
                    fences += 1
        return fences

    def count_corners(self) -> int:
        """Returns number of corners in region."""
        corners: int = 0
        for coords in self.nodes.keys():
            for vec in self.corner_vectors:
                corner_coords = [(coords[0] + v[0], coords[1] + v[1]) for v in vec]
                corner_nodes = [node in self.nodes for node in corner_coords]
                if (
                        corner_nodes[0] == corner_nodes[2]
                        and (not corner_nodes[0] or not corner_nodes[1])
                ):
                    corners += 1
        return corners

    def get_price(self, discount: bool = False) -> int:
        """Returns price of region."""
        if (discount):
            return len(self.nodes) * self.count_corners()
        return len(self.nodes) * self.count_fences()


def get_input(file_path: str) -> Map:
    """Returns content from puzzle string file tailored for today's puzzle."""
    with open(file_path, "r", encoding="utf-8") as reader:
        return Map([line.strip() for line in reader.readlines() if line.strip()])


def solve_part1(garden_map: Map) -> int:
    """Solution part 1."""
    garden_map.tag_regions()
    return sum(region.get_price() for region in garden_map.regions)


def solve_part2(garden_map: Map) -> int:
    """Solution part 2."""
    return sum(region.get_price(discount=True) for region in garden_map.regions)


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    garden_map = get_input(input_file)
    print(solve_part1(garden_map))
    print(solve_part2(garden_map))


if (__name__ == "__main__"):
    main()
