import { SolverFunction, FileReaders } from "../types";
import { readLines } from "./readFile";

const benchmarkSolver = async (solver: SolverFunction, filePath: string, fileReader: FileReaders = readLines) => {
    let startTime = performance.now();

    const result = await solver(fileReader, filePath);
    let endTime = performance.now();

    return { result, timeTaken: `${endTime - startTime} ms` };
}

export { benchmarkSolver }