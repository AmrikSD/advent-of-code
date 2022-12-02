from utils.read_file import read_lines

from day01.part_1 import solver as solver_01
from day01.part_2 import solver as solver_02


def test_part_1():
    assert solver_01(read_lines("tests/day01/input.txt")) == 24000


def test_part_2():
    assert solver_02(read_lines("tests/day01/input.txt")) == 45000
