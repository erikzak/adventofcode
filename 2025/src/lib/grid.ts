// Module for map shenanigans


export interface Coordinate {
    x: number
    y: number
}


export interface Node {
    key: string
    coords: Coordinate
    value: string | number
}


export interface Grid {
    nodes: Map<string, Node>
    width: number
    height: number
    rows: Node[][]
    columns: Node[][]
    getNode(x: number, y: number): Node | undefined
    toString(): string
}


export class Grid implements Grid {
    constructor(input: string[]) {
        const rows: Node[][] = input.slice().reverse().map((line, y) =>
            line.split("").map((char, x) => ({
                key: `${x},${y}`,
                coords: { x, y },
                value: char
            }))
        )
        this.nodes = new Map(rows.flat().map(node => [node.key, node]))
        this.width = input[0]!.length
        this.height = input.length
        this.rows = rows.reverse()
        this.columns = rows[0]!.map((_, colIndex) => rows.map(row => row[colIndex]).reverse()) as Node[][]
    }

    toString(): string {
        let map = ""
        for (let y = this.height - 1; y >= 0; y--) {
            for (let x = 0; x < this.width; x++) {
                map += this.getNode(x, y)?.value ?? " "
            }
            map += "\n"
        }
        return map
    }

    getNode(x: number, y: number): Node | undefined {
        return this.nodes.get(`${x},${y}`)
    }
}
