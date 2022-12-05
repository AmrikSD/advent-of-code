from utils.read_file import read_lines

from day05.part_1 import solver as solver_01
from day05.part_2 import solver as solver_02


def test_part_1():
    assert solver_01(read_lines("tests/day05/input.txt"), 3, 3) == "CMZ"


def test_part_2():
    assert solver_02(read_lines("tests/day05/input.txt"), 3, 3) == "MCD"
