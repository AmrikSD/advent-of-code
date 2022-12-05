from pprint import pprint
from typing import Iterator

from utils import run_solver_lines

from .stacks import parse_instruction, parse_table


def solver(file_iter: Iterator[str], rows, columns):
    table = parse_table(file_iter, rows, columns)

    next(file_iter)
    next(file_iter)  # skip unneeded lines

    for line in file_iter:
        instructions = parse_instruction(line)

        for _ in range(instructions[0]):
            table[instructions[2] - 1].append(table[instructions[1] - 1].pop())

    return "".join(stack[-1] for stack in table.values())


if __name__ == "__main__":
    run_solver_lines(solver, "day05/input.txt", {"rows": 8, "columns": 9})
