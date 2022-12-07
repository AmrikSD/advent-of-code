from typing import Iterator

from utils import run_solver_lines

from .directories import calculate_size


def calculate_sized_dirs(directories: dict):
    found_sizes = []

    for key, value in directories.items():
        if key == "size":
            if value <= 100000:
                found_sizes.append(value)

        elif key == "content":
            found_sizes.extend(calculate_sized_dirs(value))

        elif type(value) is dict:
            found_sizes.extend(calculate_sized_dirs(value))

    return found_sizes


def solver(file_iter: Iterator[str]):
    next(file_iter)

    directories = {"/": calculate_size(file_iter)}

    return sum(calculate_sized_dirs(directories["/"]))


if __name__ == "__main__":
    run_solver_lines(solver, "day07/input.txt")
