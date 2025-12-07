// Advent of Code day 6
import getDayInput from "../lib/input-parser.ts"


const DAY: string = "d06"
const TEST: boolean = false


function parseInput(input: string[]): [string[][], string[]] {
    const columns: string[][] = []
    const operators: string[] = []
    const numberIdxs: number[] = []
    input.pop()!.split("").map((char, i) => {
        if (char === "+" || char === "*") {
            operators.push(char)
            numberIdxs.push(i)
        }
    })
    for (const line of input) {
        const values: string[] = []
        for (let i = 0; i < numberIdxs.length; i++) {
            const start = numberIdxs[i]!
            const end = numberIdxs[i + 1]
            values.push(line.slice(start, end ? end - 1 : undefined))
        }
        for (const [idx, value] of values.entries()) {
            columns[idx] ?
                columns[idx].push(value)
                : columns[idx] = [value]
        }
    }
    return [columns, operators]
}



const solvePart1 = (columns: string[][], operators: string[]): number => {
    const sums: number[] = columns.map((values, idx) => {
        if (operators[idx] === "*") {
            return values.reduce((a, b) => a * parseInt(b), 1)
        }
        return values.reduce((a, b) => a + parseInt(b), 0)
    })
    if (TEST) {
        console.log(columns, operators, sums)
    }
    return sums.reduce((a, b) => a + b, 0)
}


const solvePart2 = (columns: string[][], operators: string[]): number => {
    const cephColumns: string[][] = Array.from(
        { length: columns.length },
        () => Array(columns[0]![0]!.length).fill("")
    )
    columns.map((values, colIdx) => {
        for (let i = 0; i < values[0]!.length; i++) {
            for (let number of values) {
                const char = number[i]
                if (!char?.trim()) continue
                cephColumns[colIdx]![i] += char
            }
        }
    })
    const sums: number[] = cephColumns.map((values, idx) => {
        if (operators[idx] === "*") {
            return values.reduce((a, b) => a * (parseInt(b) || 1), 1)
        }
        return values.reduce((a, b) => a + (parseInt(b) || 0), 0)
    })
    if (TEST) {
        console.log(cephColumns, operators, sums)
    }
    return sums.reduce((a, b) => a + b, 0)
}


export const main = (input: string[]): number[] => {
    const [columns, operators] = parseInput(input)
    const part1 = solvePart1(columns, operators)
    console.log(part1)
    const part2 = solvePart2(columns, operators)
    console.log(part2)
    return [part1, part2]
}


if (import.meta.main) {
    const input = getDayInput(DAY, TEST)
    main(input)
}
