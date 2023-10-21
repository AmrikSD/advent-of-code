import { results as Amrik } from './Amrik/solution'
import { results as Binda } from './Binda/solution'
import { results as Nyk } from "./Nyk/solution";

import { benchmarkSolver } from "utils/runner";
import { DaySolution, Participants } from "types";

import { readFile } from "utils/readFile";

const results: Participants = {
    Amrik,
    Binda,
    Nyk,
}

export const runSolutions = async (day: string) => {
    for (let [participant, resultSet] of Object.entries<DaySolution>(results)) {
        console.log(participant)

        const { part1, part2 } = resultSet

        const inputPath = `src/${day}/${participant}/input.txt`

        console.log("Part 1: ", await benchmarkSolver(part1, inputPath))
        console.log("Part 2: ", await benchmarkSolver(part2, inputPath))

        console.log("") //newline
    }
}