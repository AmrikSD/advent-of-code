import { SolverFunction, FileReaders } from "../../utils/types"
import { readFile } from "../../utils/readFile"
import { benchmarkSolver } from "../../utils/runner";

import { resolve } from "path";

const solver: SolverFunction = async (fileReader: FileReaders, filePath: string) => {
    let value = Number(await fileReader(filePath))
    return value;
};

const inputPath = resolve(...["Day01", "input.txt"])

benchmarkSolver(solver, inputPath, readFile)

export { solver }