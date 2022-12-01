from utils import read_lines

from .calories import get_calories


def solver(file_iter):
    return max(
        get_calories(file_iter).values(),
    )


if __name__ == "__main__":
    print(solver(read_lines("day01/input.txt")))
    # print(solver(next(read_line('input.txt'))))
