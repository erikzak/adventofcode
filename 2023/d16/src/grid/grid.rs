//! Grid class for ASCII 2D constructs
use std::collections::HashMap;

use crate::grid::point::Point;


#[derive(Clone)]
pub struct Grid {
    pub values: HashMap<Point, char>,
    pub dim: [u64; 2],
}


impl Grid {
    pub fn new(input: &str) -> Grid {
        let mut values: HashMap<Point, char> = HashMap::new();
        let mut max_x: u64 = 0;
        let mut max_y: u64 = 0;
        for (y, line) in input.lines().rev().enumerate() {
            for (x, value) in line.trim().chars().enumerate() {
                let pnt: Point = Point::from([x as i64, y as i64]);
                values.insert(pnt, value);
                max_x = x as u64;
            }
            max_y = y as u64;
        }
        let dim: [u64; 2] = [max_x + 1, max_y + 1];
        Grid { values, dim }
    }

    pub fn next_point(&self, point: &Point, direction: &char) -> Option<Point> {
        match direction {
            'E' => {
                if point.x + 1 == self.dim[0] as i64 {
                    return None;
                }
                Some(Point::from([point.x+1, point.y]))
            },
            'W' => {
                if point.x == 0 {
                    return None;
                }
                Some(Point::from([point.x-1, point.y]))
            },
            'N' => {
                if point.y + 1 == self.dim[1] as i64 {
                    return None;
                }
                Some(Point::from([point.x, point.y+1]))
            },
            'S' => {
                if point.y == 0 {
                    return None;
                }
                Some(Point::from([point.x, point.y-1]))
            },
            _ => panic!("unhandled direction"),
        }
    }
}


impl std::fmt::Display for Grid {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        let mut chars: Vec<char> = Vec::new();
        for y in (0..self.dim[1]).rev() {
            for x in 0..self.dim[0] {
                let pnt: Point = Point::from([x as i64, y as i64]);
                chars.push(self.values[&pnt]);
            }
            chars.push('\n');
        }
        let string: String = chars.into_iter().collect();
        write!(f, "{}", string.trim())
    }
}
