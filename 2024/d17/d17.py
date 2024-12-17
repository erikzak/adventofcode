"""AOC day 17"""
from __future__ import annotations
import sys
from itertools import repeat
from multiprocessing.dummy import Pool as ThreadPool


class Computer:
    """Computer stuff."""
    def __init__(self, registers: dict[str, int], program: str, parsed_program: list[int] = None):
        self.registers = registers
        self.program = program
        self._program: list[int] = parsed_program or [int(v) for v in program.split(",")]
        self.pointer = 0
        self.operations = {
            0: self.adv,
            1: self.bxl,
            2: self.bst,
            3: self.jnz,
            4: self.bxc,
            5: self.out,
            6: self.bdv,
            7: self.cdv,
        }
        self.output = []

    def __repr__(self) -> str:
        return f"Computer({self.registers}, {self.program})"

    def copy(self) -> Computer:
        return Computer(self.registers, self.program, self._program)

    def run(self) -> list[int]:
        """Runs the program and returns the output."""
        self.output = []
        while (self.pointer < len(self._program)):
            opcode = self._program[self.pointer]
            operand = self._program[self.pointer + 1]
            # print(f"{self.registers}, {self._program}, {self.pointer}: ({opcode}, {operand})")
            self.operations[opcode](operand)
            self.pointer += 2
        return ",".join(str(v) for v in self.output)

    def reset(self, a_value: int) -> None:
        """Restarts computer."""
        self.pointer = 0
        self.registers["A"] = a_value
        self.registers["B"] = 0
        self.registers["C"] = 0

    def combo(self, operand: int) -> int:
        """Returns combo operand."""
        if (0 <= operand <= 3):
            return operand
        if (operand == 4):
            return self.registers["A"]
        if (operand == 5):
            return self.registers["B"]
        if (operand == 6):
            return self.registers["C"]

    def adv(self, operand: int) -> None:
        """opcode 0"""
        self.registers["A"] = self.registers["A"] // 2 ** self.combo(operand)

    def bxl(self, operand: int) -> None:
        """opcode 1"""
        self.registers["B"] = self.registers["B"] ^ operand

    def bst(self, operand: int) -> None:
        """opcode 2"""
        self.registers["B"] = self.combo(operand) % 8

    def jnz(self, operand: int) -> None:
        """opcode 3"""
        if (self.registers["A"] == 0):
            return
        self.pointer = operand - 2

    def bxc(self, _: int) -> None:
        """opcode 4"""
        self.registers["B"] = self.registers["B"] ^ self.registers["C"]

    def out(self, operand: int) -> None:
        """opcode 5"""
        self.output.append(self.combo(operand) % 8)

    def bdv(self, operand: int) -> None:
        """opcode 6"""
        self.registers["B"] = self.registers["A"] // 2 ** self.combo(operand)

    def cdv(self, operand: int) -> None:
        """opcode 7"""
        self.registers["C"] = self.registers["A"] // 2 ** self.combo(operand)


def get_input(file_path: str) -> Computer:
    """Returns content from puzzle string file tailored for today's puzzle."""
    with open(file_path, "r", encoding="utf-8") as reader:
        content = reader.read()
    registers, program = content.split("\n\n")
    registers = [int(r.split(":")[-1].strip()) for r in registers.split("\n")]
    registers = {"A": registers[0], "B": registers[1], "C": registers[2]}
    program = program.split(":")[-1].strip()
    return Computer(registers, program)


def solve_part1(computer: Computer) -> int:
    """Solution part 1."""
    return computer.run()


def brute_threading(computer: Computer, a_value: int) -> str:
    """Multithreading brute force part 2."""
    computer = computer.copy()
    computer.reset(a_value)
    return computer.run()


def brute_part2(computer: Computer) -> int:
    """Solution part 2."""
    init = 267265166222230
    batch = 1000
    while (True):
        pool = ThreadPool(12)
        a_range = range(init, init + batch)
        results = pool.starmap(brute_threading, zip(repeat(computer), a_range))
        pool.close()
        pool.join()
        for i, output in enumerate(results):
            if (output == computer.program):
                return init + i
        init += batch


def finesse_part2(computer: Computer) -> int:
    """What is this I don't even"""
    g = computer._program
    global answer

    def solve(p, r):
        a: int
        b: int
        c: int
        if (p < 0):
            global answer
            answer = r
            return True
        for d in range(8):
            a, i = r << 3 | d, 0
            while i < len(g):
                if g[i+1] <= 3:
                    o = g[i + 1]
                elif g[i+1] == 4:
                    o = a
                elif g[i+1] == 5:
                    o = b
                elif g[i+1] == 6:
                    o = c
                if g[i] == 0:
                    a >>= o
                elif g[i] == 1:
                    b ^= g[i + 1]
                elif g[i] == 2:
                    b = o & 7
                elif g[i] == 3:
                    i = g[i + 1] - 2 if a != 0 else i
                elif g[i] == 4:
                    b ^= c
                elif g[i] == 5:
                    w = o & 7
                    break
                elif g[i] == 6:
                    b = a >> o
                elif g[i] == 7:
                    c = a >> o
                i += 2
            if w == g[p] and solve(p - 1, r << 3 | d):
                return True
        return False

    solve(len(computer._program) - 1, 0)
    return answer


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    puzzle_input = get_input(input_file)
    print(solve_part1(puzzle_input))
    print(finesse_part2(puzzle_input))


if (__name__ == "__main__"):
    main()
