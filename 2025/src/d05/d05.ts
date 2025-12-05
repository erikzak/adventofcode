// Advent of Code day 5
import getDayInput from "../lib/input-parser.ts"


const DAY: string = "d05"
const TEST: boolean = false


function parseInput(input: string[]): [[number, number][], number[]] {
    const ranges: [number, number][] = input.filter((line) => line.includes("-")).map(line => {
        const [start, end] = line.split("-").map(Number) as [number, number]
        return [start, end]
    })
    const items: number[] = input.filter((line) => !line.includes("-")).map(Number)
    return [ranges, items]
}


const solvePart1 = (input: string[]): number => {
    // Solve part 1
    const [ranges, items] = parseInput(input)
    let fresh: number = 0
    for (const item of items) {
        for (const [start, end] of ranges) {
            if (item >= start && item <= end) {
                fresh++
                break
            }
        }
    }
    return fresh
}


const solvePart2 = (input: string[]): number => {
    // Solve part 2
    const [ranges, _] = parseInput(input)
    const distinctRanges: [number, number][] = []
    for (let [startToAdd, endToAdd] of ranges) {
        // Get intersecting ranges
        const intersectingRangeIdxs: number[] = []
        for (let i=0; i < distinctRanges.length; i++) {
            const [start, end] = distinctRanges[i]!
            if (startToAdd <= end && start <= endToAdd) {
                intersectingRangeIdxs.push(i)
            }
        }
        if (!intersectingRangeIdxs.length) {
            // Does not intersect any existing range
            distinctRanges.push([startToAdd, endToAdd])
            continue
        }
        const intersectingRanges: [number, number][] = []
        // Remove intersecting ranges from distinctRanges
        for (let i = intersectingRangeIdxs.length - 1; i >= 0; i--) {
            intersectingRanges.push(distinctRanges.splice(intersectingRangeIdxs[i]!, 1)[0]!)
        }
        // Merge intersecting ranges with new range
        startToAdd = Math.min(...intersectingRanges.map(([s, _]) => s), startToAdd)
        endToAdd = Math.max(...intersectingRanges.map(([_, e]) => e), endToAdd)
        distinctRanges.push([startToAdd, endToAdd])
    }
    if (TEST) {
        console.log(distinctRanges)
    }
    const fresh = distinctRanges.reduce((sum, [start, end]) => sum + (end - start + 1), 0)
    return fresh
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
