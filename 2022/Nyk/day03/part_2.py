from typing import Iterator

from utils import run_solver_lines

# 96 for lower
# 38 for upper


def solver(file_iter: Iterator[str]):
    lower_ord_adjustment = 96
    upper_ord_adjustment = 38
    sum_of_priority = 0

    for line in file_iter:
        elf_1, elf_2, elf_3 = line, next(file_iter), next(file_iter)

        shared_type = [
            item_type
            for item_type in elf_1
            if item_type in elf_2 and item_type in elf_3
        ][0]

        sum_of_priority += ord(shared_type) - (
            lower_ord_adjustment
            if shared_type.islower()
            else upper_ord_adjustment
        )

    return sum_of_priority


if __name__ == "__main__":
    run_solver_lines(solver, "day03/input.txt")
