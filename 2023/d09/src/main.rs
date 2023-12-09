//! Advent of Code 2023 day 9

#[derive(Debug)]
struct Value {
    history: Vec<i64>,
    differences: Vec<Vec<i64>>,
}


impl Value {
    fn new(history: Vec<i64>) -> Value {
        let mut value: Value = Value{history, differences: Vec::new()};
        value.differences = value.generate_differences();
        value
    }

    fn get_differences(&self, history: &Vec<i64>) -> Vec<i64> {
        let diffs: Vec<i64> = history
            .iter()
            .enumerate()
            .filter(|&(i, _)| i < history.len() - 1)
            .map(|(i, v)| history[i+1] - v)
            .collect();
        diffs
    }

    fn generate_differences(&mut self) -> Vec<Vec<i64>> {
        let mut differences: Vec<Vec<i64>> = vec!(self.get_differences(&self.history));
        while !differences
            .last()
            .unwrap()
            .iter()
            .all(|v: &i64| v == &0_i64) {
            differences.push(self.get_differences(differences.last().unwrap()));
        }
        differences
    }

    fn extrapolate(&self) -> i64 {
        let mut prediction: i64 = *self.differences.last().unwrap().last().unwrap();
        for diffs in self.differences.iter().rev().skip(1) {
            // Iterate backwards over difference vecs
            prediction = diffs.last().unwrap() + prediction;
        }
        prediction = self.history.last().unwrap() + prediction;
        prediction
    }

    fn extrapolate_backwards(&self) -> i64 {
        let mut prediction: i64 = *self.differences.last().unwrap().first().unwrap();
        for diffs in self.differences.iter().rev().skip(1) {
            // Iterate backwards over difference vecs
            prediction = diffs.first().unwrap() - prediction;
        }
        prediction = self.history.first().unwrap() - prediction;
        prediction
    }
}


fn part_uno(_input: &str) -> i64 {
    // Solves part 1
    let mut answer: i64 = 0;
    for line in _input.lines() {
        let value_history: Vec<i64> = line
            .split_whitespace()
            .map(|v: &str| v.parse().unwrap())
            .collect();
        let value: Value = Value::new(value_history);
        answer += value.extrapolate();
    }
    println!("{answer}");
    answer
}


fn part_dos(_input: &str) -> i64 {
    // Solves part 2
    let mut answer: i64 = 0;
    for line in _input.lines() {
        let value_history: Vec<i64> = line
            .split_whitespace()
            .map(|v: &str| v.parse().unwrap())
            .collect();
        let value: Value = Value::new(value_history);
        let prediction: i64 = value.extrapolate_backwards();
        answer += prediction;
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
        assert_eq!(part_uno(input), 114);
        assert_eq!(part_dos(input), 2);
    }
}
