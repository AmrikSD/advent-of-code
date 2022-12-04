from utils.read_file import read_lines

from day03.part_1 import solver as solver_01
from day03.part_2 import solver as solver_02


def test_part_1():
    assert solver_01(read_lines("tests/day03/input.txt")) == 157


def test_part_2():
    assert solver_02(read_lines("tests/day03/input.txt")) == 70
