from typing import Iterator

from utils import run_solver_lines

from .rps import existing_state


def solver(file_iter: Iterator[str]):
    total_points = 0

    for line in file_iter:
        opponent, game_state = line.strip("\n").split(" ")
        total_points += (
            6 if game_state == "Z" else 3 if game_state == "Y" else 0
        ) + existing_state[existing_state[opponent]["states"][game_state]][
            "score"
        ]

    return total_points


if __name__ == "__main__":
    run_solver_lines(solver, "day02/input.txt")
