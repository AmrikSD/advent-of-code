import { SolverFunction, FileReaders } from "../types";
import { readLines } from "./readFile";

const benchmarkSolver = (solver: SolverFunction, filePath: string, fileReader: FileReaders = readLines) => {
    let startTime = performance.now();

    solver(fileReader, filePath).then(result => {
        let endTime = performance.now()

        console.log({ result, timeTaken: `${endTime - startTime} ms` })
    });

}

export { benchmarkSolver }