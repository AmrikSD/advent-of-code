from utils.read_file import read_lines

from day04.part_1 import solver as solver_01
from day04.part_2 import solver as solver_02


def test_part_1():
    assert solver_01(read_lines("tests/day04/input.txt")) == 2


def test_part_2():
    assert solver_02(read_lines("tests/day04/input.txt")) == 4
