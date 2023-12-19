//! Grid class for ASCII 2D constructs
use std::collections::HashMap;

use crate::grid::point::Point;
use crate::grid::extent::Extent;


#[derive(Clone)]
pub struct Grid {
    pub values: HashMap<Point, char>,
    pub extent: Extent,
}


impl Grid {
    #[allow(dead_code)]
    pub fn new(input: &str) -> Grid {
        let mut values: HashMap<Point, char> = HashMap::new();
        let mut min_x: Option<i64> = None;
        let mut max_x: Option<i64> = None;
        let mut min_y: Option<i64> = None;
        let mut max_y: Option<i64> = None;
        for (y, line) in input.lines().rev().enumerate() {
            for (x, value) in line.trim().chars().enumerate() {
                let pnt: Point = Point::from([x as i64, y as i64]);
                values.insert(pnt, value);
                if (x as i64) < min_x.unwrap_or(i64::MAX) { min_x = Some(x as i64); }
                if (x as i64) > max_x.unwrap_or(i64::MIN) { max_x = Some(x as i64); }
            }
            if (y as i64) < min_y.unwrap_or(i64::MAX) { min_y = Some(y as i64); }
            if (y as i64) > max_y.unwrap_or(i64::MIN) { max_y = Some(y as i64); }
        }
        let extent = Extent::from([min_x.unwrap(), max_x.unwrap(), min_y.unwrap(), max_y.unwrap()]);
        Grid { values, extent }
    }

    #[allow(dead_code)]
    pub fn next_point(&self, point: &Point, direction: &char) -> Option<Point> {
        match direction {
            'R' => {
                if point.x + 1 > self.extent.max_x {
                    return None;
                }
                Some(Point::from([point.x+1, point.y]))
            },
            'L' => {
                if point.x < self.extent.min_x {
                    return None;
                }
                Some(Point::from([point.x-1, point.y]))
            },
            'U' => {
                if point.y + 1 > self.extent.max_y {
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
        for y in (self.extent.min_y..=self.extent.max_y).rev() {
            for x in self.extent.min_x..=self.extent.max_x {
                let pnt: Point = Point::from([x as i64, y as i64]);
                chars.push(*self.values.get(&pnt).unwrap_or(&'.'));
            }
            chars.push('\n');
        }
        let string: String = chars.into_iter().collect();
        write!(f, "{}", string.trim())
    }
}
