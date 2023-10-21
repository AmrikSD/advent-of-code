{{#participants}}
import { results as {{ name }} } from './{{name}}/solution'
{{/participants}}

import { benchmarkSolver } from "utils/runner";
import { DaySolution, Participants } from "types";

const results: Participants = {
    {{#participants}}
        {{ name }},
    {{/participants}}
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
