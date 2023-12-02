import { SolverFunction, DaySolution } from "../../types";
import { readLines } from "utils/readFile";

const part1: SolverFunction = async (filePath: string) => {
    const limits: { [index: string]: number; } = { "red": 12, "green": 13, "blue": 14 };
    let result = 0;

    for await (const line of readLines(filePath)) {
        const [gameAndId, gameSets] = line.split(":");

        const gameId = parseInt(gameAndId.split(" ")[1]);

        let gamePossible = true;
        for (const gameSet of gameSets.split(";")) {
            for (const game of gameSet.split(",")) {
                const [, amount, colour] = game.split(" ");

                if (parseInt(amount) > limits[colour]) {
                    gamePossible = false;
                    break;
                }
            }
            if (!gamePossible) { break; }
        }

        if (gamePossible) {
            result += gameId;
        }
    }

    return result;
};

const part2: SolverFunction = async (filePath: string) => {
    const value = Number(await readLines(filePath));
    return value;
};

const results: DaySolution = {
    part1,
    part2,
};

export { results };