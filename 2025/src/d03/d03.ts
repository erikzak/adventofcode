// Advent of Code day 3
import getDayInput from "../lib/input-parser.ts"


const DAY: string = "d03"
const TEST: boolean = false


interface Candidate {
    batteries: string
    joltage: number
}


class Bank {
    batteries: string

    constructor(batteries: string) {
        this.batteries = batteries
    }

    maxJoltage(count: number): number {
        // Finds the maximum joltage that can be achieved by turning on `count` batteries
        const maxStartingJoltage: string = Math.max(...this.batteries
            .substring(0, this.batteries.length - count + 1)
            .split("")
            .map(Number)
        ).toString()
        let index: number = this.batteries.indexOf(maxStartingJoltage)
        let candidate: Candidate = {
            batteries: maxStartingJoltage,
            joltage: Number(maxStartingJoltage),
        }
        do {
            index++
            let bestBattery: number = Number(this.batteries[index])
            for (let i = index; i < this.batteries.length - count + 1 + candidate.batteries.length; i++) {
                const check = Number(this.batteries[i])
                if (check > bestBattery) {
                    bestBattery = check
                    index = i
                }
                if (bestBattery === 9) {
                    break
                }
            }
            const newBattery: string = candidate.batteries + this.batteries[index]
            candidate = {
                batteries: newBattery,
                joltage: Number(newBattery),
            }
        } while (candidate.batteries.length < count)
        if (TEST) {
            console.log(`${this.batteries}: ${candidate.joltage}`)
        }
        return candidate.joltage
    }
}


const solvePart1 = (input: string[]): number => {
    // Solve part 1
    const banks: Bank[] = input.map(line => new Bank(line))
    return banks.reduce((sum, bank) => sum + bank.maxJoltage(2), 0)
}


const solvePart2 = (input: string[]): number => {
    // Solve part 2
    const banks: Bank[] = input.map(line => new Bank(line))
    return banks.reduce((sum, bank) => sum + bank.maxJoltage(12), 0)
}


export const main = (input: string[]): number[] => {
    const part1 = solvePart1(input)
    console.log(part1)
    const part2 = solvePart2(input)
    console.log(part2)
    return [part1, part2]
}


if (import.meta.main) {
    const input = getDayInput(DAY, TEST)
    main(input)
}
