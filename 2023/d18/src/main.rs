//! Advent of Code 2023 day 18
mod grid;
use grid::point::Point;


struct Dig {
    trench: Vec<Point>,
}


impl Dig {
    fn new(input: &str, use_color: bool) -> Dig {
        let trench: Vec<Point> = if use_color { Dig::dig_part2(input) } else { Dig::dig_part1(input) };
        Dig { trench }
    }

    fn dig_part1(input: &str) -> Vec<Point> {
        let mut position = Point::from([0, 0]);
        let mut trench: Vec<Point> = vec![position];
        for line in input.lines() {
            let split: Vec<&str> = line.split(" ").collect();
            let direction: char = line.chars().nth(0).unwrap();
            let meters: i64 = split[1].parse().unwrap();
            match direction {
                'R' => { position.x += meters; },
                'D' => { position.y -= meters; },
                'L' => { position.x -= meters; },
                'U' => { position.y += meters; },
                _ => panic!("unhandled direction"),
            }
            trench.push(position);
        }
        trench
    }

    fn dig_part2(input: &str) -> Vec<Point> {
        let mut position = Point::from([0, 0]);
        let mut trench: Vec<Point> = vec![position];
        for line in input.lines() {
            let mut hexadecimal: &str = line.split(" ").collect::<Vec<&str>>()[2];
            hexadecimal = &hexadecimal[2..hexadecimal.len()-1];
            let direction: char = match hexadecimal.chars().last().unwrap() {
                '0' => 'R',
                '1' => 'D',
                '2' => 'L',
                '3' => 'U',
                _ => panic!("unhandled direction code"),
            };
            let distance_hex: &str = &hexadecimal[..hexadecimal.len()-1];
            let meters = i64::from_str_radix(distance_hex, 16).unwrap();
            match direction {
                'R' => { position.x += meters; },
                'D' => { position.y -= meters; },
                'L' => { position.x -= meters; },
                'U' => { position.y += meters; },
                _ => panic!("unhandled direction"),
            }
            trench.push(position);
        }
        trench
    }

    fn calculate_volume(&self) -> u64 {
        // Shoelace formula
        let n: usize = self.trench.len();
        let mut sum: f64 = 0.0;
        for i in 0..n {
            let j = (i + 1) % n;
            sum += (
                self.trench[i].x * self.trench[j].y -
                self.trench[j].x * self.trench[i].y
            ) as f64;
        }
        let mut volume: f64 = sum.abs() * 0.5;
        // Add trench halfsies
        for i in 0..n-1 {
            volume += (
                (self.trench[i+1].x - self.trench[i].x).abs() +
                (self.trench[i+1].y - self.trench[i].y).abs()
            ) as f64 * 0.5;
        }
        volume as u64 + 1
    }
}


fn part_uno(input: &str) -> u64 {
    // Solves part 1
    let dig = Dig::new(input, false);
    let answer: u64 = dig.calculate_volume();
    println!("{answer}");
    answer
}

fn part_dos(input: &str) -> u64 {
    // Solves part 2
    let dig = Dig::new(input, true);
    let answer: u64 = dig.calculate_volume();
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
        assert_eq!(part_uno(input), 62);
        assert_eq!(part_dos(input), 952408144115);
    }
}
