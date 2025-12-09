// Advent of Code day 9
import getDayInput from "../lib/input-parser.ts"
import type { Node } from "../lib/grid.ts"


const DAY: string = "d09"
const TEST: boolean = false


function solvePart1(input: string[]): number {
    const redNodes: Node[] = input.map(line => {
        const [x, y]: [number, number] = line.split(",").map(Number) as [number, number]
        return { key: `${x},${y}`, coords: { x, y } }
    })
    let maxArea = 0
    for (let i = 0; i < redNodes.length; i++) {
        for (let j = i + 1; j < redNodes.length; j++) {
            const a = redNodes[i]!
            const b = redNodes[j]!
            const area = (Math.abs(a.coords.x - b.coords.x) + 1) * (Math.abs(a.coords.y - b.coords.y) + 1)
            if (area > maxArea) {
                maxArea = area
            }
        }
    }
    return maxArea
}


function nodeIsInside(node: Node, redNodes: Node[], insideNodes: {[key: string]: Node}): boolean {
    /**
     Determine if the point is on the path, corner, or boundary of the polygon.
     https://en.wikipedia.org/wiki/Even%E2%80%93odd_rule
     */
    if (insideNodes[node.key]) {
        return insideNodes[node.key]!.value === "X"
    }
    const { x, y } = node.coords
    // Check if point is exactly on a vertex or edge
    for (let i = 0; i < redNodes.length; i++) {
        const curr = redNodes[i]!
        const next = redNodes[(i + 1) % redNodes.length]!
        // Point is a vertex
        if (x === curr.coords.x && y === curr.coords.y) {
            node.value = "X"
            insideNodes[node.key] = node
            return true
        }
        // Point is on an edge (between curr and next)
        const minX = Math.min(curr.coords.x, next.coords.x)
        const maxX = Math.max(curr.coords.x, next.coords.x)
        const minY = Math.min(curr.coords.y, next.coords.y)
        const maxY = Math.max(curr.coords.y, next.coords.y)
        if (x >= minX && x <= maxX && y >= minY && y <= maxY) {
            // Check if point is collinear with the edge
            const cross = (next.coords.y - curr.coords.y) * (x - curr.coords.x) - 
                         (next.coords.x - curr.coords.x) * (y - curr.coords.y)
            if (cross === 0) {
                node.value = "X"
                insideNodes[node.key] = node
                return true
            }
        }
    }
    // Standard even-odd ray casting
    let inside = false
    for (let i = 0, j = redNodes.length - 1; i < redNodes.length; j = i++) {
        const [xi, yi]: [number, number] = [redNodes[i]!.coords.x, redNodes[i]!.coords.y]
        const [xj, yj]: [number, number] = [redNodes[j]!.coords.x, redNodes[j]!.coords.y]
        const intersect = ((yi > y) != (yj > y))
            && (x < (xj - xi) * (y - yi) / (yj - yi) + xi)
        if (intersect) inside = !inside
    }
    node.value = inside ? "X" : "."
    insideNodes[node.key] = node
    return inside
}


function rectangleIsInside(a: Node, b: Node, redNodes: Node[], insideNodes: {[key: string]: Node}): boolean {
    const minX: number = Math.min(a.coords.x, b.coords.x)
    const maxX: number = Math.max(a.coords.x, b.coords.x)
    const minY: number = Math.min(a.coords.y, b.coords.y)
    const maxY: number = Math.max(a.coords.y, b.coords.y)
    for (let x of [minX, maxX]) {
        for (let y=minY; y <= maxY; y++) {
            if (!nodeIsInside({ key: `${x},${y}`, coords: { x, y } }, redNodes, insideNodes)) {
                return false
            }
        }
    }
    for (let y of [minY, maxY]) {
        for (let x=minX; x <= maxX; x++) {
            if (!nodeIsInside({ key: `${x},${y}`, coords: { x, y } }, redNodes, insideNodes)) {
                return false
            }
        }
    }
    return true
}


function compressCoordinates(nodes: Node[]): { compressedNodes: Node[], xMap: Map<number, number>, yMap: Map<number, number> } {
    // Extract unique x and y coordinates and sort them
    const xCoords = [...new Set(nodes.map(n => n.coords.x))].sort((a, b) => a - b)
    const yCoords = [...new Set(nodes.map(n => n.coords.y))].sort((a, b) => a - b)
    // Create mappings from original to compressed coordinates
    const xMap = new Map<number, number>()
    const yMap = new Map<number, number>()
    xCoords.forEach((x, i) => xMap.set(x, i))
    yCoords.forEach((y, i) => yMap.set(y, i))
    // Create compressed nodes
    const compressedNodes = nodes.map(node => ({
        key: `${xMap.get(node.coords.x)},${yMap.get(node.coords.y)}`,
        coords: { x: xMap.get(node.coords.x)!, y: yMap.get(node.coords.y)! }
    }))
    return { compressedNodes, xMap, yMap }
}


function printGrid(redNodes: Node[], insideNodes: {[key: string]: Node}): void {
    for (let y = 0; y < Math.max(...redNodes.map(node => node.coords.y)) + 2; y++) {
        let line = ""
        for (let x = 0; x < Math.max(...redNodes.map(node => node.coords.x)) + 3; x++) {
            const nodeKey = `${x},${y}`
            if (redNodes.find(node => node.key === nodeKey)) {
                line += "#"
            } else if (insideNodes[nodeKey]) {
                line += insideNodes[nodeKey].value
            } else {
                line += "."
            }
        }
        console.log(line)
    }
}


function solvePart2(input: string[]): number {
    const originalNodes: Node[] = input.map(line => {
        const [x, y]: [number, number] = line.split(",").map(Number) as [number, number]
        return { key: `${x},${y}`, coords: { x, y } }
    })
    // Compress coordinates to work with smaller grid
    const { compressedNodes: redNodes, xMap, yMap } = compressCoordinates(originalNodes)
    // Create reverse maps for calculating actual areas
    const reverseXMap = new Map<number, number>()
    const reverseYMap = new Map<number, number>()
    xMap.forEach((compressed, original) => reverseXMap.set(compressed, original))
    yMap.forEach((compressed, original) => reverseYMap.set(compressed, original))
    const insideNodes: {[key: string]: Node} = {}
    if (TEST) {
        for (let i = 0; i < redNodes.length; i++) {
            for (let j = i + 1; j < redNodes.length; j++) {
                const a = redNodes[i]!
                const b = redNodes[j]!
                for (let x = Math.min(a.coords.x, b.coords.x); x <= Math.max(a.coords.x, b.coords.x); x++) {
                    for (let y = Math.min(a.coords.y, b.coords.y); y <= Math.max(a.coords.y, b.coords.y); y++) {
                        const node: Node = { key: `${x},${y}`, coords: { x, y } }
                        if (nodeIsInside(node, redNodes, insideNodes)) {
                            node.value = "X"
                            insideNodes[node.key] = node
                        } else {
                            node.value = "."
                            insideNodes[node.key] = node
                        }
                    }
                }
            }
        }
        printGrid(redNodes, insideNodes)
    }
    let maxArea = 0
    for (let i = 0; i < redNodes.length - 1; i++) {
        for (let j = i + 1; j < redNodes.length; j++) {
            const a = redNodes[i]!
            const b = redNodes[j]!
            // Check if all nodes are inside our boundary (using compressed coordinates)
            if (!rectangleIsInside(a, b, redNodes, insideNodes)) {
                continue
            }
            // Calculate actual area using original coordinates
            const origX1 = reverseXMap.get(a.coords.x)!
            const origX2 = reverseXMap.get(b.coords.x)!
            const origY1 = reverseYMap.get(a.coords.y)!
            const origY2 = reverseYMap.get(b.coords.y)!
            const area = (Math.abs(origX1 - origX2) + 1) * (Math.abs(origY1 - origY2) + 1)
            if (area > maxArea) {
                maxArea = area
            }
        }
    }
    return maxArea
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
