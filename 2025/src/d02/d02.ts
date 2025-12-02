// Advent of Code day 2
import getDayInput from "../lib/input-parser.ts"


const DAY: string = "d02"
const TEST: boolean = false


interface Product {
    firstId: number
    secondId: number
    validate(part: number): number
    validateId(id: number): boolean
}


class Product implements Product {
    constructor(firstId: number, secondId: number) {
        this.firstId = firstId
        this.secondId = secondId
    }

    validate(part: number): number {
        // Validates the range of product IDs and returns the sum of invalid IDs
        let sum = 0
        for (let i = this.firstId; i <= this.secondId; i++) {
            const validator = part === 1 ? this.validateIdPart1 : this.validateIdPart2
            if (!validator(i)) {
                sum += i
            }
        }
        if (TEST) {
            console.log(`${this.firstId}-${this.secondId} += ${sum}`)
        }
        return sum
    }

    validateIdPart1(id: number): boolean {
        // Checks if id has repeating pattern of digits
        const idStr = id.toString()
        const patternLength = Math.floor(idStr.length / 2)
        const check = idStr.substring(0, patternLength)
        const compare = idStr.substring(patternLength)
        if (check === compare) {
            if (TEST) {
                console.log(`${id}: ${check}${compare}`)
            }
            return false
        }
        return true
    }

    validateIdPart2(id: number): boolean {
        // Checks if id has repeating pattern of digits
        const idStr = id.toString()
        const maxPatternLength = Math.floor(idStr.length / 2)
        for (let patternLength = 1; patternLength <= maxPatternLength; patternLength++) {
            for (let i = 0; i <= idStr.length - 2 * patternLength; i++) {
                const check = idStr.substring(i, i + patternLength)
                if (check.startsWith("0")) {
                    continue
                }
                for (let c = i + patternLength; c <= patternLength; c += patternLength) {
                    if (!isRepeating(idStr.substring(c), check)) {
                        continue
                    }
                    if (TEST) {
                        console.log(`${id}: ${check}`)
                    }
                    return false
                }
            }
        }
        return true
    }
}


const isRepeating = (str: string, pattern: string): boolean => {
    // Checks if str consists of repeating pattern
    const patternLength = pattern.length
    // console.log(`  checking ${pattern} against ${str}`)
    for (let i = 0; i < str.length; i += patternLength) {
        const check = str.substring(i, i + patternLength)
        if (check !== pattern) {
            return false
        }
    }
    return true
}


const solve = (input: string[], part: number): number => {
    // Solver for parts 1 & 2
    const products: Product[] = getProducts(input)
    return products.reduce((sum, product) => sum + product.validate(part), 0)
}


const getProducts = (input: string[]): Product[] => {
    const products: Product[] = (input[0] || "")
        .split(",")
        .map(productLine => {
            const [first, second] = productLine.split("-").map(Number)
            return new Product(first || 0, second || 0)
        })
    return products
}


export const main = (input: string[]): number[] => {
    const part1 = solve(input, 1)
    console.log(part1)
    const part2 = solve(input, 2)
    console.log(part2)
    return [part1, part2]
}


if (import.meta.main) {
    const input = getDayInput(DAY, TEST)
    main(input)
}
