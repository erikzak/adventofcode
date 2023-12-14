//! Advent of Code 2023 day 12
use memoize::memoize;


struct Record {
    rows: Vec<Row>,
}


impl Record {
    fn new(input: &str) -> Record {
        let mut rows: Vec<Row> = Vec::new();
        for line in input.lines() {
            let split: Vec<&str> = line.split(" ").collect();
            let springs: String = split[0].to_string();
            let groups_part: Vec<&str> = split[1].split(",").collect::<Vec<&str>>();
            let groups: Vec<u32> = groups_part.iter().map(|s| s.parse().unwrap()).collect();
            rows.push(Row{ springs, groups });
        }
        Record{ rows }
    }

    fn unfold(&mut self) {
        for row in self.rows.iter_mut() {
            let mut springs: Vec<&str> = Vec::new();
            let mut groups: Vec<u32> = Vec::new();
            for _ in 0..5 {
                springs.push(&row.springs);
                groups.extend(row.groups.clone());
            }
            row.springs = springs.join("?");
            row.groups = groups;
        }
    }
}


#[derive(Debug)]
struct Row {
    springs: String,
    groups: Vec<u32>,
}


#[memoize]
fn sum_arrangements(springs: String, groups: Vec<u32>) -> u64 {
    // Did we run out of groups? We might still be valid
    if groups.len() == 0 {
        // Make sure there aren't any more damaged springs, if so, we're valid
        if !springs.contains("#") {
            // This will return true even if row is empty, which is valid
            return 1;
        }
        // More damaged springs that we can't fit
        return 0;
    }

    // There are more groups, but no more springs
    if springs.len() == 0 {
        // We can't fit, exit
        return 0;
    }

    // Look at the next element in each row and group
    let next_character: char = springs.chars().collect::<Vec<char>>()[0];
    let next_group: usize = groups[0] as usize;

    let sum: u64;
    if next_character == '#' {
        // Test pound logic
        sum = pound(springs.clone(), groups.clone(), next_group.clone())
    } else if next_character == '.' {
        // Test dot logic
        sum = dot(springs.clone(), groups.clone())
    } else if next_character == '?' {
        // This character could be either character, so we'll explore both
        // possibilities
        sum = dot(springs.clone(), groups.clone()) + pound(springs.clone(), groups.clone(), next_group.clone())
    } else {
        panic!("unhandled character {next_character}");
    }
    // println!("{springs}, {:?}, {sum}", groups);
    sum
}


#[memoize]
// Logic that treats the first character as pound
fn pound(springs: String, groups: Vec<u32>, next_group: usize) -> u64 {
    // If the first is a pound, then the first n characters must be
    // able to be treated as a pound, where n is the first group number
    let this_group: String = springs[..if next_group > springs.len() { springs.len() } else { next_group }].to_string();
    let this_group: String = this_group.replace("?", "#");

    // If the next group can't fit all the damaged springs, then abort
    if this_group != "#".repeat(next_group) {
        return 0;
    }

    // If the rest of the record is just the last group, then we're
    // done and there's only one possibility
    if springs.len() == next_group {
        // Make sure this is the last group
        if groups.len() == 1 {
            // We are valid
            return 1;
        }
        // There's more groups, we can't make it work
        return 0;
    }

    // Make sure the character that follows this group can be a seperator
    let test: char = springs.chars().collect::<Vec<char>>()[next_group];
    if test == '?' || test == '.' {
        // It can be seperator, so skip it and reduce to the next group
        return sum_arrangements(springs[next_group+1..].to_string(), groups[1..].to_vec());
    }

    // Can't be handled, there are no possibilites
    0
}


#[memoize]
// Logic that treats the first character as a dot
fn dot(springs: String, groups: Vec<u32>) -> u64 {
    // We just skip over the dot looking for the next pound
    sum_arrangements(springs[1..].to_string(), groups.to_vec())
}


fn part_uno(_input: &str) -> u64 {
    // Solves part 1
    let mut answer: u64 = 0;
    let record = Record::new(_input);
    for row in record.rows {
        answer += sum_arrangements(row.springs, row.groups);
    }
    println!("{answer}");
    answer
}


fn part_dos(_input: &str) -> u64 {
    // Solves part 2
    let mut answer: u64 = 0;
    let mut record = Record::new(_input);
    record.unfold();
    for row in record.rows {
        answer += sum_arrangements(row.springs, row.groups);
    }
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
        assert_eq!(part_uno(input), 21);
        assert_eq!(part_dos(input), 525152);
    }
}
