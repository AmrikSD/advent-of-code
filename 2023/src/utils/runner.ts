import { SolverFunction } from "../types";

const benchmarkSolver = async (solver: SolverFunction, filePath: string,) => {
    let startTime = performance.now();

    const result = await solver(filePath);
    let endTime = performance.now();

    return { result, timeTaken: `${endTime - startTime} ms` };
}

export { benchmarkSolver }