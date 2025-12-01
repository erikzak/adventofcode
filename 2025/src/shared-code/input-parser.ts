// Reads input file for a given day and returns its content as a string, optionally for test data
import { readFileSync } from "fs"
import { EOL } from "os"


export default function getDayInput(folder: string, test: boolean = false): string[] {
    /**
     * Reads input file for a given day and returns its content as a string, optionally for test data
     * 
     * Parameters:
     *  folder: string - folder path where input files are located
     *  test: boolean - whether to read test input or actual input (default: false)
     * Returns:
     *  string - content of the input or test file
     */
    const filePath: string = `./src/${folder}/${test ? "test" : "input"}.txt`
    const content: string = readFileSync(filePath, "utf8")
    // Strip empty lines
    const lines: string[] = content.split(EOL).filter(line => line.trim() !== "")
    return lines
}
