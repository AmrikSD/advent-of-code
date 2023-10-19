import { benchmarkSolver } from "utils/runner";
import { daySolution as Nyk } from "./Nyk/solution";
import { DaySolution, Participants } from "types";
import { readFile } from "utils/readFile";

const results: Participants = {
    Nyk,
    Amrik: Nyk
}

for (let [participant, resultSet] of Object.entries<DaySolution>(results)) {
    console.log(participant)

    const { part1, part2, path } = resultSet

    console.log("Part 1: ", await benchmarkSolver(part1, path, readFile))
    console.log("Part 2: ", await benchmarkSolver(part2, path, readFile))

    console.log("") //newline
}