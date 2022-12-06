from typing import Iterator

from utils import run_solver_line


def solver(file_iter: str):
    for index in range(3, len(file_iter)):
        if len(set(file_iter[index - 4 : index])) == 4:
            return index

    return 0


if __name__ == "__main__":
    run_solver_line(solver, "day06/input.txt")
