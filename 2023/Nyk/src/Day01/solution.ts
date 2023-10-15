import { SolverFunction, FileReaders } from "../types"
import { readFile } from "../utils/readFile"
import { benchmarkSolver } from "../utils/runner";

const solver: SolverFunction = async (fileReader: FileReaders, filePath: string) => {
    let value = Number(await fileReader(filePath))
    return value;
};

//benchmarkSolver(solver, "input.txt", readFile)

export { solver }