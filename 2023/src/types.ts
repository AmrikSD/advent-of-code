import { readFile, readLines } from "./utils/readFile";

type FileReaders = typeof readFile | typeof readLines;

type SolverResult = Promise<number>;
type SolverFunction = { (fileReader: FileReaders, filePath: string): SolverResult };

interface DaySolution {
    part1: SolverFunction,
    part2: SolverFunction
    path: string
}

interface Participants {
    [s: string]: DaySolution,
}

export { FileReaders, SolverResult, SolverFunction, DaySolution, Participants }