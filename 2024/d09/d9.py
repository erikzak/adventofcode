"""AOC day 9"""
import sys


class Disk:
    """Keeps track of disk properties."""
    def __init__(self, disk_map: str):
        self.disk_map = disk_map
        self.blocks: list[int | None] = []
        self.file_sizes: dict[int, int] = {}
        current_file_id = 0
        for i, block_size in enumerate(disk_map):
            block_size = int(block_size)
            if (i % 2 != 0):
                self.blocks.extend([None] * block_size)
            else:
                self.blocks.extend([current_file_id] * block_size)
                self.file_sizes[current_file_id] = block_size
                current_file_id += 1

    def __str__(self):
        return "".join(str(block) if (block is not None) else "." for block in self.blocks)

    def defragment(self):
        """Defragments files from end of disk to earliest free space."""
        free_space = self.blocks.index(None)
        for i, file_id in reversed(list(enumerate(self.blocks))):
            if (free_space > i):
                return
            if (file_id is None):
                continue
            self.blocks[free_space] = file_id
            self.blocks[i] = None
            free_space = self.blocks.index(None, free_space)

    def get_next_empty_block(self, start: int) -> int:
        """Returns the first empty block after the start index, or None if none found."""
        try:
            return self.blocks.index(None, start)
        except ValueError:
            return None

    def get_next_data_block(self, start: int) -> int:
        """Returns the first block with data after the start index, or None if none found."""
        for i, block in enumerate(self.blocks[start:], start):
            if (block is not None):
                return i
        return None

    def defragment_whole_files(self):
        """Defragments files from end of disk to earliest free space moving whole files."""
        i = len(self.blocks) - 1
        while (i > 0):
            file_id = self.blocks[i]
            if (file_id is None):
                i -= 1
                continue
            file_size = self.file_sizes[file_id]
            free_space_idx = 0
            while ((free_space_idx := self.get_next_empty_block(free_space_idx)) is not None):
                if (free_space_idx > i):
                    break
                next_data = self.get_next_data_block(free_space_idx)
                free_blocks = next_data - free_space_idx
                if (free_blocks >= file_size):
                    self.blocks[free_space_idx:free_space_idx + file_size] = [file_id] * file_size
                    self.blocks[i - file_size + 1:i + 1] = [None] * file_size
                    break
                free_space_idx = next_data + self.file_sizes[self.blocks[next_data]]
            i -= file_size

    @property
    def checksum(self):
        """Returns the checksum of the disk."""
        return sum(i * file_id for i, file_id in enumerate(self.blocks) if (file_id is not None))


def get_input(file_path: str) -> Disk:
    """Returns content from puzzle string file tailored for today's puzzle."""
    with open(file_path, "r", encoding="utf-8") as reader:
        return Disk(reader.read().strip())


def solve_part1(disk: Disk) -> int:
    """Solution part 1."""
    disk.defragment()
    return disk.checksum


def solve_part2(disk: Disk) -> int:
    """Solution part 2."""
    disk.defragment_whole_files()
    return disk.checksum


def main():
    """Runs puzzle solutions."""
    test = any(arg in ["-t", "--test"] for arg in sys.argv)
    input_file = "test.txt" if (test) else "input.txt"
    disk = get_input(input_file)
    print(solve_part1(disk))
    disk = get_input(input_file)
    print(solve_part2(disk))


if (__name__ == "__main__"):
    main()
