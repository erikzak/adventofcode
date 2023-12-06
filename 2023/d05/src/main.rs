//! Advent of Code 2023 day 5
use std::collections::HashMap;


fn parse_input(input: &str, part1: bool) -> (Vec<Seed>, HashMap<String, Map>) {
    // Parses input and returns list of Seeds instances, along with hashmaps of
    // map type keys pointing to Map values
    let mut seeds: Vec<Seed> = Vec::new();
    let mut maps: HashMap<String, Map> = HashMap::new();
    let binding: String = input.replace("\r\n", "\n");
    let sections: Vec<&str> = binding.split("\n\n").map(|s: &str| s.trim()).collect();
    let seed_numbers: &[&str] = &sections[0].split_whitespace().collect::<Vec<_>>()[1..];
    if part1 {
        for seed_number in seed_numbers {
            seeds.push(Seed::new(seed_number.parse().unwrap(), 0));
        }
    } else {
        for (i, seed_number) in seed_numbers.iter().enumerate() {
            if i % 2 != 0 { continue; }
            let number: u64 = seed_number.parse().unwrap();
            let range: u64 = seed_numbers[i+1].parse().unwrap();
            seeds.push(Seed::new(number, range));
        }
    }
    for section in &sections[1..] {
        let map: Map = Map::new(section);
        maps.insert(map.map_type.clone(), map);
    }
    return (seeds, maps);
}


struct Map {
    // Keeps track of mapped sources, destinations and ranges
    map_type: String,
    sources: Vec<u64>,
    destinations: Vec<u64>,
    ranges: Vec<u64>,
}

impl Map {
    fn new(input: &str) -> Map {
        let lines: Vec<&str> = input.lines().collect();
        let map_type: String = lines[0].split_whitespace().collect::<Vec<&str>>()[0].to_string();
        let mut sources: Vec<u64> = Vec::new();
        let mut destinations: Vec<u64> = Vec::new();
        let mut ranges: Vec<u64> = Vec::new();
        for line in &lines[1..] {
            let numbers: Vec<&str> = line.split_whitespace().collect();
            destinations.push(numbers[0].parse().unwrap());
            sources.push(numbers[1].parse().unwrap());
            ranges.push(numbers[2].parse().unwrap());
        }
        Map{ map_type, sources, destinations, ranges }
    }

    fn forward(&self, source: u64) -> u64 {
        for (i, d0) in self.destinations.iter().enumerate() {
            let s0: u64 = self.sources[i];
            let r: u64 = self.ranges[i];
            if source >= s0 && source < s0 + r {
                return d0 + source - s0
            }
        }
        source
    }

    fn reverse(&self, destination: u64) -> u64 {
        for (i, s0) in self.sources.iter().enumerate() {
            let d0: u64 = self.destinations[i];
            let r: u64 = self.ranges[i];
            if destination >= d0 && destination < d0 + r {
                return s0 + destination - d0
            }
        }
        destination
    }
}

struct Seed {
    // Keeps track of seed properties 
    number: u64,
    range: u64,
}


impl Seed {
    // Implements Seed methods to map seeds to locations
    fn new(number: u64, range: u64) -> Seed {
        Seed{ number, range }
    }
}

fn map_to_location(seed: u64, maps: &HashMap<String, Map>) -> u64 {
    let soil: u64 = maps["seed-to-soil"].forward(seed);
    let fertilizer: u64 = maps["soil-to-fertilizer"].forward(soil);
    let water: u64 = maps["fertilizer-to-water"].forward(fertilizer);
    let light: u64 = maps["water-to-light"].forward(water);
    let temperature: u64 = maps["light-to-temperature"].forward(light);
    let humidity: u64 = maps["temperature-to-humidity"].forward(temperature);
    maps["humidity-to-location"].forward(humidity)
}


fn map_to_seed(location: u64, maps: &HashMap<String, Map>) -> u64 {
    let humidity: u64 = maps["humidity-to-location"].reverse(location);
    let temperature: u64 = maps["temperature-to-humidity"].reverse(humidity);
    let light: u64 = maps["light-to-temperature"].reverse(temperature);
    let water: u64 = maps["water-to-light"].reverse(light);
    let fertilizer: u64 = maps["fertilizer-to-water"].reverse(water);
    let soil: u64 = maps["soil-to-fertilizer"].reverse(fertilizer);
    maps["seed-to-soil"].reverse(soil)
}


fn part_uno(_input: &str) {
    // Solves part 1
    let mut answer: u64 = 0;
    let (seeds, maps) = parse_input(_input, true);
    for seed in seeds {
        let location = map_to_location(seed.number, &maps);
        if answer == 0 || location < answer {
            answer = location;
        }
    }
    println!("Part one: {answer}");
}


fn part_dos(_input: &str) {
    // Solves part 2
    let mut location = 0;
    let (seeds, maps) = parse_input(_input, false);
    loop {
        let seed_number: u64 = map_to_seed(location, &maps);
        for seed in &seeds {
            if seed_number >= seed.number && seed_number < seed.number + seed.range {
                println!("Part two: {location}");
                return;
            }
        }
        location += 1;
    }
}


fn main() {
    // Reads input and runs solve functions for parts
    let input: &str = include_str!("../input.txt").trim();
    part_uno(input);
    part_dos(input);
}
