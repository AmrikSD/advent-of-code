from glob import glob
from pathlib import Path
from shutil import copyfile
from sys import argv


def copy_files(copy_from_path: Path, copy_to_path: Path):
    files = glob("*.*", root_dir=copy_from_path)

    for file in files:
        copyfile(copy_from_path / file, copy_to_path / file)


def copy_solution(copy_from_path: Path, copy_to_path: Path):
    if not copy_to_path.exists():
        copy_to_path.mkdir()
    else:
        raise Exception("Solution folder already exists")

    copy_files(copy_from_path / "solutions", copy_to_path)


def copy_tests(copy_from_path: Path, copy_to_path: Path):
    if not copy_to_path.exists():
        copy_to_path.mkdir()
    else:
        raise Exception("Test folder already exists")

    copy_files(copy_from_path / "tests", copy_to_path)


def copy_template(day: str):
    from_path = Path("./utils/generate/template").absolute()
    to_path_solution = Path(f"./day{day}/").absolute()
    to_path_tests = Path(f"./tests/day{day}/").absolute()

    copy_solution(from_path, to_path_solution)
    copy_tests(from_path, to_path_tests)


if __name__ == "__main__":
    day = argv[1]

    copy_template(day)
