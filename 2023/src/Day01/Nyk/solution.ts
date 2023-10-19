import { SolverFunction, FileReaders, DaySolution } from "../../types"

const part1: SolverFunction = async (fileReader: FileReaders, filePath: string) => {
    let value = Number(await fileReader(filePath))
    return value;
};

const part2: SolverFunction = async (fileReader: FileReaders, filePath: string) => {
    let value = Number(await fileReader(filePath))
    return value;
};

const path: string = "src/Day01/Nyk/input.txt"

const daySolution: DaySolution = {
    part1,
    part2,
    path

}

export { daySolution, }