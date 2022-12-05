from typing import Iterator


def parse_table(input: Iterator[str], rows, columns):
    stacks = {index: [] for index in range(columns)}

    for _ in range(rows):
        row = (letter for letter in next(input)[1::4])

        for index, value in enumerate(row):
            if not value == " ":
                stacks[index].insert(0, value)

    return stacks


def _clean_values(line: str):
    return (
        line.replace("move ", "")
        .replace(" from ", ",")
        .replace(" to ", ",")
        .strip("\n")
        .split(",")
    )


def parse_instruction(line: str):
    return [int(value) for value in _clean_values(line)]
