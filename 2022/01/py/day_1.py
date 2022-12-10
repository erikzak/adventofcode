"""
Advent of Code, day 1.

https://adventofcode.com/2022/day/1
"""
from typing import List


INPUT = "input.txt"


class Food:
    """Keeps track of a food item and its calorie count."""
    def __init__(self, calories) -> None:
        self.calories = calories
        return


class Elf:
    """Keeps track of an elf's foodstuff and sum calories."""
    def __init__(self) -> None:
        self.foodstuff = []
        self.sum_calories = 0
        return

    def add_food(self, food: Food) -> None:
        """Adds a food item to the elf, and adds the calorie count to its sum."""
        self.foodstuff.append(food)
        self.sum_calories += food.calories
        return


def parse_elves(input_txt: str) -> List[Elf]:
    """
    Parses puzzle input from txt file.

    Returns a list of elves with foodstuff.
    """
    elves = []
    with open(input_txt) as reader:
        elf = Elf()
        for line in reader:
            line = line.strip()
            if (not line):
                elves.append(elf)
                elf = Elf()
                continue
            calories = int(line)
            food = Food(calories)
            elf.add_food(food)
        if (elf.foodstuff):
            elves.append(elf)
    return elves


def main() -> None:
    """
    Parses input as txt into a list of elves then sorts them by sum calories
    to find the ones with the most calories.
    """
    elves = parse_elves(INPUT)

    # Sort by sum calories
    elves.sort(key=lambda elf: elf.sum_calories)

    # Part 1: print sum calories of elf with the most calories
    print(elves[-1].sum_calories)

    # Part 2: print sum calories of the three elves with the most calories
    top_three_sum_calories = [elf.sum_calories for elf in elves[-3:]]
    print(sum(top_three_sum_calories))
    return


if (__name__ == "__main__"):
    main()
