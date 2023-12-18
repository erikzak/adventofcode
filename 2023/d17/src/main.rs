//! Advent of Code 2023 day 17
use std::cmp::Ordering;
use std::collections::{BinaryHeap, HashMap};

mod grid;
use grid::grid::Grid;
use grid::point::Point;

struct City {
    grid: Grid,
}

impl City {
    fn reconstruct_path(&self, came_from: &HashMap<StateKey, StateKey>, end: StateKey) -> Vec<Point> {
        let mut current: StateKey = end;
        let mut path: Vec<Point> = vec![current.point];
        while came_from.contains_key(&current) {
            current = came_from[&current];
            path.push(current.point);
        }
        path.reverse();
        path
    }

    fn calculate_cost(&self, path: &Vec<Point>) -> u32 {
        let mut cost: u32 = 0;
        for pnt in &path[1..] {
            cost += self.grid.values[&pnt];
        }
        cost
    }

    #[allow(dead_code)]
    fn debug_path(&self, path: &Vec<Point>) {
        let mut chars: Vec<char> = Vec::new();
        for y in (0..self.grid.dim[1]).rev() {
            for x in 0..self.grid.dim[0] {
                let pnt: Point = Point::from([x as i64, y as i64]);
                let v: char = if path.contains(&pnt) { '*' } else {
                    char::from_digit(self.grid.values[&pnt], 10).unwrap()
                };
                chars.push(v);
            }
            chars.push('\n');
        }
        let string: String = chars.into_iter().collect();
        println!("{}\n", string.trim())
    }

    fn get_neighbors(&self, point: &Point) -> Vec<(Point, char)> {
        let mut neighbors: Vec<(Point, char)> = Vec::new();
        let offsets: Vec<(i64, i64, char)> = vec![(0, 1, 'N'), (1, 0, 'E'), (0, -1, 'S'), (-1, 0, 'W')];
        for (dx, dy, direction) in offsets {
            let x: i64 = point.x + dx;
            let y: i64 = point.y + dy;
            if x < 0 || y < 0 || x >= self.grid.dim[0] as i64 || y >= self.grid.dim[1] as i64 {
                continue;
            }
            neighbors.push((Point { x, y }, direction));
        }
        neighbors
    }

    fn get_cart_neighbors(
            &self, node: &State, min_blocks: u32, max_blocks: u32
    ) -> Vec<(Point, char)> {
        let mut neighbors: Vec<(Point, char)> = self.get_neighbors(&node.point);
        neighbors.retain(|neighbor| {
            let direction: char = neighbor.1;
            // Has to move <min_blocks> before turning
            if node.direction != '-' && direction != node.direction && node.steps < min_blocks { return false; }
            // Can't move more than <max_blocks> consecutive blocks in the same direction
            if direction == node.direction && node.steps == max_blocks { return false; }

            // Can't reverse
            match direction {
                'N' => { if node.direction == 'S' { return false; }},
                'S' => { if node.direction == 'N' { return false; }},
                'E' => { if node.direction == 'W' { return false; }},
                'W' => { if node.direction == 'E' { return false; }},
                _ => panic!("unhandled direction"),
            }

            // Valid neighbor
            true
        });
        neighbors
    }

    fn dijkstra(&self, start: Point, goal: Point, min_blocks: u32, max_blocks: u32) -> Vec<Point> {
        // Dijkstra's shortest path algorithm
        let mut came_from: HashMap<StateKey, StateKey> = HashMap::new();
        let mut heap: BinaryHeap<State> = BinaryHeap::new();
        let first = State{ point: start, cost: 0, direction: '-', steps: 0 };
        let mut costs: HashMap<StateKey, u32> = HashMap::new();
        costs.insert(StateKey::from(first), 0);
        heap.push(first);

        while let Some(current) = heap.pop() {
            let current_key = StateKey::from(current);
            if current.point == goal && current.steps >= min_blocks {
                return self.reconstruct_path(&came_from, current_key);
            }
            if current.cost > costs[&StateKey::from(current)] { continue; }
            for (neighbor, direction) in self.get_cart_neighbors(&current, min_blocks, max_blocks) {
                let next = State {
                    point: neighbor,
                    cost: current.cost + self.grid.values[&neighbor],
                    direction,
                    steps: if direction == current.direction { current.steps + 1 } else { 1 },
                };
                let next_key = StateKey::from(next);
                if next.cost < *costs.get(&next_key).unwrap_or(&u32::MAX) {
                    came_from.insert(next_key, current_key);
                    costs.insert(next_key, next.cost);
                    heap.push(next);
                }
            }
        }
        panic!("no path!");
    }
}

#[derive(Copy, Clone, Debug, Hash, PartialEq, Eq)]
struct State {
    point: Point,
    cost: u32,
    direction: char,
    steps: u32,
}

#[derive(Copy, Clone, Hash, PartialEq, Eq)]
struct StateKey {
    point: Point,
    direction: char,
    steps: u32,
}

impl StateKey {
    fn from(state: State) -> StateKey {
        StateKey{ point: state.point, direction: state.direction, steps: state.steps }
    }
}

impl Ord for State {
    fn cmp(&self, other: &Self) -> Ordering {
        other.cost.cmp(&self.cost)
            .then_with(|| other.direction.cmp(&self.direction))
            .then_with(|| other.steps.cmp(&self.steps))
            .then_with(|| other.point.x.cmp(&self.point.x))
            .then_with(|| other.point.y.cmp(&self.point.y))
    }
}

impl PartialOrd for State {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

fn part_uno(_input: &str) -> u32 {
    // Solves part 1
    let city = City { grid: Grid::new(_input) };
    let start = Point { x: 0, y: city.grid.dim[1] as i64 - 1 };
    let goal = Point { x: city.grid.dim[0] as i64 - 1, y: 0 };
    let path: Vec<Point> = city.dijkstra(start, goal, 1, 3);
    let answer: u32 = city.calculate_cost(&path);
    println!("{answer}");
    answer
}

fn part_dos(_input: &str) -> u32 {
    // Solves part 2
    let city = City { grid: Grid::new(_input) };
    let start = Point { x: 0, y: city.grid.dim[1] as i64 - 1 };
    let goal = Point { x: city.grid.dim[0] as i64 - 1, y: 0 };
    let path: Vec<Point> = city.dijkstra(start, goal, 4, 10);
    let answer: u32 = city.calculate_cost(&path);
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
        assert_eq!(part_uno(input), 102);
        assert_eq!(part_dos(input), 94);
        let input_2b: &str = include_str!("../test2b.txt").trim();
        assert_eq!(part_dos(input_2b), 71);
    }
}
