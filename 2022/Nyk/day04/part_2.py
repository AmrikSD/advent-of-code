from typing import Iterator

from utils import run_solver_lines

from .elf_sections import get_elf_sections


def solver(file_iter: Iterator[str]):
    non_overlap_pairs = 0
    lines = 0

    for line in file_iter:
        lines += 1

        elf_1_1, elf_1_2, elf_2_1, elf_2_2 = get_elf_sections(line)

        if elf_1_2 < elf_2_1 or elf_1_1 > elf_2_2:
            non_overlap_pairs += 1

    return lines - non_overlap_pairs


if __name__ == "__main__":
    run_solver_lines(solver, "day04/input.txt")
