//! Advent of Code 2023 day 5
use std::iter::zip;


struct Race {
    time: u64,
    record: u64,
}


impl Race {
    fn ways(self) -> u64 {
        // Returns ways of winning the race
        let mut ways: u64 = 0;
        for ms in 1..self.time-1 {
            if self.winning(ms) {
                ways += 1;
            }
        }
        ways
    }

    fn winning(&self, milliseconds: u64) -> bool {
        // Returns boolean for if holding the button for given milliseconds wins the race
        let speed: u64 = milliseconds;
        let moving_time: u64 = self.time - milliseconds;
        let distance: u64 = moving_time * speed;
        distance > self.record
    }
}


fn part_uno(_input: &str) -> u64  {
    // Solves part 1
    let mut answer: u64 = 0;
    // Parse input
    let mut races: Vec<Race> = Vec::new();
    let lines: Vec<&str> = _input.lines().collect();
    let times: Vec<u64> = lines[0].split(":").collect::<Vec<&str>>()[1]
        .split_whitespace()
        .map(|t| { t.parse().unwrap() })
        .collect();
    let records: Vec<u64> = lines[1].split(":").collect::<Vec<&str>>()[1]
        .split_whitespace()
        .map(|t| { t.parse().unwrap() })
        .collect();
    for (time, record) in zip(times, records) {
        races.push(Race{time, record});
    }
    for race in races {
        let ways: u64 = race.ways();
        answer = if answer == 0 { ways } else { answer * ways };
    }
    println!("Part one: {answer}");
    answer
}


fn part_dos(_input: &str) -> u64 {
    // Solves part 2
    let lines: Vec<&str> = _input.lines().collect();
    let time: u64 = lines[0].split(":").collect::<Vec<&str>>()[1]
        .chars().filter(|c| !c.is_whitespace()).collect::<String>()
        .parse().unwrap();
    let record: u64 = lines[1].split(":").collect::<Vec<&str>>()[1]
        .chars().filter(|c| !c.is_whitespace()).collect::<String>()
        .parse().unwrap();
    let race: Race = Race{time, record};
    let ways: u64 = race.ways();
    println!("Part two: {ways}");
    ways
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
        assert_eq!(part_uno(input), 288);
        assert_eq!(part_dos(input), 71503);
    }
}
