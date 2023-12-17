//! Point class for ASCII 2D construct coordinates


#[derive(Debug, Eq, Hash, PartialEq, Copy, Clone)]
pub struct Point {
    pub x: i64,
    pub y: i64,
}


impl std::fmt::Display for Point {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(f, "({}, {})", self.x, self.y)
    }
}


impl Point {
    pub fn from(coords: [i64; 2]) -> Point {
        Point{ x: coords[0], y: coords[1] }
    }
}
