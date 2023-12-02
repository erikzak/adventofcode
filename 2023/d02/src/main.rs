//! Advent of Code 2023 day 2
use std::fs;


fn read_file(file_path: &str) -> String {
    // Returns file contents as string
    fs::read_to_string(file_path)
        .expect("wtf")
        .trim().to_string()
}


fn get_game_id_and_sets(game: &str) -> (i32, Vec<&str>) {
    // Returns game id and list of set strings
    let game_split: Vec<&str> = game.split(": ").collect();
    let id: i32 = game_split[0].rsplit_once(" ").unwrap().1.parse().unwrap();
    let sets: Vec<&str> = game_split[1].split("; ").collect();
    return (id, sets);
}


fn get_cube_color_and_count(set: &str) -> (&str, i32) {
    // Returns color and count of cubes in subset
    let split: Vec<&str> = set.split(" ").collect();
    let count: i32 = split[0].parse().unwrap();
    let color: &str = split[1];
    return (color, count);
}


fn validate_game(game: &str) -> i32 {
    // Returns the game id if all sets are valid, else 0
    let (id, sets) = get_game_id_and_sets(game);
    for set in sets {
        let subsets: Vec<&str> = set.split(", ").collect();
        for cubes in subsets{
            let (color, count) = get_cube_color_and_count(cubes);
            if color == "red" && count > 12 {
                return 0;
            }
            if color == "green" && count > 13 {
                return 0;
            }
            if color == "blue" && count > 14 {
                return 0;
            }
        }
    }
    return id;
}


fn get_game_power(game: &str) -> i32 {
    // Returns minimum red, blue and green cubes for game multiplied
    let (_, sets) = get_game_id_and_sets(game);
    let mut min_red = 0;
    let mut min_green = 0;
    let mut min_blue = 0;
    for set in sets {
        let subsets: Vec<&str> = set.split(", ").collect();
        for cubes in subsets{
            let (color, count) = get_cube_color_and_count(cubes);
            if color == "red" && count > min_red {
                min_red = count;
            }
            if color == "green" && count > min_green {
                min_green = count;
            }
            if color == "blue" && count > min_blue {
                min_blue = count;
            }
        }
    }
    min_red * min_green * min_blue
}


fn part_uno(input: &str) {
    // Solves part 1
    let mut answer: i32 = 0;
    for game in input.lines() {
        answer += validate_game(game);
    }
    println!("Part one: {answer}");
}


fn part_dos(input: &str) {
    // Solves part 2
    let mut answer: i32 = 0;
    for game in input.lines() {
        answer += get_game_power(game);
    }
    println!("Part two: {answer}");
}


fn main() {
    // Reads input and runs solve functions for parts
    let file_path: &str = "input.txt";
    let input: String = read_file(file_path);
    part_uno(&input);
    part_dos(&input);
}
