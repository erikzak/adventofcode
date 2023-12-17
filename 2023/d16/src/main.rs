//! Advent of Code 2023 day 16
use std::collections::HashMap;

mod grid;
use grid::grid::Grid;
use grid::point::Point;


struct Construct {
    grid: Grid,
    beams: Vec<Beam>,
    energized: HashMap<Point, Vec<char>>,
}


impl Construct {
    fn new(input: &str) -> Construct {
        let grid: Grid = Grid::new(input);
        Construct{ grid, beams: Vec::new(), energized: HashMap::new() }
    }

    fn reset(&mut self) {
        self.energized = HashMap::new();
        self.beams = Vec::new();
    }

    fn tick(&mut self) {
        let mut new_beams: Vec<Beam> = Vec::new();
        self.beams.retain_mut(|beam: &mut Beam| {
            let Some(next_point): Option<Point> = self.grid.next_point(&beam.point, &beam.direction)
                else { return false; };
            // Check if another beam has already energized the next point in the same direction
            let already_energized: bool = self.energized.contains_key(&next_point) &&
                self.energized[&next_point].contains(&beam.direction);
            if already_energized { return false; }
            // Track energized squares
            match self.energized.get_mut(&next_point) {
                Some(dirs) => {
                    dirs.push(beam.direction);
                },
                None => {
                    self.energized.insert(next_point, vec![beam.direction]);
                },
            }
            // Process beam
            beam.point = next_point;
            let next_value: char = self.grid.values[&next_point];
            match next_value {
                '.' => (),
                '\\' => match beam.direction {
                    'E' => { beam.direction = 'S'; },
                    'W' => { beam.direction = 'N'; },
                    'N' => { beam.direction = 'W'; },
                    'S' => { beam.direction = 'E'; },
                    _ => panic!("unhandled direction"),
                },
                '/' => match beam.direction {
                    'E' => { beam.direction = 'N'; },
                    'W' => { beam.direction = 'S'; },
                    'N' => { beam.direction = 'E'; },
                    'S' => { beam.direction = 'W'; },
                    _ => panic!("unhandled direction"),
                },
                '-' => {
                    if beam.direction == 'N' || beam.direction == 'S' {
                        beam.direction = 'E';
                        new_beams.push(Beam{ point: beam.point.clone(), direction: 'W' });
                    }
                },
                '|' => {
                    if beam.direction == 'W' || beam.direction == 'E' {
                        beam.direction = 'N';
                        new_beams.push(Beam{ point: beam.point.clone(), direction: 'S' });
                    }
                },
                _ => panic!("unhandled grid value"),
            }
            true
        });
        self.beams.extend(new_beams);
    }
}


impl std::fmt::Display for Construct {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        let mut chars: Vec<char> = Vec::new();
        for y in (0..self.grid.dim[1]).rev() {
            for x in 0..self.grid.dim[0] {
                let pnt: Point = Point::from([x as i64, y as i64]);
                let v: char;
                if self.energized.contains_key(&pnt) {
                    v = '#';
                } else {
                    v = self.grid.values[&pnt];
                }
                chars.push(v);
            }
            chars.push('\n');
        }
        let string: String = chars.into_iter().collect();
        write!(f, "{}", string.trim())
    }
}

#[derive(Clone, Debug)]
struct Beam {
    point: Point,
    direction: char,
}


fn part_uno(_input: &str) -> u64 {
    // Solves part 1
    let mut construct: Construct = Construct::new(_input);
    let starting_point: Point = Point::from([-1, construct.grid.dim[1] as i64 - 1]);
    let beam: Beam = Beam{ point: starting_point, direction: 'E' };
    construct.beams.push(beam);
    while construct.beams.len() > 0 {
        construct.tick();
    }
    let answer: u64 = construct.energized.len() as u64;
    println!("{answer}");
    answer
}


fn part_dos(_input: &str) -> u64 {
    // Solves part 2
    let mut answer: u64 = 0;
    let mut construct: Construct = Construct::new(_input);
    for x in 0..construct.grid.dim[0] {
        let starting_point: Point = Point::from([x as i64, -1]);
        let beam: Beam = Beam{ point: starting_point, direction: 'N' };
        construct.beams.push(beam);
        while construct.beams.len() > 0 {
            construct.tick();
        }
        let energized: u64 = construct.energized.len() as u64;
        if energized > answer {
            answer = energized;
        }
        construct.reset();
        let starting_point: Point = Point::from([x as i64, construct.grid.dim[1] as i64]);
        let beam: Beam = Beam{ point: starting_point, direction: 'S' };
        construct.beams.push(beam);
        while construct.beams.len() > 0 {
            construct.tick();
        }
        let energized: u64 = construct.energized.len() as u64;
        if energized > answer {
            answer = energized;
        }
        construct.reset();
    }
    for y in 0..construct.grid.dim[1] {
        let starting_point: Point = Point::from([-1, y as i64]);
        let beam: Beam = Beam{ point: starting_point, direction: 'E' };
        construct.beams.push(beam);
        while construct.beams.len() > 0 {
            construct.tick();
        }
        let energized: u64 = construct.energized.len() as u64;
        if energized > answer {
            answer = energized;
        }
        construct.reset();
        let starting_point: Point = Point::from([construct.grid.dim[0] as i64, y as i64]);
        let beam: Beam = Beam{ point: starting_point, direction: 'W' };
        construct.beams.push(beam);
        while construct.beams.len() > 0 {
            construct.tick();
        }
        let energized: u64 = construct.energized.len() as u64;
        if energized > answer {
            answer = energized;
        }
        construct.reset();
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
        assert_eq!(part_uno(input), 46);
        assert_eq!(part_dos(input), 51);
    }
}
