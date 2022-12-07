from typing import Iterator


def is_command(line: str):
    return "$" in line


def is_move_dir_command(line: str):
    return "cd" in line


def is_list_dir_command(line: str):
    return "ls" in line


def is_move_up(line: str):
    return ".." in line


def parse_cd_command(line: str):
    _, _, new_dir = line.split()
    return new_dir


def is_directory(line: str):
    return "dir" in line


def is_file(line: str):
    return "dir" not in line


def parse_file(line: str):
    size, file_name = line.split()
    return int(size), file_name


def calculate_size(file_iter: Iterator[str]):
    directory = {"size": 0, "content": {}}

    while True:
        try:
            line = next(file_iter).strip("\n")
        except StopIteration:
            break

        if is_command(line):
            if is_move_dir_command(line):
                if is_move_up(line):
                    return directory

                new_dir = parse_cd_command(line)
                new_dict = calculate_size(file_iter)

                directory["content"][new_dir] = {
                    "size": new_dict["size"],
                    "content": new_dict["content"],
                }

                directory["size"] += directory["content"][new_dir]["size"]

        else:
            if is_file(line):
                size, file_name = parse_file(line)

                directory["size"] += size
                directory["content"][file_name] = size

    return directory
