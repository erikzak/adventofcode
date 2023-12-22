//! Advent of Code 2023 day 20
use num::integer::lcm;
use std::collections::{HashMap, VecDeque};
use std::fmt;

#[derive(Clone, Debug)]
struct Pulse {
    high: bool,
    source: String,
    target: String,
}

impl fmt::Display for Pulse {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let pulse: &str = if self.high { "-high->" } else { "-low->" };
        write!(f, "{} {pulse} {}", self.source, self.target)
    }
}

#[derive(Clone, Debug)]
struct FlipFlopModule {
    name: String,
    targets: Vec<String>,
    state: bool,
}

impl FlipFlopModule {
    fn new(name: &str, targets: Vec<&str>,) -> FlipFlopModule {
        let targets: Vec<String> = targets.iter().map(|s| s.to_string()).collect();
        FlipFlopModule{ name: name.to_string(), targets, state: false }
    }
}

impl Module for FlipFlopModule {
    fn receive(&mut self, pulse: &Pulse) -> Vec<Pulse> {
        if pulse.high { return Vec::new(); }
        self.state = if self.state { false } else { true };
        self.send(pulse)
    }

    fn send(&self, _: &Pulse) -> Vec<Pulse> {
        let mut pulses: Vec<Pulse> = Vec::new();
        for target in &self.targets {
            pulses.push(Pulse {
                high: self.state,
                source: self.name.to_string(),
                target: target.to_string()
            });
        }
        pulses
    }

    fn get_targets(&self) -> &Vec<String> {
        &self.targets
    }
}

#[derive(Clone, Debug)]
struct ConjunctionModule {
    name: String,
    targets: Vec<String>,
    last_pulse: HashMap<String, Pulse>,
}

impl ConjunctionModule {
    fn new(name: &str, targets: Vec<&str>, inputs: Vec<String>) -> ConjunctionModule {
        let targets: Vec<String> = targets.iter().map(|s| s.to_string()).collect();
        let mut last_pulse: HashMap<String, Pulse> = HashMap::new();
        for input in inputs {
            let pulse = Pulse {
                high: false,
                source: input.to_string(),
                target: name.to_string(),
            };
            last_pulse.insert(input.to_string(), pulse);
        }
        ConjunctionModule { name: name.to_string(), targets, last_pulse }
    }
}

impl Module for ConjunctionModule {
    fn receive(&mut self, pulse: &Pulse) -> Vec<Pulse> {
        self.last_pulse.insert(pulse.source.to_string(), pulse.clone());
        self.send(pulse)
    }

    fn send(&self, _: &Pulse) -> Vec<Pulse> {
        let mut pulses: Vec<Pulse> = Vec::new();
        let high_pulse: bool = self.last_pulse.values().any(|v| !v.high);
        for target in &self.targets {
            pulses.push(Pulse {
                high: high_pulse,
                source: self.name.to_string(),
                target: target.to_string()
            });
        }
        pulses
    }

    fn get_targets(&self) -> &Vec<String> {
        &self.targets
    }
}

#[derive(Clone, Debug)]
struct BroadcastModule {
    name: String,
    targets: Vec<String>,
}

impl BroadcastModule {
    fn new(name: &str, targets: Vec<&str>) -> BroadcastModule {
        let targets: Vec<String> = targets.iter().map(|s| s.to_string()).collect();
        BroadcastModule { name: name.to_string(), targets }
    }
}

impl Module for BroadcastModule {
    fn receive(&mut self, pulse: &Pulse) -> Vec<Pulse> {
        self.send(pulse)
    }

    fn send(&self, pulse: &Pulse) -> Vec<Pulse> {
        let mut pulses: Vec<Pulse> = Vec::new();
        for target in &self.targets {
            pulses.push(Pulse {
                high: pulse.high,
                source: self.name.to_string(),
                target: target.to_string()
            });
        }
        pulses
    }

    fn get_targets(&self) -> &Vec<String> {
        &self.targets
    }
}

trait Module {
    fn receive(&mut self, pulse: &Pulse) -> Vec<Pulse>;
    fn send(&self, pulse: &Pulse) -> Vec<Pulse>;
    fn get_targets(&self) -> &Vec<String>;
}

fn parse_input(input: &str) -> HashMap<String, Box<dyn Module>> {
    let mut modules: HashMap<String, Box<dyn Module>> = HashMap::new();
    for line in input.lines() {
        let split: Vec<&str> = line.split(" -> ").collect();
        let name: &str = split[0];
        let targets: Vec<&str> = split[1].split(", ").collect();
        if name == "broadcaster" {
            let module  = BroadcastModule::new(name, targets);
            modules.insert(module.name.clone(), Box::new(module));
        } else if name.chars().nth(0).unwrap() == '%' {
            let module  = FlipFlopModule::new(&name[1..], targets);
            modules.insert(module.name.clone(), Box::new(module));
        }
    }
    for line in input.lines() {
        let split: Vec<&str> = line.split(" -> ").collect();
        let name: &str = split[0];
        let targets: Vec<&str> = split[1].split(", ").collect();
        if name.chars().nth(0).unwrap() == '&' {
            let mut inputs: Vec<String> = Vec::new();
            for (source_name, source_module) in &modules {
                for target in source_module.get_targets() {
                    if *target == name[1..] {
                        inputs.push(source_name.clone());
                    } 
                }
            }
            let module  = ConjunctionModule::new(&name[1..], targets, inputs );
            modules.insert(module.name.clone(), Box::new(module));
        }
    }
    modules
}


fn part_uno(_input: &str) -> u64 {
    // Solves part 1
    let mut low_pulses: u64 = 0;
    let mut high_pulses: u64 = 0;
    let mut modules: HashMap<String, Box<dyn Module>> = parse_input(_input);
    let mut pulses: VecDeque<Pulse> = VecDeque::new();
    for _ in 0..1000 {
        pulses.push_back(Pulse {
            high: false,
            source: String::from("button"),
            target: String::from("broadcaster"),
        });
        while let Some(pulse) = pulses.pop_front() {
            // println!("{pulse}");
            if pulse.high { high_pulses += 1; }
            else { low_pulses += 1; }
            let module: Option<&mut Box<dyn Module>> = modules.get_mut(&pulse.target);
            if module.is_some() {
                pulses.extend(module.unwrap().receive(&pulse));
            }
        }
    }
    let answer: u64 = low_pulses * high_pulses;
    println!("{answer}");
    answer
}

fn part_dos(_input: &str) -> u64 {
    // Solves part 2 through GraphViz magic..
    let cycles: Vec<u64> = vec![4051, 3877, 3847, 3797];
    let answer: u64 = cycles.into_iter().reduce(lcm).unwrap();
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
    // use super::{part_uno, part_dos};
    use super::part_uno;
    #[test]
    fn test() {
        // Reads input and runs solve functions for parts
        let input: &str = include_str!("../test.txt").trim();
        assert_eq!(part_uno(input), 11687500);
        // assert_eq!(part_dos(input), 0);
    }
}
