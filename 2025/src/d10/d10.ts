// Advent of Code day 10
import { init } from "z3-solver";

import getDayInput from "../lib/input-parser.ts"


const DAY: string = "d10"
const TEST: boolean = false


interface IMachine {
    lights: string
    lightDiagram: string
    buttons: number[][]
    requiredJoltage: number[]
    pushes: number
}


class Machine implements IMachine {
    lights: string
    lightDiagram: string
    buttons: number[][]
    requiredJoltage: number[]
    pushes: number

    constructor(lightDiagram: string, buttons: number[][], requiredJoltage: number[]) {
        this.lightDiagram = lightDiagram
        this.lights = ".".repeat(lightDiagram.length)
        this.buttons = buttons
        this.requiredJoltage = requiredJoltage
        this.pushes = 0
    }

    push(button: number[]): void {
        for (const idx of button) {
            this.lights = this.lights.slice(0, idx) +
                (this.lights[idx] === "." ? "#" : ".") +
                this.lights.slice(idx + 1)
        }
        this.pushes++
    }

    reset(): void {
        this.lights = ".".repeat(this.lightDiagram.length)
        this.pushes = 0
    }
}


function parseInput(input: string[]): Machine[] {
    const machines: Machine[] = []
    for (let line of input) {
        const lightDiagram: string = line.match(/\[(.*?)\]/g)![0].slice(1, -1)
        const buttons: number[][] = Array.from(line.matchAll(/\(([^\)]+)\)/g)!, match => match[1]!.split(",").map(Number))
        const requiredJoltage: number[] = line.match(/{([^}]*)}/g)![0].slice(1, -1).split(",").map(Number)
        machines.push(new Machine(lightDiagram, buttons, requiredJoltage))
    }
    return machines
}


function getPermutations(buttons: number[][], length: number): number[][][] {
    const results: number[][][] = []
    function permute(current: number[][], used: Set<number>) {
        if (current.length === length) {
            results.push([...current])
            return
        }
        for (let i = 0; i < buttons.length; i++) {
            if (!used.has(i)) {
                used.add(i)
                permute([...current, buttons[i]!], used)
                used.delete(i)
            }
        }
    }
    permute([], new Set())
    return results
}


function solvePart1(input: string[]): number {
    const machines: Machine[] = parseInput(input)
    for (const machine of machines) {
        let pushes: number = 1
        while (true) {
            const permutations: number[][][] = getPermutations(machine.buttons, pushes)
            for (const buttons of permutations) {
                buttons.map(button => machine.push(button))
                if (machine.lights === machine.lightDiagram) break
                machine.reset()
            }
            if (machine.lights === machine.lightDiagram) break
            pushes++
        }
    }
    return machines.reduce((sum, m) => sum + m.pushes, 0)
}


async function solvePart2(input: string[]): Promise<number> {
    const machines: Machine[] = parseInput(input)
    const { Context } = await init();
    const { Int, Optimize } = Context("main");
    let result = 0
    for (const machine of machines) {
        const buttons = machine.buttons
        const requirement = machine.requiredJoltage
        const solver = new Optimize()
        const variables = []
        for (let i = 0; i < buttons.length; i++) {
            const value = Int.const(String.fromCodePoint(i + 97))
            solver.add(value.ge(0))
            variables.push(value)
        }
        for (let i = 0; i < requirement.length; i++) {
            let condition = Int.val(0)
            for (let j = 0; j < buttons.length; j++) {
                if (buttons[j]!.includes(i)) condition = condition.add(variables[j]!)
            }
            condition = condition.eq(Int.val(requirement[i]!))
            solver.add(condition)
        }
        const sum = variables.reduce((a, x) => a.add(x), Int.val(0))
        solver.minimize(sum)
        if ((await solver.check()) == "sat") result += parseInt(solver.model().eval(sum).toString())
    }
    return result
}


export async function main(input: string[]): Promise<number[]> {
    const part1 = solvePart1(input)
    console.log(part1)
    const part2 = await solvePart2(input)
    console.log(part2)
    return [part1, part2]
}


if (import.meta.main) {
    const input = getDayInput(DAY, TEST)
    main(input)
}
