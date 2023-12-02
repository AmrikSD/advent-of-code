import { SolverFunction } from "../types";

const benchmarkSolver = async (solver: SolverFunction, filePath: string,) => {
    const startTime = performance.now();

    const result = await solver(filePath);
    const endTime = performance.now();

    return { result, timeTaken: `${(endTime - startTime).toFixed(2)} ms` };
};

export { benchmarkSolver };
