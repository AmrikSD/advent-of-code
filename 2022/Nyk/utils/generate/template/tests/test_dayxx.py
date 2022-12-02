from dayxx.part_1 import solver as solver_01
from dayxx.part_2 import solver as solver_02
from utils.read_file import read_lines


def test_part_1():
    assert solver_01(read_lines("tests/dayxx/input.txt")) == "Part 1 Done"


def test_part_2():
    assert solver_02(read_lines("tests/dayxx/input.txt")) == "Part 2 Done"
