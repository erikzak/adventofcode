//! Advent of Code 2023 day 2
use std::{fs, collections::HashMap};


fn read_file(file_path: &str) -> String {
    // Returns file contents as string
    fs::read_to_string(file_path)
        .expect("wtf")
        .trim().to_string()
}


fn generate_hashmaps(input: &str) -> (HashMap<[u8; 2], u32>, HashMap<[u8; 2], char>) {
    let mut numbers: HashMap<[u8; 2], u32> = HashMap::new();
    let mut symbols: HashMap<[u8; 2], char> = HashMap::new();
    // Generate hashmaps of schematic numbers and symbols
    for (y, line) in input.lines().rev().enumerate() {
        let mut x: u8 = 0;
        // Grab symbols and numbers per line
        let mut value: String = String::from("");
        for c in line.chars() {
            if c.is_numeric() {
                value.push(c);
            } else {
                if c != '.' {
                    symbols.insert([x, y as u8], c);
                }
                if !value.is_empty() {
                    let coords = [x - value.chars().count() as u8, y as u8];
                    numbers.insert(coords, value.parse::<u32>().unwrap());
                    value = String::from("");
                }
            }
            x += 1;
        }
        if !value.is_empty() {
            let coords = [x - value.chars().count() as u8, y as u8];
            numbers.insert(coords, value.parse::<u32>().unwrap());
        }
    }
    return (numbers, symbols);
}


fn touches_symbol(coords: &[u8; 2], number: &u32, symbols: &HashMap<[u8; 2], char>) -> bool {
    let x_min: u8 = if coords[0] == 0 {0} else {coords[0] - 1};
    let x_max: u8 = coords[0] + number.to_string().chars().count() as u8;
    let y_min: u8 = if coords[1] == 0 {0} else {coords[1] - 1};
    let y_max: u8 = coords[1] + 1 as u8;
    for xx in x_min..=x_max {
        for yy in y_min..=y_max {
            if symbols.contains_key(&[xx, yy]) {
                return true;
            }
        }
    }
    false
}


fn touches_number(coords: &[u8; 2], number_coords: &[u8; 2], number: &u32) -> bool {
    let x_min: u8 = if number_coords[0] == 0 {0} else {number_coords[0] - 1};
    let x_max: u8 = number_coords[0] + number.to_string().chars().count() as u8;
    let y_min: u8 = if number_coords[1] == 0 {0} else {number_coords[1] - 1};
    let y_max: u8 = number_coords[1] + 1 as u8;
    for xx in x_min..=x_max {
        for yy in y_min..=y_max {
            if coords[0] == xx && coords[1] == yy {
                return true;
            }
        }
    }
    false
}


fn part_uno(input: &str) {
    // Solves part 1
    let mut answer: u32 = 0;
    let (numbers, symbols) = generate_hashmaps(input);
    // Loop through schematic numbers and check which of them touch a symbol
    for (coords, number) in numbers {
        if touches_symbol(&coords, &number, &symbols) {
            answer += number;
        }
    }
    println!("Part one: {answer}");
}


fn part_dos(input: &str) {
    // Solves part 2
    let mut answer: u32 = 0;
    let (numbers, symbols) = generate_hashmaps(input);
    for (symbol_coords, symbol) in symbols {
        if symbol != '*' {
            continue;
        }
        let mut touches: Vec<u32> = Vec::new();
        for (number_coords, number) in &numbers {
            if touches_number(&symbol_coords, &number_coords, &number) {
                touches.push(*number);
            }
        }
        if touches.len() == 2 {
            answer += touches[0] * touches[1];
        }
    }
    println!("Part two: {answer}");
}


fn main() {
    // Reads input and runs solve functions for parts
    let input: String = read_file("input.txt");
    part_uno(&input);
    part_dos(&input);
}
