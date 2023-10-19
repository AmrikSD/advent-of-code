import { readFile, readLines } from "./readFile";

type FileReaders = typeof readFile | typeof readLines;

type SolverResult = Promise<number>;
type SolverFunction = { (fileReader: FileReaders, filePath: string): SolverResult };

export { FileReaders, SolverResult, SolverFunction }