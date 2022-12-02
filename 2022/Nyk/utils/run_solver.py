from timeit import default_timer as timer

from .read_file import read_line, read_lines


def run_solver_lines(solver, file_name="input.txt"):
    start = timer()
    result = solver(read_lines(file_name))
    end = timer()

    print(f"Result: {result}")
    print("Time elapsed: {:.3f}s".format(end - start))


def run_solver_line(solver, file_name="input.txt"):
    start = timer()
    result = solver(next(read_line(file_name)))
    end = timer()

    print(f"Result: {result}")
    print("Time elapsed: {:.3f}s".format(end - start))
