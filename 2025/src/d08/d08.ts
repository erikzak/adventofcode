// Advent of Code day 8
import getDayInput from "../lib/input-parser.ts"


const DAY: string = "d08"
const TEST: boolean = false


interface Position {
    x: number
    y: number
    z: number
}


interface Distance {
    from: JunctionBox
    to: JunctionBox
    length: number
}


interface IJunctionBox {
    id: number
    position: Position
    distances: Distance[]
    circuit: Circuit | null
    toString(): string
}


class JunctionBox implements IJunctionBox {
    id: number
    position: Position
    distances: Distance[]
    circuit: Circuit | null

    constructor(id: number, position: Position) {
        this.id = id
        this.position = position
        this.distances = []
        this.circuit = null
    }

    toString(): string {
        return this.positionString()
            + `[${this.distances[0]!.length.toFixed(2)} > ${this.distances[0]!.to.id}]`
    }

    positionString(): string {
        return `(${this.id}: ${this.position.x},${this.position.y},${this.position.z})`
    }
}


interface ICircuit {
    id: number
    junctionBoxes: JunctionBox[]
    count: number
    toString(): string
}


class Circuit implements ICircuit {
    id: number
    junctionBoxes: JunctionBox[]
    count: number

    constructor(id: number, junctionBoxes: JunctionBox[]) {
        this.id = id
        this.junctionBoxes = junctionBoxes
        this.count = junctionBoxes.length
    }

    add(junctionBox: JunctionBox): void {
        this.junctionBoxes.push(junctionBox)
        this.count++
    }

    toString(): string {
        const boxStrings = this.junctionBoxes.map(
            jb => `(${jb.id}: ${jb.position.x},${jb.position.y},${jb.position.z})`
        )
        return "[" + boxStrings.join(", ") + "]"
    }
}


function distanceBetween(jba: JunctionBox, jbb: JunctionBox): number {
    const dx = jba.position.x - jbb.position.x
    const dy = jba.position.y - jbb.position.y
    const dz = jba.position.z - jbb.position.z
    return Math.sqrt(dx * dx + dy * dy + dz * dz)
}


function solvePart1(input: string[]): number {
    const junctionBoxes: JunctionBox[] = parseInput(input)
    if (TEST) {
        junctionBoxes.sort((a, b) => (a.distances[0]!.length - b.distances[0]!.length))
        console.log(junctionBoxes.map(jb => jb.toString()))
    }
    // Connect N junction boxes into circuits, sorted by shortest distance
    const maxConnections = TEST ? 10 : 1000
    const circuits: Circuit[] = connectBoxes(junctionBoxes, maxConnections) as Circuit[]
    // Find the three largest circuits and multiply their sizes
    circuits.sort((a, b) => b.junctionBoxes.length - a.junctionBoxes.length)
    if (TEST) {
        console.log(
            circuits.length + junctionBoxes.reduce((sum, jb) => sum + (jb.circuit ? 0 : 1), 0),
            circuits.map(c => c.toString())
        )
    }
    return circuits.slice(0, 3).reduce((sum, circuit) => sum * circuit.count, 1)
}


function solvePart2(input: string[]): number {
    const junctionBoxes: JunctionBox[] = parseInput(input)
    const [fromBox, toBox]: [JunctionBox, JunctionBox] = connectBoxes(junctionBoxes, Infinity) as [JunctionBox, JunctionBox]
    if (TEST) {
        console.log(`Connected all boxes by linking ${fromBox.positionString()} & ${toBox.positionString()}`)
    }
    return fromBox.position.x * toBox.position.x
}


function connectBoxes(junctionBoxes: JunctionBox[], maxConnections: number): Circuit[] | [JunctionBox, JunctionBox] {
    const circuits: Circuit[] = []
    let connections: number = 0
    while (connections < maxConnections) {
        junctionBoxes.sort((a, b) => (a.distances[0]!.length - b.distances[0]!.length))
        const nextConnection: Distance = junctionBoxes[0]!.distances.shift()!
        const [fromBox, toBox]: [JunctionBox, JunctionBox] = [nextConnection.from, nextConnection.to]
        toBox.distances = toBox.distances.filter(d => d.to.id !== fromBox.id)
        connections++
        if (fromBox.circuit && toBox.circuit && fromBox.circuit.id === toBox.circuit.id) {
            if (TEST) {
                console.log(`Skipping link between same circuit: ${fromBox.positionString()} & ${toBox.positionString()}`)
            }
            continue
        }
        if (!fromBox.circuit && !toBox.circuit) {
            const circuit: Circuit = new Circuit(
                Math.max(...circuits.map(c => c.id), -1) + 1,
                [fromBox, toBox]
            )
            fromBox.circuit = circuit
            toBox.circuit = circuit
            circuits.push(circuit)
            if (TEST) {
                console.log(`New circuit: ${circuit.toString()}`)
            }
        } else if (fromBox.circuit && toBox.circuit) {
            const fromCircuit: Circuit = circuits.splice(circuits.indexOf(fromBox.circuit), 1)[0]!
            const toCircuit: Circuit = circuits.splice(circuits.indexOf(toBox.circuit), 1)[0]!
            const newCircuit: Circuit = new Circuit(
                Math.max(...circuits.map(c => c.id), -1) + 1,
                [...fromCircuit.junctionBoxes, ...toCircuit.junctionBoxes]
            )
            for (const junctionBox of newCircuit.junctionBoxes) {
                junctionBox.circuit = newCircuit
            }
            if (TEST) {
                console.log(`Combining circuits ${fromCircuit.toString()} & ${toCircuit.toString()}: ${newCircuit.toString()}`)
            }
            circuits.push(newCircuit)
        } else if (fromBox.circuit) {
            fromBox.circuit.add(toBox)
            toBox.circuit = fromBox.circuit
            if (TEST) {
                console.log(`Adding ${toBox.id} to circuit: ${fromBox.circuit.toString()}`)
            }
        } else {
            toBox.circuit!.add(fromBox)
            fromBox.circuit = toBox.circuit!
            if (TEST) {
                console.log(`Adding ${fromBox.id} to circuit: ${toBox.circuit!.toString()}`)
            }
        }
        if (circuits[0]!.count == junctionBoxes.length) {
            return [fromBox, toBox]
        }
    }
    return circuits
}


function parseInput(input: string[]): JunctionBox[] {
    // Parse input into junction boxes
    const junctionBoxes: JunctionBox[] = input.map((line, index) => {
        const [x, y, z] = line.split(",").map(Number) as [number, number, number]
        return new JunctionBox(index, { x, y, z })
    })
    // Find distances between all junction boxes
    for (const junctionBox of junctionBoxes) {
        junctionBox.distances = []
        for (const otherBox of junctionBoxes) {
            if (junctionBox.id === otherBox.id) {
                continue
            }
            const dist = distanceBetween(junctionBox, otherBox)
            junctionBox.distances.push({
                from: junctionBox,
                to: otherBox,
                length: dist,
            })
        }
        junctionBox.distances.sort((a, b) => a.length - b.length)
    }
    return junctionBoxes
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
