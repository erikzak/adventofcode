//! Advent of Code 2023 day 1
use std::collections::HashMap;
use std::fs;


fn read_file(file_path: &str) -> String {
    // Returns file contents as string
    fs::read_to_string(file_path)
        .expect("wat")
        .trim().to_string()
}


fn get_first_number_char_in_string(line: &str, reverse: bool) -> char {
    // Returns first character in string that is numeric as a char,
    // optionally with boolean flag to start looking from end of string.
    let text: String;
    if reverse {
        text = line.chars().rev().collect();
    } else {
        text = line.to_string();
    }
    for c in text.chars() {
        if c.is_numeric() {
            return c;
        }
    }
    panic!("No numeric characters found");
}


fn part_uno(input: &str) {
    let mut sum: i32 = 0;
    for line in input.lines() {
        let first: char = get_first_number_char_in_string(line, false);
        let last: char = get_first_number_char_in_string(line, true);
        let number: String = format!("{first}{last}");
        sum += number.parse::<i32>().unwrap();
    }
    println!("Part one: {sum}");
}


fn part_dos(input: &str) {
    let numbers: HashMap<&str, &str> = HashMap::from([
        ("one", "o1e"),
        ("two", "t2o"),
        ("three", "t3e"),
        ("four", "f4r"),
        ("five", "f5e"),
        ("six", "s6x"),
        ("seven", "s7n"),
        ("eight", "e8t"),
        ("nine", "n9e"),
    ]);
    let mut sum: i32 = 0;
    for line in input.lines() {
        let mut line: String = line.to_string();
        for (word, number) in &numbers {
            line = line.replace(word, number);
        }
        let first: char = get_first_number_char_in_string(line.as_str(), false);
        let last: char = get_first_number_char_in_string(line.as_str(), true);
        let number: String = format!("{first}{last}");
        sum += number.parse::<i32>().unwrap();
    }
    println!("Part two: {sum}");
}


fn main() {
    let file_path: &str = "input.txt";
    let input: String = read_file(file_path);
    part_uno(&input);
    part_dos(&input);
}
