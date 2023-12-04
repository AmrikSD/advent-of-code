import { SolverFunction, DaySolution } from "../../types";
import { readLines, readFile } from "utils/readFile";

const splitNumberStrings = (line: string) => {
    const numbers = line.split(":");
    const [winningNumbers, playersNumbers] = numbers[1].split(" | ");

    const winning = winningNumbers.split(" ").filter(value => value !== "").map(value => parseInt(value));
    const players = playersNumbers.split(" ").filter(value => value !== "").map(value => parseInt(value));

    return [winning, players];
};

const part1: SolverFunction = async (filePath: string) => {
    let result = 0;

    for await (const line of readLines(filePath)) {
        const [winning, players] = splitNumberStrings(line);

        result += players
            .reduce<number>((result, value) => winning.includes(value) ? result === 0 ? 1 : result * 2 : result, 0);
    };

    return result;
};

const part2: SolverFunction = async (filePath: string) => {
    return (await readFile(filePath))
        .split("\n")
        .map(line => ({ line, instances: 1 }))
        .reduce<number>((result, value, index, lines) => {
            const [winning, players] = splitNumberStrings(value.line);

            const matchingNumbers = players.reduce<number>((result, value) => winning.includes(value) ? result + 1 : result, 0);
            for (let lineIndex = index + 1; lineIndex < index + matchingNumbers + 1; lineIndex++) {

                lines[lineIndex].instances += value.instances;
            }

            return result + value.instances;;
        }, 0);
};

const results: DaySolution = {
    part1,
    part2,
};

export { results };