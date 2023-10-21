import { readFile, readLines } from "./utils/readFile";

type FileReaders = typeof readFile | typeof readLines;

type SolverResult = Promise<number>;
type SolverFunction = { (filePath: string): SolverResult };

interface DaySolution {
    part1: SolverFunction,
    part2: SolverFunction
}

interface Participants {
    [s: string]: DaySolution,
}

export { FileReaders, SolverResult, SolverFunction, DaySolution, Participants }