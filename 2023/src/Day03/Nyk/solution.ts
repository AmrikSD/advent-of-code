import { SolverFunction, DaySolution } from "../../types";
import { readLines } from "utils/readFile";

const clamp = (num: number, min: number, max: number) => Math.min(Math.max(num, min), max);

const getAdjacentIndexes = (numIndexes: number[], lineIndex: number, schematicLength: number, lineLength: number) => {
    const minHorizontalIndex = clamp(numIndexes.reduce<number>((result, value) => Math.min(result, value), Number.MAX_SAFE_INTEGER) - 1, 0, lineLength);
    const maxHorizontalIndex = clamp(numIndexes.reduce<number>((result, value) => Math.max(result, value), Number.MIN_SAFE_INTEGER) + 1, 0, lineLength);

    const minVerticalIndex = clamp(lineIndex - 1, 0, schematicLength);
    const maxVerticalIndex = clamp(lineIndex + 1, 0, schematicLength);

    const indexes: number[][] = [];
    for (let verticalIndex = minVerticalIndex; verticalIndex < maxVerticalIndex + 1; verticalIndex++) {
        for (let horizontalIndex = minHorizontalIndex; horizontalIndex < maxHorizontalIndex + 1; horizontalIndex++) {
            indexes.push([verticalIndex, horizontalIndex]);
        }
    }

    return indexes;
};

const isValidNumber = (indexesToCheck: number[][], schematic: string[][]): boolean => {
    const notANumber = (value: string): boolean => isNaN(parseInt(value));
    const notADot = (value: string): boolean => value !== ".";

    return indexesToCheck
        .map(([vertical, horizontal]) => schematic[vertical][horizontal])
        .some(([value]) => notANumber(value) && notADot(value));
};

const part1: SolverFunction = async (filePath: string) => {
    const schematic: string[][] = [];
    for await (const line of readLines(filePath)) {
        schematic.push(line.split(""));
    }

    const re = /[0-9]{1,3}/g;
    let result = 0;

    for (let lineIndex = 0; lineIndex < schematic.length; lineIndex++) {
        const line = schematic[lineIndex].join("");

        let reResult;
        while ((reResult = re.exec(line)) !== null) {
            const indexes = [reResult.index, reResult.index + reResult[0].length - 1];
            const indexesToCheck = getAdjacentIndexes(indexes, lineIndex, schematic.length - 1, schematic[lineIndex].length - 1);

            if (isValidNumber(indexesToCheck, schematic)) {
                const value = parseInt(reResult[0]);
                result += value;
            }
        }

    };

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