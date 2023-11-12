type SolverResult = Promise<number>;
type SolverFunction = { (filePath: string): SolverResult };

interface DaySolution {
    part1: SolverFunction,
    part2: SolverFunction
}

interface Participants {
    [s: string]: DaySolution,
}

export { SolverResult, SolverFunction, DaySolution, Participants }