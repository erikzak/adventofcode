// Advent of Code day 4
import getDayInput from "../lib/input-parser.ts"
import { Grid } from "../lib/grid.ts"


const DAY: string = "d04"
const TEST: boolean = false


const NEIGHBORHOOD: [number, number][] = [
    [-1, -1], [0, -1], [1, -1],
    [-1,  0],          [1,  0],
    [-1,  1], [0,  1], [1,  1],
]


const solvePart1 = (input: string[]): number => {
    // Solve part 1
    const grid = new Grid(input)
    let accessibleRolls = 0
    for (const node of grid.nodes.values()) {
        if (node.value !== "@") {
            continue
        }
        let adjacentRolls = 0
        for (const idx of NEIGHBORHOOD) {
            const neighbor = grid.getNode(node.coords.x + idx[0], node.coords.y + idx[1])
            if (!neighbor) {
                continue
            }
            if (neighbor.value === "@") {
                adjacentRolls++
            }
        }
        if (adjacentRolls < 4) {
            accessibleRolls++
        }
    }
    return accessibleRolls
}


const solvePart2 = (input: string[]): number => {
    // Solve part 2
    const grid = new Grid(input)
    let rollsRemoved = 0
    let rollsRemovedThisRound = 1
    while (rollsRemovedThisRound > 0) {
        rollsRemovedThisRound = 0
        for (const node of grid.nodes.values()) {
            if (node.value !== "@") {
                continue
            }
            let adjacentRolls = 0
            for (const idx of NEIGHBORHOOD) {
                const neighbor = grid.getNode(node.coords.x + idx[0], node.coords.y + idx[1])
                if (!neighbor) {
                    continue
                }
                if (neighbor.value === "@") {
                    adjacentRolls++
                }
            }
            if (adjacentRolls < 4) {
                node.value = "."
                rollsRemoved++
                rollsRemovedThisRound++
            }
        }
    }
    return rollsRemoved
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
