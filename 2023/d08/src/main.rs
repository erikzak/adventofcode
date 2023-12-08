//! Advent of Code 2023 day 7
use std::collections::HashMap;

#[derive(Debug)]
struct Node {
    left: String,
    right: String,
}

impl Node {
    fn move_to(&self, direction: &char) -> &str {
        if direction == &'L' {
            return &self.left;
        }
        &self.right
    }
}


fn parse_input(input: &str) -> (Vec<char>, HashMap<String, Node>) {
    let newline_fix: String = input.replace("\r\n", "\n").to_string();
    let split: Vec<&str> = newline_fix.split("\n\n").collect();
    let instructions: Vec<char> = split[0].chars().collect();
    let node_map: HashMap<String, Node> = generate_node_map(split[1]);
    return (instructions, node_map);
}


fn generate_node_map(input: &str) -> HashMap<String, Node> {
    let mut node_map: HashMap<String, Node> = HashMap::new();
    for line in input.lines() {
        let split: Vec<&str> = line.split_whitespace().collect();
        let name: String = split[0].to_string();
        let left: String = split[2][1..split[2].len()-1].to_string();
        let right: String = split[3][..split[3].len()-1].to_string();
        node_map.insert(name, Node{left, right});
    }
    node_map
}


fn lcm_of_vector(numbers: &[u64]) -> Option<u64> {
    if numbers.is_empty() {
        None
    } else {
        let result = numbers
            .iter()
            .fold(numbers[0], |acc, &x| lcm(acc, x));
        Some(result)
    }
}


fn lcm(a: u64, b: u64) -> u64 {
    if a == 0 || b == 0 {
        return 0;
    } else {
        return a / gcd(a, b) * b;
    }
}


fn gcd(mut a: u64, mut b: u64) -> u64 {
    while b != 0 {
        let temp = b;
        b = a % b;
        a = temp;
    }

    a
}


fn part_uno(_input: &str) -> u64 {
    // Solves part 1
    let mut steps: u64 = 0;
    let (instructions, node_map) = parse_input(_input);
    let mut location: String = String::from("AAA");
    loop {
        for direction in &instructions {
            let node: &Node = &node_map[&location];
            location = node.move_to(&direction).to_string();
            steps += 1;
            if location == "ZZZ" { break; }
        }
        if location == "ZZZ" { break; }
    }
    println!("Part one: {steps}");
    steps
}


fn part_dos(_input: &str) -> u64 {
    // Solves part 2
    let (instructions, node_map) = parse_input(_input);
    let mut locations: Vec<String> = Vec::new();
    let mut steps: Vec<u64> = Vec::new();
    for loc in node_map.keys() {
        if loc.ends_with("A") {
            locations.push(loc.to_string());
        }
    }
    for location in locations.iter_mut() {
        let mut ghost_steps: u64 = 0;
        loop {
            for direction in &instructions {
                let node: &Node = &node_map[location];
                *location = node.move_to(&direction).to_string();
                ghost_steps += 1;
                if location.ends_with("Z") { break; }
            }
            if location.ends_with("Z") { break; }
        }
        steps.push(ghost_steps);
    }
    let lcm: u64 = lcm_of_vector(&steps).unwrap();
    println!("Part two: {lcm}");
    lcm
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
        assert_eq!(part_uno(input), 6);
        assert_eq!(part_dos(input), 6);
    }
}
