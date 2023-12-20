//! Advent of Code 2023 day 19
use std::collections::HashMap;
use std::fmt;

fn parse_workflows(input: &str) -> HashMap<String, Workflow> {
    let mut workflows: HashMap<String, Workflow> = HashMap::new();
    for line in input.replace("\r\n", "\n")
            .split("\n\n").collect::<Vec<&str>>()[0]
            .lines() {
        let split: Vec<&str> = line.split("{").collect();
        let name: String = split[0].to_string();
        let rules: Vec<&str> = split[1][..split[1].len()-1].split(",").collect();
        let workflow = Workflow::new(rules);
        workflows.insert(name, workflow );
    }
    workflows
}

fn parse_parts(input: &str) -> Vec<Part> {
    let mut parts: Vec<Part> = Vec::new();
    for line in input.replace("\r\n", "\n")
            .split("\n\n").collect::<Vec<&str>>()[1]
            .lines() {
        let split: Vec<&str> = line[1..line.len()-1].split(",").collect();
        let values: HashMap<char, u64> = HashMap::from([
            ('x', split[0][2..].parse().unwrap()),
            ('m', split[1][2..].parse().unwrap()),
            ('a', split[2][2..].parse().unwrap()),
            ('s', split[3][2..].parse().unwrap()),
        ]);
        parts.push(Part { values });
    }
    parts
}

struct Workflow {
    rules: Vec<Rule>,
    default: String,
}

impl Workflow {
    fn new(input: Vec<&str>) -> Workflow {
        let mut rules: Vec<Rule> = Vec::new();
        let mut default: String = String::from("");
        for rule in input {
            if !rule.contains(":") {
                default = rule.to_string();
                continue;
            }
            rules.push(Rule::new(rule));
        }
        if default == "" {
            panic!("invalid workflow default outcome");
        }
        Workflow { rules, default }
    }

    fn check(&self, part: &Part) -> &str {
        for rule in &self.rules {
            if rule.validate(part) {
                return &rule.if_valid;
            }
        }
        &self.default
    }
}

#[derive(Debug)]
struct Rule {
    variable: char,
    operator: char,
    value: u64,
    if_valid: String,
}

impl fmt::Display for Rule {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "Rule {{ {} {} {} => {} }}", self.variable, self.operator, self.value, self.if_valid)
    }
}

impl Rule {
    fn new(input: &str) -> Rule {
        let variable: char = input.chars().nth(0).unwrap();
        let operator: char = input.chars().nth(1).unwrap();
        let split: Vec<&str> = input[2..].split(":").collect();
        let value: u64 = split[0].parse().unwrap();
        let if_valid: String = split[1].to_string();
        Rule { variable, operator, value, if_valid }
    }

    fn validate(&self, part: &Part) -> bool {
        if (self.operator == '>' && part.values[&self.variable] > self.value) ||
            (self.operator == '<' && part.values[&self.variable] < self.value) {
                return true;
            }
        false
    }
}

#[derive(Debug)]
struct Part {
    values: HashMap<char, u64>,
}

impl Part {
    fn get_value(&self) -> u64 {
        self.values[&'x'] + self.values[&'m'] + self.values[&'a'] + self.values[&'s']
    }
}

#[derive(Clone, Debug)]
struct Combination {
    x: [u64; 2],
    m: [u64; 2],
    a: [u64; 2],
    s: [u64; 2],
    next_workflow: String,
}

impl Combination {
    fn new() -> Combination {
        Combination {
            x: [1, 4000],
            m: [1, 4000],
            a: [1, 4000],
            s: [1, 4000],
            next_workflow: String::from("in"),
        }
    }

    fn update(&mut self, rule: &Rule, inverse: bool) {
        if rule.operator == '>' {
            match rule.variable {
                'x' => if !inverse { self.x[0] = rule.value + 1 } else { self.x[1] = rule.value },
                'm' => if !inverse { self.m[0] = rule.value + 1 } else { self.m[1] = rule.value },
                'a' => if !inverse { self.a[0] = rule.value + 1 } else { self.a[1] = rule.value },
                's' => if !inverse { self.s[0] = rule.value + 1 } else { self.s[1] = rule.value },
                _ => panic!("unhandled rule variable"),
            };
        } else if rule.operator == '<' {
            match rule.variable {
                'x' => if !inverse { self.x[1] = rule.value - 1 } else { self.x[0] = rule.value },
                'm' => if !inverse { self.m[1] = rule.value - 1 } else { self.m[0] = rule.value },
                'a' => if !inverse { self.a[1] = rule.value - 1 } else { self.a[0] = rule.value },
                's' => if !inverse { self.s[1] = rule.value - 1 } else { self.s[0] = rule.value },
                _ => panic!("unhandled rule variable"),
            };
        }
    }

    fn distinct_combinations(&self) -> u64 {
        (self.x[1] - self.x[0] + 1) *
        (self.m[1] - self.m[0] + 1) *
        (self.a[1] - self.a[0] + 1) *
        (self.s[1] - self.s[0] + 1)
    }
}


fn part_uno(_input: &str) -> u64 {
    // Solves part 1
    let mut answer: u64 = 0;
    let workflows: HashMap<String, Workflow> = parse_workflows(_input);
    let parts: Vec<Part> = parse_parts(_input);
    for part in parts {
        let mut workflow: &str = "in";
        while workflow != "A" && workflow != "R" {
            workflow = workflows[workflow].check(&part)
        }
        if workflow == "A" {
            answer += part.get_value();
        }
    }
    println!("{answer}");
    answer
}

fn part_dos(_input: &str) -> u64 {
    // Solves part 2 with debug heaven
    let mut answer: u64 = 0;
    let workflows: HashMap<String, Workflow> = parse_workflows(_input);
    let mut combinations: Vec<Combination> = vec![Combination::new()];
    loop {
        let mut combo: Combination = combinations.pop().unwrap();
        // println!("\nchecking {:?}", combo);
        let workflow: &Workflow = &workflows[&combo.next_workflow];
        for rule in &workflow.rules {
            // println!("processing {rule}");
            let mut split: Combination = combo.clone();
            combo.update(rule, false);
            split.update(rule, true);
            if rule.if_valid == "A" {
                answer += combo.distinct_combinations();
                // println!("rule => A: adding {} to answer from {:?}", combo.distinct_combinations(), combo);
            } else if rule.if_valid != "R" {
                combo.next_workflow = rule.if_valid.to_string();
                // println!("queuing {:?}", combo);
                combinations.push(combo);
            }
            combo = split;
            // println!("\nchecking {:?}", combo);
        }

        if workflow.default == "A" {
            answer += combo.distinct_combinations();
            // println!("workflow default => A: adding {} to answer from {:?}", combo.distinct_combinations(), combo);
        } else if workflow.default != "R" {
            combo.next_workflow = workflow.default.to_string();
            // println!("queuing {:?}", combo);
            combinations.push(combo);
        }

        if combinations.len() == 0 { break; }
        // println!("combinations: {:?}\n", combinations);
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
        assert_eq!(part_uno(input), 19114);
        assert_eq!(part_dos(input), 167409079868000);
    }
}
