# Advent of Code 2025
In TypeScript. I don't actually know the best way to set up a TS module, but this worked.

## Setup
1. `npm init` and installed packages using `npm install`
2. `typescript tsc --init` and added relevant updates to tsconfig.ts

Added days as .ts modules with main functions logging/returning answers, then
run using e.g. `npx ts-node src/d01/d01.ts`

## Usage
1. Clone repo: `git clone https://github.com/erikzak/adventofcode.git`
2. Enter 2025 folder: `cd adventofcode/2025`
3. Install dependencies: `npm install`
4. Run individual days: `npx ts-node src/d01/d01.ts`


## Interactive Calendar
I had Copilot build a react frontend on its own, using Claude, with clickable
doors for AoC days.

1. Run `npm run dev` to start the development server
2. Open your browser to the URL shown (typically http://localhost:5173)
3. Click on any calendar door (1-12) to open it
4. The solution will run on the test input for that day
5. View the results for Part 1, Part 2, and execution time
