import { SolverFunction, DaySolution } from "../../types"
import { readFile } from "utils/readFile";

const part1: SolverFunction = async (filePath: string) => {
    const value = Number(await readFile(filePath))
    return value;
};

const part2: SolverFunction = async (filePath: string) => {
    const value = Number(await readFile(filePath))
    return value;
};

const results: DaySolution = {
    part1,
    part2,
}

export { results }