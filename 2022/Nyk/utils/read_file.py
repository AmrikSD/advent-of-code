from pathlib import Path


def read_lines(file_name: str):
    with open(Path(file_name).absolute(), "r") as file:
        for line in file.readlines():
            yield line


def read_line(file_name: str):
    with open(Path(file_name).absolute(), "r") as file:
        yield file.readline()
