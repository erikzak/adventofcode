//! Advent of Code 2023 day 10
use std::collections::{HashMap, HashSet};
use std::fmt;


#[derive(Debug)]
#[derive(Eq)]
#[derive(PartialEq)]
#[derive(Hash)]
struct Point {
    x: u8,
    y: u8,
}


impl fmt::Display for Point {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "({}, {})", self.x, self.y)
    }
}


impl Point {
    fn from(coords: [u8; 2]) -> Point {
        Point{ x: coords[0], y: coords[1] }
    }

    fn copy(&self) -> Point {
        Point{ x: self.x, y: self.y }
    }

    fn north(&self) -> Point {
        Point::from([self.x, self.y + 1])
    }

    fn south(&self) -> Point {
        Point::from([self.x, if self.y == 0 { 0 } else { self.y - 1 }])
    }

    fn east(&self) -> Point {
        Point::from([self.x + 1, self.y])
    }

    fn west(&self) -> Point {
        Point::from([if self.x == 0 { 0 } else { self.x - 1 }, self.y])
    }
}


struct Maze {
    pipe_map: HashMap<Point, char>,
    pipe_loop: HashSet<Point>,
    dim: [u8; 2],
    start: Point,
}


impl Maze {
    fn new(input: &str) -> Maze {
        let mut pipe_map: HashMap<Point, char> = HashMap::new();
        let mut pipe_loop: HashSet<Point> = HashSet::new();
        let mut start: Point = Point::from([0, 0]);
        let mut max_x: usize = 0;
        let mut max_y: usize = 0;
        for (y, line) in input.lines().rev().enumerate() {
            for (x, value) in line.trim().chars().enumerate() {
                let pnt: Point = Point::from([x as u8, y as u8]);
                if value == 'S' {
                    start = pnt.copy();
                }
                pipe_map.insert(pnt, value);
                max_x = x;
            }
            max_y = y;
        }
        let dim: [u8; 2] = [max_x as u8 + 1, max_y as u8 + 1];
        // Walk through once to map loop
        let mut maze = Maze{ pipe_map, pipe_loop: HashSet::new(), dim, start };
        let mut pnt: Point = maze.walk(&maze.start, &maze.start);
        pipe_loop.insert(pnt.copy());
        let mut last: Point = maze.start.copy();
        loop {
            let next_pnt: Point = maze.walk(&pnt, &last);
            pipe_loop.insert(next_pnt.copy());
            if next_pnt == maze.start {
                break;
            }
            last = pnt.copy();
            pnt = next_pnt;
        }
        maze.pipe_loop = pipe_loop;
        maze
    }

    fn walk(&self, pnt: &Point, from: &Point) -> Point {
        let pipe: char = self.pipe_map[pnt];
        if pipe == 'S' {
            if "|JL".chars().any(|c| c == self.pipe_map[&pnt.south()]) { return pnt.south(); }
            if "|F7".chars().any(|c| c == self.pipe_map[&pnt.north()]) { return pnt.north(); }
            if "-J7".chars().any(|c| c == self.pipe_map[&pnt.east()]) { return pnt.east(); }
            if "-LF".chars().any(|c| c == self.pipe_map[&pnt.west()]) { return pnt.west(); }
        }
        if pipe == '|' {
            return if pnt.south() == *from { pnt.north() } else { pnt.south() }
        }
        if pipe == '-' {
            return if pnt.west() == *from { pnt.east() } else { pnt.west() }
        }
        if pipe == 'L' {
            return if pnt.north() == *from { pnt.east() } else { pnt.north() }
        }
        if pipe == 'J' {
            return if pnt.north() == *from { pnt.west() } else { pnt.north() }
        }
        if pipe == '7' {
            return if pnt.south() == *from { pnt.west() } else { pnt.south() }
        }
        if pipe == 'F' {
            return if pnt.south() == *from { pnt.east() } else { pnt.south() }
        }
        panic!("didn't move")
    }

    fn on_loop(&self, pnt: &Point) -> bool {
        self.pipe_loop.contains(pnt)
    }

    fn at_edge(&self, pnt: &Point) -> bool {
        pnt.x <= 0 || pnt.y <= 0 ||
        pnt.x >= self.dim[0] - 1 || pnt.y >= self.dim[1] - 1
    }

    fn get_farthest_point(&self) -> u32 {
        let mut pnt: Point = self.walk(&self.start, &self.start);
        let mut last: Point = self.start.copy();
        let mut steps: u32 = 1;
        loop {
            steps += 1;
            let next_pnt: Point = self.walk(&pnt, &last);
            if next_pnt == self.start {
                break;
            }
            last = pnt.copy();
            pnt = next_pnt;
        }
        let farthest_point: u32 = (steps as f32 / 2.0).floor() as u32;
        farthest_point
    }

    fn get_enclosed_points(&self) -> u32 {
        let mut enclosed_points = 0;
        for x in 0..=self.dim[0] - 1 {
            for y in 0..=self.dim[1] - 1 {
                let pnt: Point = Point::from([x, y]);
                if self.contains_point(pnt) {
                    enclosed_points += 1;
                }
            }
        }
        enclosed_points
    }

    fn contains_point(&self, pnt: Point) -> bool {
        // Ray casting algorithm
        if self.at_edge(&pnt) || self.on_loop(&pnt) { return false };
        let edges: Vec<&str> = vec!("|JL", "|F7");
        for edge_case in edges.iter() {
            let mut crossings: u8 = 0;
            let ray: std::ops::RangeInclusive<u8>;
            if (0..=pnt.x).any(|i: u8| Point::from([i, pnt.y]) == self.start) {
                ray = pnt.x..=self.dim[0] - 1;
            } else {
                ray = 0..=pnt.x;
            }
            for x in ray {
                let check: &Point = &Point::from([x, pnt.y]);
                if self.on_loop(check) && edge_case.chars().any(|c| c == self.pipe_map[check]) {
                    crossings += 1;
                }
            }
            let inside: bool = if crossings == 0 || crossings % 2 == 0 { false } else { true };
            if inside {
                return inside;
            }
        }
        false
    }
}


fn part_uno(_input: &str) -> u32 {
    // Solves part 1
    let maze: Maze = Maze::new(_input);
    let farthest_point: u32 = maze.get_farthest_point();
    println!("{farthest_point}");
    farthest_point
}


fn part_dos(_input: &str) -> u32 {
    // Solves part 2
    let maze: Maze = Maze::new(_input);
    let enclosed_points: u32 = maze.get_enclosed_points();
    println!("{enclosed_points}");
    enclosed_points
}


fn main() {
    // Reads input and runs solve functions for parts
    let input: &str = include_str!("../input.txt").trim();
    part_uno(input);
    part_dos(input);
}

#[cfg(test)]
mod tests {
    use super::{part_uno, part_dos};
    #[test]
    fn test() {
        // Reads input and runs solve functions for parts
        let input1: &str = include_str!("../test1.txt").trim();
        assert_eq!(part_uno(input1), 8);
        let input2a: &str = include_str!("../test2a.txt").trim();
        assert_eq!(part_dos(input2a), 4);
        let input2b: &str = include_str!("../test2b.txt").trim();
        assert_eq!(part_dos(input2b), 8);
        let input2c: &str = include_str!("../test2c.txt").trim();
        assert_eq!(part_dos(input2c), 10);
    }
}
