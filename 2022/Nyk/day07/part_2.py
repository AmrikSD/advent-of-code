from typing import Iterator

from utils import run_solver_lines

from .directories import calculate_size


def calculate_sized_dirs(directories: dict, free_space_needed):
    found_sizes = []

    for key, value in directories.items():
        if key == "size":
            if value >= free_space_needed:
                found_sizes.append(value)

        elif key == "content":
            found_sizes.extend(calculate_sized_dirs(value, free_space_needed))

        elif type(value) is dict:
            found_sizes.extend(calculate_sized_dirs(value, free_space_needed))

    return found_sizes


def solver(file_iter: Iterator[str]):
    next(file_iter)

    directories = {"/": calculate_size(file_iter)}
    free_space_needed = 30000000 - (70000000 - directories["/"]["size"])

    return min(calculate_sized_dirs(directories["/"], free_space_needed))


if __name__ == "__main__":
    run_solver_lines(solver, "day07/input.txt")
