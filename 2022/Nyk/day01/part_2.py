from utils import read_lines

from .calories import get_calories


def solver(file_iter):
    return sum(sorted(list(get_calories(file_iter).values()))[-3:])


if __name__ == "__main__":
    print(solver(read_lines("day01/input.txt")))
    # print(solver(next(read_line('input.txt'))))
