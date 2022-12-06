from utils.read_file import read_line

from day06.part_1 import solver as solver_01
from day06.part_2 import solver as solver_02


def test_part_1():
    assert solver_01(next(read_line("tests/day06/input.txt"))) == 7


def test_part_2():
    assert solver_02(next(read_line("tests/day06/input.txt"))) == 19
