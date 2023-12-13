//! Advent of Code 2023 day 13

use std::collections::HashMap;


struct Map {
    grid: HashMap<Point, char>,
    dim: [usize; 2],
}


impl Map {
    fn new(input: &str) -> Map {
        let mut grid: HashMap<Point, char> = HashMap::new();
        let mut max_x: usize = 0;
        let mut max_y: usize = 0;
        for (y, line) in input.lines().rev().enumerate() {
            for (x, value) in line.trim().chars().enumerate() {
                let pnt: Point = Point::from([x, y]);
                grid.insert(pnt, value);
                max_x = x;
            }
            max_y = y;
        }
        let dim: [usize; 2] = [max_x + 1, max_y + 1];
        Map { grid, dim }
    }

    fn x_is_mirror(&self, x: usize, offset: usize, smudge: &Point) -> bool {
        if offset > x { return true; }
        let from: usize = x - offset;
        let to: usize = x + 1 + offset;
        if to >= self.dim[0] { return true; }
        for y in 0..self.dim[1] {
            let left_pnt: Point = Point::from([from, y]);
            let right_pnt: Point = Point::from([to, y]);
            let mut left_value: char = self.grid[&left_pnt];
            if left_pnt == *smudge {
                left_value = if left_value == '.' { '#' } else { '.' };
            }
            let mut right_value: char = self.grid[&right_pnt];
            if right_pnt == *smudge {
                right_value = if right_value == '.' { '#' } else { '.' };
            }
            if left_value != right_value {
                return false;
            }
        }
        true
    }

    fn y_is_mirror(&self, y: usize, offset: usize, smudge: &Point) -> bool {
        if offset > y { return true; }
        let from: usize = y - offset;
        let to: usize = y + 1 + offset;
        if to >= self.dim[1] { return true; }
        for x in 0..self.dim[0] {
            let bottom_pnt: Point = Point::from([x, from]);
            let top_pnt: Point = Point::from([x, to]);
            let mut bottom_value: char = self.grid[&bottom_pnt];
            if bottom_pnt == *smudge {
                bottom_value = if bottom_value == '.' { '#' } else { '.' };
            }
            let mut top_value: char = self.grid[&top_pnt];
            if top_pnt == *smudge {
                top_value = if top_value == '.' { '#' } else { '.' };
            }
            if bottom_value != top_value {
                return false;
            }
        }
        true
    }

    fn smudge_in_reflection(&self, idx: usize, offset: usize, smudge: &Point, x_axis: bool) -> bool {
        if smudge.x == self.dim[0] && smudge.y == self.dim[1] { return true; }
        let smudge_idx: usize = if x_axis { smudge.x } else { smudge.y };
        smudge_idx >= idx - offset && smudge_idx <= idx + 1 + offset
    }

    fn find_mirror(&self, smudge: &Point) -> (i32, i32) {
        for x in 0..self.dim[0]-1 {
            let max_offset: usize = if self.dim[0] - 2 - x < x { self.dim[0] - 2 - x } else { x };
            if (0..=max_offset).any(|offset| !self.x_is_mirror(x, offset, smudge)) {
                continue;
            }
            if self.smudge_in_reflection(x, max_offset, smudge, true) {
                return (x as i32, -1);
            }
        }
        for y in 0..self.dim[1]-1 {
            let max_offset: usize = if self.dim[1] - 2 - y < y { self.dim[1] - 2 - y } else { y };
            if (0..=max_offset).any(|offset| !self.y_is_mirror(y, offset, smudge)) {
                continue;
            }
            if self.smudge_in_reflection(y, max_offset, smudge, false) {
                return (-1, y as i32);
            }
        }
        (-1, -1)
    }
}


#[derive(Eq, Hash, PartialEq)]
struct Point {
    x: usize,
    y: usize,
}

impl std::fmt::Display for Point {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(f, "({}, {})", self.x, self.y)
    }
}

impl Point {
    fn from(coords: [usize; 2]) -> Point {
        Point{ x: coords[0], y: coords[1] }
    }
}


fn process_smudge_patterns(pattern: &Map) -> u32 {
    for x in 0..pattern.dim[0] {
        for y in 0..pattern.dim[1] {
            let smudge: Point = Point::from([x, y]);
            let mirror: (i32, i32) = pattern.find_mirror(&smudge);
            if mirror.0 >= 0 {
                return mirror.0 as u32 + 1;
            }
            if mirror.1 >= 0 {
                return (pattern.dim[1] as u32 - mirror.1 as u32 - 1) * 100;
            }
        }
    }
    panic!("no smudged mirror found")
}


fn part_uno(_input: &str) -> u32 {
    // Solves part 1
    let mut answer: u32 = 0;
    for pattern_input in _input.replace("\r\n", "\n").split("\n\n") {
        let pattern: Map = Map::new(pattern_input);
        let smudge = Point::from([pattern.dim[0], pattern.dim[1]]);
        let mirror: (i32, i32) = pattern.find_mirror(&smudge);
        if mirror.0 >= 0 {
            answer += mirror.0 as u32 + 1;
            continue;
        }
        if mirror.1 >= 0 {
            answer += (pattern.dim[1] as u32 - mirror.1 as u32 - 1) * 100;
            continue;
        }
        panic!("no mirror found");
    }
    println!("{answer}");
    answer
}


fn part_dos(_input: &str) -> u32 {
    // Solves part 2
    let mut answer: u32 = 0;
    for pattern_input in _input.replace("\r\n", "\n").split("\n\n") {
        let pattern: Map = Map::new(pattern_input);
        answer += process_smudge_patterns(&pattern);
    }
    println!("{answer}");
    answer
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
        let input: &str = include_str!("../test.txt").trim();
        assert_eq!(part_uno(input), 405);
        assert_eq!(part_dos(input), 400);
    }
}
