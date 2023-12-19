//! Extent class for ASCII 2D construct dimensions

use crate::grid::point::Point;

#[derive(Debug, Eq, Hash, PartialEq, Copy, Clone)]
pub struct Extent {
    pub min_x: i64,
    pub max_x: i64,
    pub min_y: i64,
    pub max_y: i64,
}

impl std::fmt::Display for Extent {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(
            f, "(min_x: {}, max_x: {}, min_y: {}, max_y: {})",
            self.min_x, self.max_x, self.min_y, self.max_y,
        )
    }
}

impl Extent {
    #[allow(dead_code)]
    pub fn new() -> Extent {
        Extent { min_x: 0, max_x: 0, min_y: 0, max_y: 0 }
    }

    pub fn from(dims: [i64; 4]) -> Extent {
        Extent{ min_x: dims[0], max_x: dims[1], min_y: dims[2], max_y: dims[3] }
    }

    #[allow(dead_code)]
    pub fn update(&mut self, point: Point) {
        if point.x < self.min_x { self.min_x = point.x; }
        if point.x > self.max_x { self.max_x = point.x; }
        if point.y < self.min_y { self.min_y = point.y; }
        if point.y > self.max_y { self.max_y = point.y; }
    }
}
