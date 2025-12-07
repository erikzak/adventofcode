// Advent of Code day 7
import getDayInput from "../lib/input-parser.ts"
import { type Coordinate, Grid } from "../lib/grid.ts"


const DAY: string = "d07"
const TEST: boolean = false


const STATES: { [key: string]: number } = {}


function solvePart1(input: string[]): number {
    const grid = new Grid(input)
    let splits: number = 0
    for (const row of grid.rows) {
        for (const node of row) {
            if (node.value !== "S" && node.value !== "|") {
                continue
            }
            const down = grid.getNode(node.coords.x, node.coords.y - 1)
            if (!down) continue
            if (down.value === ".") {
                down.value = "|"
                continue
            }
            if (down.value === "^") {
                const downLeft = grid.getNode(node.coords.x - 1, node.coords.y - 1)!
                const downRight = grid.getNode(node.coords.x + 1, node.coords.y - 1)!
                let split: boolean = true
                if (downLeft.value === ".") {
                    downLeft.value = "|"
                    split = true
                }
                if (downRight.value === ".") {
                    downRight.value = "|"
                    split = true
                }
                splits += split ? 1 : 0
            }
        }
    }
    if (TEST) {
        console.log(grid.toString())
    }
    return splits
}


function getTimelines(tachyon: Coordinate, grid: Grid): number {
    const key = `${tachyon.x},${tachyon.y}`
    if (STATES[key] !== undefined) {
        return STATES[key]!
    }
    let timelines: number = 0
    const down = grid.getNode(tachyon.x, tachyon.y - 1)
    if (!down) {
        return 1
    }
    if (down.value === ".") {
        timelines += getTimelines(down.coords, grid)
    }
    if (down.value === "^") {
        const downLeft = grid.getNode(tachyon.x - 1, tachyon.y - 1)!
        const downRight = grid.getNode(tachyon.x + 1, tachyon.y - 1)!
        if (downLeft.value === ".") {
            timelines += getTimelines(downLeft.coords, grid)
        }
        if (downRight.value === ".") {
            timelines += getTimelines(downRight.coords, grid)
        }
    }
    STATES[`${tachyon.x},${tachyon.y}`] = timelines
    return timelines
}



function solvePart2(input: string[]): number {
    const grid = new Grid(input)
    const start = grid.rows[0]!.filter(node => node.value === "S").map(node => node.coords)[0]!
    return getTimelines(start, grid)
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
