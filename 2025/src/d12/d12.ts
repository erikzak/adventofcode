// Advent of Code day 12
import getDayInput from "../lib/input-parser.ts"
import { Grid } from "../lib/grid.ts"


const DAY: string = "d12"
const TEST: boolean = false


interface Gift {
    index: number
    grid: Grid
    area: number
}


interface Tree {
    width: number
    height: number
    requiredGifts: number[]
    area: number
}


function parseInput(input: string[]): [Gift[], Tree[]] {
    const firstTree: number = input.findIndex(line => line.indexOf("x") >= 0)
    const trees: Tree[] = input.slice(firstTree).map(line => {
        const [dim, giftIdxs]: [string, string] = line.split(":").map(s => s.trim()) as [string, string]
        const [width, height]: [number, number] = dim.split("x").map(Number) as [number, number]
        const requiredGifts: number[] = giftIdxs.split(" ").map(Number)
        const area = width * height
        return { width, height, requiredGifts, area }
    })
    const gifts: Gift[] = []
    for (let i=0; i<firstTree-1; i += 4) {
        const index: number = Number(input[i]!.split(":")[0]!)
        const gridLines: string[] = input.slice(i+1, i+4)
        const grid: Grid = new Grid(gridLines)
        const area = Array.from(grid.nodes.values()).reduce((sum, node) => sum + (node.value === "#" ? 1 : 0), 0)
        gifts.push({ index, grid, area })
    }
    return [gifts, trees]
}


function solvePart1(input: string[]): number {
    const [gifts, trees]: [Gift[], Tree[]] = parseInput(input)
    let canFit: number = 0
    for (const tree of trees) {
        const numGifts: number = tree.requiredGifts.reduce((sum, num) => sum + num, 0)
        if (numGifts * 9 > tree.area) continue
        canFit++
    }
    return canFit
}


export const main = (input: string[]): number[] => {
    const part1 = solvePart1(input)
    console.log(part1)
    return [part1, 0]
}


if (import.meta.main) {
    const input = getDayInput(DAY, TEST)
    main(input)
}
