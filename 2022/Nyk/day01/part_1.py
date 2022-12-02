from utils import run_solver_lines

from .calories import get_calories


def solver(file_iter):
    return max(
        get_calories(file_iter).values(),
    )


if __name__ == "__main__":
    run_solver_lines(solver, "day01/input.txt")
