// Advent of Code day 11
import { LRUCache } from "lru-cache"
import getDayInput from "../lib/input-parser.ts"


const DAY: string = "d11"
const TEST: boolean = false


interface Server {
    id: string
    outputs: string[]
}


const cache = new LRUCache<string, number>({ max: 999999 })


function dfs(serverId: string, targetId: string, servers: {[id: string]: Server}, visited: Set<string>): number {
    const key = `${serverId},${targetId}`
    if (cache.has(key)) {
        return cache.get(key)!
    }
    if (visited.has(serverId)) {
        return 0
    }
    visited.add(serverId)
    let paths = 0
    if (!servers[serverId]) {
        visited.delete(serverId)
        cache.set(key, 0)
        return 0
    }
    const server = servers[serverId]!
    for (const outputId of server.outputs) {
        if (outputId === targetId) {
            paths++
            continue
        }
        paths += dfs(outputId, targetId, servers, visited)
    }
    visited.delete(serverId)
    cache.set(key, paths)
    return paths
}


function parseInput(input: string[]): {[id: string]: Server} {
    const servers: {[id: string]: Server} = {}
    for (const line of input) {
        const split = line.split(":").map(s => s.trim())
        const id = split.shift()!
        const outputs = split[0]!.split(" ")
        servers[id] = { id, outputs }
    }
    return servers
}


function solvePart1(input: string[]): number {
    if (TEST) {
        input = input.slice(0, input.length / 2)
    }
    const servers = parseInput(input)
    cache.clear()
    const paths: number = dfs("you", "out", servers, new Set())
    return paths
}


function solvePart2(input: string[]): number {
    if (TEST) {
        input = input.slice(input.length / 2 - 1)
    }
    const servers = parseInput(input)
    cache.clear()
    const paths: number = (
        dfs("svr", "dac", servers, new Set()) * dfs("dac", "fft", servers, new Set()) * dfs("fft", "out", servers, new Set())
        + dfs("svr", "fft", servers, new Set()) * dfs("fft", "dac", servers, new Set()) * dfs("dac", "out", servers, new Set())
    )
    return paths
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
