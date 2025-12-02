// Advent of Code day 1
import getDayInput from "../lib/input-parser.ts"


const DAY: string = "d01"
const TEST: boolean = false


const solvePart1 = (input: string[]): number => {
    // Count how many times the dial hits zero after rotations
    let value: number = 50
    let zeros: number = 0
    for (const line of input) {
        const rotation = getRotation(line)
        value += rotation
        if (value % 100 === 0) {
            zeros += 1
        }
    }
    return zeros
}


const solvePart2 = (input: string[]): number => {
    // Count how many times the dial hits zero *during* rotations
    let value: number = 50
    let zeros: number = 0
    for (const line of input) {
        let rotation = getRotation(line)
        while (rotation !== 0) {
            const step = rotation > 0 ? 1 : -1
            value += step
            if (value % 100 === 0) {
                zeros += 1
            }
            rotation -= step
        }
    }
    return zeros
}


const getRotation = (line: string): number => {
    const direction: string = line[0] as string
    const steps: number = parseInt(line.slice(1))
    return direction === "L" ? -steps : steps
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
