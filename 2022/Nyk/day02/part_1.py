from typing import Iterator

from utils import run_solver_lines

from .rps import winning_states


def solver(file_iter: Iterator[str]):
    total_points = 0

    for line in file_iter:
        opponent, us = line.strip("\n").split(" ")
        total_points += (
            6
            if winning_states[us]["winning_state"] == opponent
            else 3
            if winning_states[us]["equal_state"] == opponent
            else 0
        ) + winning_states[us]["score"]

    return total_points


if __name__ == "__main__":
    run_solver_lines(solver, "day02/input.txt")
