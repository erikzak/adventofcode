"""AOC day 14"""
from __future__ import annotations
import sys
from functools import reduce


Coords = tuple[int, int]
Velocity = tuple[int, int]

MAP_DIM = (11, 7) if any(arg in ["-t", "--test"] for arg in sys.argv) else (101, 103)


class Robot:
    """Robot stuff."""
    def __init__(self, position: Coords, velocity: Velocity) -> None:
        self.x, self.y = (position[0], MAP_DIM[1] - position[1] - 1)  # I want positive Y up
        self.start_position: Coords = self.coords
        self.velocity: Velocity = velocity

    @property
    def coords(self) -> Coords:
        """Returns robots current coordinates."""
        return (self.x, self.y)

    def move(self, bathroom_map: Map) -> None:
        """Moves robot based on velocity."""
        self.x += self.velocity[0]
        self.y -= self.velocity[1]
        if (not bathroom_map.get_node(self.coords)):
            dim = bathroom_map.dim
            if (not 0 <= self.x < dim[0]):
                self.x = self.x - dim[0] if (self.x >= 0) else self.x + dim[0]
            if (not 0 <= self.y < dim[1]):
                self.y = self.y - dim[1] if (self.y >= 0) else self.y + dim[1]

    def reset(self) -> None:
        """Reset robot to initial coordinates."""
        self.x, self.y = self.start_position


class Node:
    """Map node."""
    def __init__(self, coords: Coords, value: str) -> None:
        self.coords = coords
        self.x, self.y = coords
        self.value = value


class Map:
    """Map stuff."""
    def __init__(self, robots: list[Robot]) -> None:
        self.robots: list[Robot] = robots
        self.index: dict[Coords, Node] = {}
        for y in range(MAP_DIM[1]):
            for x in range(MAP_DIM[0]):
                node = Node((x, y), ".")
                self.index[(x, y)] = node
        self.dim = MAP_DIM
        self.x_mid = MAP_DIM[0] // 2
        self.y_mid = MAP_DIM[1] // 2

    def __str__(self) -> str:
        output = ""
        for y in range(self.dim[1] - 1, -1, -1):
            for x in range(self.dim[0]):
                node = self.get_node((x, y))
                robots = len([r for r in self.robots if r.coords == node.coords])
                output += str(robots) if robots else node.value
            output += "\n"
        return output

    def get_node(self, coords: Coords) -> Node | None:
        """Returns node at given coordinates, or None if out of bounds."""
        return self.index.get(coords)

    def get_quadrant(self, coords: Coords) -> int | None:
        """Returns quadrant based on given coords, or None if in middle."""
        if (coords[0] < self.x_mid and coords[1] > self.y_mid):
            return 0
        if (coords[0] > self.x_mid and coords[1] > self.y_mid):
            return 1
        if (coords[0] < self.x_mid and coords[1] < self.y_mid):
            return 2
        if (coords[0] > self.x_mid and coords[1] < self.y_mid):
            return 3
        return None

    @property
    def security_factor(self) -> int:
        """Returns security factor based no robots per quadrant."""
        quadrants = {0: 0, 1: 0, 2: 0, 3: 0}
        for robot in self.robots:
            quadrant = self.get_quadrant(robot.coords)
            if (quadrant is not None):
                quadrants[quadrant] += 1
        return reduce(lambda x, y: x * y, quadrants.values())

    def draw(self) -> None:
        """Outputs the map to command line."""
        print(str(self))

    def move_robots(self) -> None:
        """Moves robot based on individual velocities."""
        for robot in self.robots:
            robot.move(self)


def get_input(file_path: str) -> Map:
    """Returns content from puzzle string file tailored for today's puzzle."""
    with open(file_path, "r", encoding="utf-8") as reader:
        lines = [line.strip() for line in reader.readlines()]
    robots = []
    for line in lines:
        pos_str, vel_str = line.split()
        position = list(map(int, (pos_str.split(",")[0].split("=")[1], pos_str.split(",")[-1])))
        velocity = list(map(int, (vel_str.split(",")[0].split("=")[1], vel_str.split(",")[-1])))
        robots.append(Robot(position, velocity))
    bathroom_map = Map(robots)
    return bathroom_map


def solve_part1(bathroom_map: Map) -> int:
    """Solution part 1."""
    for _ in range(100):
        bathroom_map.move_robots()
    return bathroom_map.security_factor


def solve_part2(bathroom_map: Map) -> int:
    """Solution part 2."""
    lowest_security_factor, seconds = float("inf"), float("inf")
    for robot in bathroom_map.robots:
        robot.reset()
    for i in range(1, 10000):
        bathroom_map.move_robots()
        security_factor = bathroom_map.security_factor
        if (security_factor < lowest_security_factor):
            lowest_security_factor = security_factor
            seconds = i
    return seconds


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    puzzle_input = get_input(input_file)
    print(solve_part1(puzzle_input))
    print(solve_part2(puzzle_input))


if (__name__ == "__main__"):
    main()
