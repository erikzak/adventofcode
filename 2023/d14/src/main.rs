//! Advent of Code 2023 day 14
use either::*;
use std::collections::{HashMap, HashSet};

#[derive(Clone)]
struct Platform {
    grid: HashMap<Point, char>,
    dim: [usize; 2],
    rocks: HashSet<Point>,
}


impl Platform {
    fn new(input: &str) -> Platform {
        let mut grid: HashMap<Point, char> = HashMap::new();
        let mut rocks: HashSet<Point> = HashSet::new();
        let mut max_x: usize = 0;
        let mut max_y: usize = 0;
        for (y, line) in input.lines().rev().enumerate() {
            for (x, value) in line.trim().chars().enumerate() {
                let pnt: Point = Point::from([x, y]);
                let c: char;
                if value == 'O' {
                    rocks.insert(pnt);
                    c = '.';
                } else {
                    c = value;
                }
                grid.insert(pnt, c);
                max_x = x;
            }
            max_y = y;
        }
        let dim: [usize; 2] = [max_x + 1, max_y + 1];
        Platform { grid, dim, rocks }
    }

    fn find_cycle(&mut self) -> (u32, u32) {
        let mut sequence: HashMap<String, u32> = HashMap::new();
        let mut cycle: u32 = 0;
        sequence.insert(format!("{self}"), cycle);
        loop {
            self.spin();
            cycle += 1;
            let grid: String = format!("{self}");
            if sequence.contains_key(&grid) {
                let start: u32 = sequence[&grid];
                let period = cycle - start;
                return (start, period);
            }
            sequence.insert(grid, cycle);
        }
    }

    fn spin_cycle(&mut self, count: u64) {
        for _ in 0..count {
            self.spin();
        }
    }

    fn spin(&mut self) {
        self.tilt('N');
        self.tilt('W');
        self.tilt('S');
        self.tilt('E');
    }

    fn tilt(&mut self, direction: char) {
        if direction == 'N' || direction == 'S' {
            let y_range: Either<std::ops::Range<usize>, std::iter::Rev<std::ops::Range<usize>>>;
            if direction == 'N' {
                y_range = either::Right((0..self.dim[1]-1).rev());
            } else {
                y_range = either::Left(1..self.dim[1]);
            }
            for y in y_range {
                for x in 0..self.dim[0] {
                    let from: Point = Point::from([x, y]);
                    if self.rocks.contains(&from) {
                        self.roll(from, direction);
                    }
                }
            }
        } else if direction == 'E' || direction == 'W' {  
            let x_range: Either<std::ops::Range<usize>, std::iter::Rev<std::ops::Range<usize>>>;
            if direction == 'E' {
                x_range = either::Right((0..self.dim[0]-1).rev());
            } else {
                x_range = either::Left(1..self.dim[0]);
            }
            for x in x_range {
                for y in 0..self.dim[1] {
                    let from: Point = Point::from([x, y]);
                    if self.rocks.contains(&from) {
                        self.roll(from, direction)
                    }
                }
            }
        } else {
            panic!("unknown direction: {direction}");
        }
    }

    fn roll(&mut self, pnt: Point, direction: char) {
        let path: Vec<Point>;
        if direction == 'N' {
            path = (pnt.y+1..self.dim[1]).map(|y: usize| Point::from([pnt.x, y])).collect();
        } else if direction == 'S' {
            path = (0..pnt.y).rev().map(|y: usize| Point::from([pnt.x, y])).collect();
        } else if direction == 'E' {
            path = (pnt.x+1..self.dim[0]).map(|x: usize| Point::from([x, pnt.y])).collect();
        } else if direction == 'W' {
            path = (0..pnt.x).rev().map(|x: usize| Point::from([x, pnt.y])).collect();
        } else {
            panic!("unknown direction: {direction}");
        }
        let mut from = pnt;
        for next_pnt in path {
            if self.rocks.contains(&next_pnt) || self.grid[&next_pnt] != '.' {
                break;
            }
            self.rocks.remove(&from);
            self.rocks.insert(next_pnt);
            from = next_pnt;
        }
    }

    fn calculate_load(&self) -> u64 {
        let mut total: u64 = 0;
        for rock in &self.rocks {
            total += rock.y as u64 + 1;
        }
        total
    }
}


impl std::fmt::Display for Platform {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        let mut chars: Vec<char> = Vec::new();
        for y in (0..self.dim[1]).rev() {
            for x in 0..self.dim[0] {
                let pnt: Point = Point::from([x, y]);
                chars.push(if self.rocks.contains(&pnt) { 'O' } else { self.grid[&pnt] });
            }
            chars.push('\n');
        }
        let string: String = chars.into_iter().collect();
        write!(f, "{string}")
    }
}


#[derive(Debug, Eq, Hash, PartialEq, Copy, Clone)]
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


fn part_uno(_input: &str) -> u64 {
    // Solves part 1
    let mut platform: Platform = Platform::new(_input);
    platform.tilt('N');
    let answer: u64 = platform.calculate_load();
    println!("{answer}");
    answer
}


fn part_dos(_input: &str) -> u64 {
    // Solves part 2
    let mut platform: Platform = Platform::new(_input);
    let (start, period) = platform.clone().find_cycle();
    platform.spin_cycle(start as u64 + (1000000000 - start as u64) % period as u64);
    let answer: u64 = platform.calculate_load();
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
        assert_eq!(part_uno(input), 136);
        assert_eq!(part_dos(input), 64);
    }
}
