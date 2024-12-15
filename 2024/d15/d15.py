"""AOC day 15"""
from __future__ import annotations
import sys
from typing import Generator


Coords = tuple[int, int]


class Lanternfish:
    """Lanternfish stuff."""
    VECTORS = {"<": (-1, 0), "v": (0, -1), ">": (1, 0), "^": (0, 1)}

    def __init__(self, start: Coords, moves: list[str]) -> None:
        self.start = start
        self.x, self.y = start
        self.moves = moves

    @property
    def coords(self) -> Coords:
        """Returns current coordinates."""
        return (self.x, self.y)

    def _next(self, coords: Coords, direction: str) -> Coords:
        """Returns next coords in move direction."""
        return tuple(map(sum, zip(coords, self.VECTORS[direction])))

    def move(self, direction: str, warehouse: Map) -> None:
        """Moves the lanternfish, pushing crates if possible."""
        next_coords: Coords = self._next(self.coords, direction)
        next_node: Node = warehouse.get_node(next_coords)
        if (not next_node or next_node.value == "#"):
            return
        if (next_node.value == "."):
            self.x, self.y = next_coords
            return
        if (next_node.value in "O[]"):
            pushed = self.push_rock(next_node, warehouse, direction)
            if (pushed):
                self.x, self.y = next_coords
            return
        raise ValueError(f"Unhandled node value for next move: {next_node.value}")

    def push_rock(self, rock: Node, warehouse: Map, direction: str) -> bool:
        """Tries to push rock(s) ahead of lanternfish, returns True if successful."""
        if (rock.value == "."):
            # Already pushed
            return True
        if (rock.value == "O" or direction in "<>"):
            # Horizontal rock pushing is the same for both parts
            return self.push_p1(rock, warehouse, direction)
        return self.push_p2(rock, warehouse, direction)

    def push_p1(self, rock, warehouse, direction) -> bool:
        next_node = warehouse.get_node(self._next(rock.coords, direction))
        if (next_node.value == "#"):
            return False
        if (
                (next_node.value in "O[]" and self.push_rock(next_node, warehouse, direction))
                or next_node.value == "."
        ):
            next_node.value = rock.value
            rock.value = "."
            return True
        return False

    def push_p2(self, rock, warehouse, direction) -> bool:
        other_rock_coords = (
            rock.coords[0] + (1 if rock.value == "[" else -1), rock.coords[1]
        )
        next_nodes = [
            warehouse.get_node(self._next(rock.coords, direction)),
            warehouse.get_node(self._next(other_rock_coords, direction))
        ]
        if (any(node.value == "#" for node in next_nodes)):
            return False
        if (
                all(node.value == "." for node in next_nodes)
                or all(self.inspect_rock(node, warehouse, direction) for node in next_nodes)
        ):
            # Next rock can be pushed vertically
            for node in next_nodes:
                if (node.value in "[]"):
                    self.push_rock(node, warehouse, direction)
            next_nodes[0].value = rock.value
            next_nodes[1].value = "]" if (rock.value == "[") else "["
            rock.value = "."
            warehouse.get_node(other_rock_coords).value = "."
            return True
        return False

    def inspect_rock(self, rock: Node, warehouse: Map, direction: str) -> bool:
        """Returns True if rock node can be moved *vertically* (or is empty space)."""
        if (rock.value == "."):
            return True
        other_rock_coords = (rock.coords[0] + (1 if rock.value == "[" else -1), rock.coords[1])
        next_nodes = [
            warehouse.get_node(self._next(rock.coords, direction)),
            warehouse.get_node(self._next(other_rock_coords, direction))
        ]
        if (any(node.value == "#" for node in next_nodes)):
            return False
        return (all(
            node.value == "." or (
                node.value in "[]"
                and self.inspect_rock(node, warehouse, direction)
            )
            for node in next_nodes
        ))


class Node:
    """Map node."""
    def __init__(self, coords: Coords, value: str) -> None:
        self.coords = coords
        self.x, self.y = coords
        self.value = value

    def __str__(self) -> str:
        return f"({self.coords}, {self.value})"

    def __repr__(self) -> str:
        return f"({self.coords}, {self.value})"


class Map:
    """Map stuff."""
    def __init__(self, puzzle_input: list[str], part2: bool = False) -> None:
        data, lanternfish_moves = puzzle_input.split("\n\n")
        data = data.split("\n")
        lanternfish_moves = lanternfish_moves.replace("\n", "").strip()
        self.index: dict[Coords, Node] = {}
        for y, line in enumerate(reversed(data)):
            line = self.widen_line(line) if (part2) else line
            for x, char in enumerate(line):
                if (char == "@"):
                    self.lanternfish = Lanternfish((x, y), lanternfish_moves)
                    char = "."
                node = Node((x, y), char)
                self.index[(x, y)] = node
        self.dim = [len(line), len(data)]

    def __str__(self) -> str:
        output = ""
        for y in range(self.dim[1] - 1, -1, -1):
            for x in range(self.dim[0]):
                node = self.get_node((x, y))
                output += node.value if (node.coords != self.lanternfish.coords) else "@"
            output += "\n"
        return output

    @staticmethod
    def widen_line(line: str):
        """Widens line for part 2."""
        wider_line = ""
        for char in line:
            if (char == "#"):
                wider_line += "##"
            elif (char == "O"):
                wider_line += "[]"
            elif (char == "@"):
                wider_line += "@."
            else:
                wider_line += ".."
        return wider_line

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

    def move_fish(self) -> None:
        """Moves lanternfish all moves."""
        for move in self.lanternfish.moves:
            self.lanternfish.move(move, self)

    @property
    def box_sum(self) -> int:
        """Returns sum of all boxes' GPS coordinates."""
        box_sum = 0
        for node in self.nodes:
            if (node.value not in "O["):
                continue
            box_sum += 100 * (self.dim[1] - node.coords[1] - 1) + node.coords[0]
        return box_sum


def get_input(file_path: str, part2: bool = False) -> Map:
    """Returns content from puzzle string file tailored for today's puzzle."""
    with open(file_path, "r", encoding="utf-8") as reader:
        return Map(reader.read(), part2)


def solve_part1(warehouse_map: Map) -> int:
    """Solution part 1."""
    warehouse_map.move_fish()
    return warehouse_map.box_sum


def solve_part2(warehouse_map: Map) -> int:
    """Solution part 2."""
    warehouse_map.move_fish()
    return warehouse_map.box_sum


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    puzzle_input = get_input(input_file)
    print(solve_part1(puzzle_input))
    puzzle_input = get_input(input_file, True)
    print(solve_part2(puzzle_input))


if (__name__ == "__main__"):
    main()
