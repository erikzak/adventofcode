//! Advent of Code 2023 day 4
use std::collections::HashSet;


struct Card {
    // Keeps track of card properties and implements methods for parsing numbers from input
    input: String,
    id: u8,
    numbers: Vec<u8>,
    winning_numbers: HashSet<u8>,
    matches: u8,
    points: u32,
}


impl Card {
    fn new(input: &str) -> Card {
        let mut card: Card = Card{
            input: input.to_string(), id: 0,
            numbers: Vec::new(), winning_numbers: HashSet::new(),
            matches: 0, points: 0
        };
        card.parse_id();
        card.parse_card_numbers();
        card.parse_winning_numbers();
        card.check_matching_numbers();
        card.calculate_points();
        card
    }

    fn parse_id(&mut self) {
        // Parses card number (id) from input
        self.id = self.input.split(": ").collect::<Vec<&str>>()[0]
            .split_whitespace().collect::<Vec<&str>>()[1]
            .parse::<u8>().unwrap();
    }

    fn parse_card_numbers(&mut self) {
        // Parses card numbers from input as vector
        let numbers_string: &str = self.input.split(" | ").collect::<Vec<&str>>()[1];
        self.numbers = parse_numbers(numbers_string);
    }

    fn parse_winning_numbers(&mut self) {
        // Parses winning card numbers from input and stores it in hashset
        let winning_numbers_string: &str = self.input
            .split(" | ").collect::<Vec<&str>>()[0]
            .split(": ").collect::<Vec<&str>>()[1];
        for number in parse_numbers(winning_numbers_string) {
            self.winning_numbers.insert(number);
        }
    }

    fn check_matching_numbers(&mut self) {
        // Checks how many card numbers match winning numbers
        self.matches = 0;
        for number in &self.numbers {
            if self.winning_numbers.contains(&number) {
                self.matches += 1;
            }
        }
    }

    fn calculate_points(&mut self) {
        // Calculates points based on matching numbers
        self.points = 0;
        if self.matches > 0 {
            self.points += 2_u32.pow(self.matches as u32 - 1);
        }
    }
}


fn parse_numbers(numbers_string: &str) -> Vec<u8> {
    // Returns a vector of integers parsed from string of numbers split by space
    let mut numbers: Vec<u8> = Vec::new();
    for number in numbers_string.split_whitespace().collect::<Vec<&str>>() {
        numbers.push(number.parse().unwrap());
    }
    numbers
}


fn part_uno(_input: &str) {
    // Solves part 1
    let mut answer: u32 = 0;
    for card in _input.lines() {
        let card: Card = Card::new(card);
        answer += card.points;
    }
    println!("Part one: {answer}");
}


fn part_dos(_input: &str) {
    // Solves part 2
    let card_lines: Vec<&str> = _input.lines().collect();
    let mut cards: Vec<u32> = vec![0; card_lines.len()];
    for card in card_lines {
        let card: Card = Card::new(card);
        let idx: usize = card.id as usize - 1;
        cards[idx] += 1;
        for won_idx in idx + 1..idx + 1 + card.matches as usize {
            cards[won_idx] += cards[idx];
        }
    }
    let answer: u32 = cards.iter().sum();
    println!("Part two: {answer}");
}


fn main() {
    // Reads input and runs solve functions for parts
    let input: &str = include_str!("../input.txt").trim();
    part_uno(input);
    part_dos(input);
}
