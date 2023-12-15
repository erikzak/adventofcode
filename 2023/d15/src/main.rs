//! Advent of Code 2023 day 15
use std::collections::HashMap;


struct Box {
    lenses: Vec<Lens>,
}


impl Box {
    fn new() -> Box {
        Box{ lenses: Vec::new() }
    }

    fn remove_lens(&mut self, label: &str) {
        for (i, lens) in self.lenses.iter().enumerate() {
            if lens.label == label {
                self.lenses.remove(i);
                return;
            }
        }
    }

    fn upsert_lens(&mut self, label: &str, focal_length: u8) {
        for lens in self.lenses.iter_mut() {
            if lens.label == label {
                lens.focal_length = focal_length;
                return;
            }
        }
        self.lenses.push(Lens{ label: label.to_string(), focal_length });
    }
}


struct Lens {
    label: String,
    focal_length: u8,
}


fn hash(value: &str) -> u8 {
    let mut current_value: u32 = 0;
    for c in value.chars() {
        current_value += c as u32;
        current_value *= 17;
        current_value = current_value % 256;
    }
    current_value as u8
}


fn part_uno(_input: &str) -> u32 {
    // Solves part 1
    let mut answer: u32 = 0;
    for part in _input.split(",") {
        answer += hash(part) as u32;
    }
    println!("{answer}");
    answer
}


fn part_dos(_input: &str) -> u32 {
    // Solves part 2
    let mut answer: u32 = 0;
    let mut boxes: HashMap<u8, Box> = HashMap::new();
    for part in _input.split(",") {
        let operation: char;
        let split: Vec<&str>;
        if part.contains("=") {
            operation = '=';
            split = part.split(operation).collect();
        } else {
            operation = '-';
            split = part.split(operation).collect();
        }
        let label: &str = split[0];
        let box_number: u8 = hash(label);
        if operation == '-' {
            if boxes.contains_key(&box_number) {
                let bøx: &mut Box = boxes.get_mut(&box_number).unwrap();
                bøx.remove_lens(&label);
            }
        } else {
            let focal_length: u8 = split[1].parse().expect(format!("what: {:?}", part).as_str());
            if !boxes.contains_key(&box_number) {
                boxes.insert(box_number, Box::new());
            }
            let bøx: &mut Box = boxes.get_mut(&box_number).unwrap();
            bøx.upsert_lens(label, focal_length);
        }
    }
    for (box_number, bøx) in boxes {
        for (i, lens) in bøx.lenses.iter().enumerate() {
            let power: u32 = (1 + box_number as u32) * (i as u32 + 1) * lens.focal_length as u32;
            answer += power;
        }
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
        assert_eq!(part_uno(input), 1320);
        assert_eq!(part_dos(input), 145);
    }
}
