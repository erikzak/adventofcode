//! Advent of Code 2023 day 11
use std::collections::HashSet;


struct Universe {
    grid: Vec<Vec<char>>,
    dim: [usize; 2],
    galaxies: HashSet<Point>,
    expansion_factor: u32,
    empty_x: HashSet<usize>,
    empty_y: HashSet<usize>,
}


impl std::fmt::Display for Universe {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        let lines: Vec<String> = self.grid.iter().rev()
            .map(|line| String::from_iter(line))
            .collect();
        write!(f, "{}", lines.join("\n"))
    }
}


impl Universe {
    fn new(input: &str, expansion_factor: u32) -> Universe {
        let mut grid: Vec<Vec<char>> = Vec::new();
        let mut galaxies: HashSet<Point> = HashSet::new();
        for (y, line) in input.lines().rev().enumerate() {
            let mut gridline: Vec<char> = Vec::new();
            for (x, c) in line.chars().enumerate() {
                if c == '#' {
                    galaxies.insert(Point{ x, y });
                }
                gridline.push(c);
            }
            grid.push(gridline);
        }
        let dim: [usize; 2] = [grid[0].len(), grid.len()];
        let mut universe = Universe{
            grid, dim, galaxies, expansion_factor,
            empty_x: HashSet::new(), empty_y: HashSet::new()
        };
        universe.find_empty_space();
        universe
    }

    fn find_empty_space(&mut self) {
        for x in (0..self.dim[0]).rev() {
            if (0..self.dim[1]).all(|y| self.grid[y][x] == '.') {
                self.empty_x.insert(x);
            }
        }
        self.dim[0] = self.grid[0].len();
        for y in (0..self.dim[1]).rev() {
            if self.grid[y].iter().all(|x| x == &'.') {
                self.empty_y.insert(y);
            }
        }
    }

    fn distance(&self, from: &Point, to: &Point) -> u128 {
        let mut x_distance: u128 = 0;
        let x_range: std::ops::Range<usize> = if to.x > from.x { from.x..to.x } else { to.x..from.x };
        for x in x_range {
            x_distance += (if self.empty_x.contains(&x) { self.expansion_factor } else { 1 }) as u128;
        }

        let mut y_distance: u128 = 0;
        let y_range: std::ops::Range<usize> = if to.y > from.y { from.y..to.y } else { to.y..from.y };
        for y in y_range {
            y_distance += (if self.empty_y.contains(&y) { self.expansion_factor } else { 1 }) as u128;
        }

        x_distance + y_distance
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


fn part_uno(_input: &str, expansion_factor: u32) -> u128 {
    // Solves part 1
    let mut answer: u128 = 0;
    let universe: Universe = Universe::new(_input, expansion_factor);
    let mut done: HashSet<&Point> = HashSet::new();
    for from in &universe.galaxies {
        for to in &universe.galaxies {
            if from == to || done.contains(to) {
                continue;
            }
            answer += universe.distance(from, to);
        }
        done.insert(&from);
    }
    println!("{answer}");
    answer
}


fn part_dos(_input: &str, expansion_factor: u32) -> u128 {
    // Solves part 2
    let mut answer: u128 = 0;
    let universe: Universe = Universe::new(_input, expansion_factor);
    let mut done: HashSet<&Point> = HashSet::new();
    for from in &universe.galaxies {
        for to in &universe.galaxies {
            if from == to || done.contains(to) {
                continue;
            }
            answer += universe.distance(from, to);
        }
        done.insert(&from);
    }
    println!("{answer}");
    answer
}


fn main() {
    // Reads input and runs solve functions for parts
    let input: &str = include_str!("../input.txt").trim();
    part_uno(input, 2);
    part_dos(input, 1000000);
}

#[cfg(test)]
mod tests {
    use super::{part_uno, part_dos};
    #[test]
    fn test() {
        // Reads input and runs solve functions for parts
        let input: &str = include_str!("../test.txt").trim();
        assert_eq!(part_uno(input, 2), 374);
        assert_eq!(part_dos(input, 100), 8410);
    }
}
