//! Advent of Code 2023 day 7
use core::cmp::Ordering;
use std::collections::HashMap;
use std::iter::zip;


struct Hand {
    cards: String,
    card_value: Vec<u8>,
    bid: u32,
    value: u32,
    rank: u32,
}


impl Hand {
    fn new(cards: &str, bid: u32, jokers: bool) -> Hand {
        // Inits hand and identifies value
        let mut card_value: Vec<u8> = Vec::new();
        for c in cards.chars() {
            match c {
                'A' => card_value.push(14),
                'K' => card_value.push(13),
                'Q' => card_value.push(12),
                'J' => card_value.push(if jokers { 1 } else { 11 }),
                'T' => card_value.push(10),
                _ => card_value.push(c.to_string().parse().unwrap()),
            }
        }
        let mut hand: Hand = Hand{cards: cards.to_string(), card_value, bid, value: 0, rank: 0};
        hand.value = if !jokers { hand.calculate_p1_value() } else { hand.calculate_p2_value() };
        hand
    }

    fn calculate_p1_value(&mut self) -> u32 {
        // Calculates value based on hand contents:
        // Five of a kind = 6
        // Four of a kind = 5
        // Full house = 4
        // Three of a kind = 3
        // Two pair = 2
        // One pair = 1
        // High card = 0
        let mut counts: HashMap<char, u8> = HashMap::new();
        for c in self.cards.chars() {
            *counts.entry(c).or_insert(0) += 1;
        }
        let mut sorted_counts: Vec<&u8> = counts.values().collect();
        sorted_counts.sort();
        sorted_counts.reverse();
        // *Cries in no match on vecs*
        if *sorted_counts[0] == 5 {
            return 6;
        } else if *sorted_counts[0] == 4 {
            return 5;
        } else if *sorted_counts[0] == 3 && *sorted_counts[1] == 2 {
            return 4;
        } else if *sorted_counts[0] == 3 {
            return 3;
        } else if *sorted_counts[0] == 2 && *sorted_counts[1] == 2 {
            return 2;
        } else if *sorted_counts[0] == 2 {
            return 1;
        }
        return 0;
    }

    fn calculate_p2_value(&mut self) -> u32 {
        // Calculates value based on hand contents, with jokers.
        let mut counts: HashMap<char, u8> = HashMap::new();
        let mut jokers: u8 = 0;
        counts.insert('J', 0);
        for c in self.cards.chars() {
            if c == 'J' {
                jokers += 1;
                continue;
            }
            *counts.entry(c).or_insert(0) += 1;
        }
        let mut sorted_counts: Vec<&u8> = counts.values().collect();
        sorted_counts.sort();
        sorted_counts.reverse();
        if *sorted_counts[0] + jokers == 5 {
            return 6;
        } else if *sorted_counts[0] + jokers == 4 {
            return 5;
        } else if *sorted_counts[0] + jokers == 3 && *sorted_counts[1] == 2 {
            return 4;
        } else if *sorted_counts[0] + jokers == 3 {
            return 3;
        } else if *sorted_counts[0] + jokers == 2 && *sorted_counts[1] == 2 {
            return 2;
        } else if *sorted_counts[0] + jokers == 2 {
            return 1;
        }
        return 0;
    }

    fn partial_cmp(&self, other: &Hand) -> Option<Ordering> {
        // Compare value, then highest card
        if self.value != other.value { return self.value.partial_cmp(&other.value); }
        for (a, b) in zip(self.card_value.iter(), other.card_value.iter()) {
            if a != b { return a.partial_cmp(&b); }
        }
        Some(Ordering::Equal)
    }
}


fn part_uno(_input: &str) -> u32 {
    // Solves part 1
    let answer: u32;
    let mut hands: Vec<Hand> = Vec::new();
    // Generate list of hands
    for line in _input.lines() {
        let split: Vec<&str> = line.split_whitespace().collect();
        let hand: Hand = Hand::new(split[0], split[1].parse().unwrap(), false);
        hands.push(hand);
    }
    // Sort hands by value and assign rank
    hands.sort_by(|a, b| a.partial_cmp(b).unwrap());
    for (rank, hand) in hands.iter_mut().enumerate() {
        hand.rank = rank as u32 + 1;
    }
    // Calculate answer
    answer = hands.iter().fold(0, |acc: u32, hand: &Hand| { acc + hand.bid * hand.rank });
    println!("Part one: {answer}");
    answer
}


fn part_dos(_input: &str) -> u32 {
    // Solves part 2
    let answer: u32;
    let mut hands: Vec<Hand> = Vec::new();
    // Generate list of hands
    for line in _input.lines() {
        let split: Vec<&str> = line.split_whitespace().collect();
        let hand: Hand = Hand::new(split[0], split[1].parse().unwrap(), true);
        hands.push(hand);
    }
    // Sort hands by value and assign rank
    hands.sort_by(|a, b| a.partial_cmp(b).unwrap());
    for (rank, hand) in hands.iter_mut().enumerate() {
        hand.rank = rank as u32 + 1;
    }
    // Calculate answer
    answer = hands.iter().fold(0, |acc: u32, hand: &Hand| { acc + hand.bid * hand.rank });
    println!("Part two: {answer}");
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
        assert_eq!(part_uno(input), 6440);
        assert_eq!(part_dos(input), 5905);
    }
}
