import { SolverFunction, DaySolution } from "../../types";
import { readLines } from "utils/readFile";

const clamp = (num: number, min: number, max: number) => Math.min(Math.max(num, min), max);

const getAdjacentIndexes = (currentNumberBuilder: NumberBuilder, lineIndex: number, schematicLength: number, lineLength: number) => {
    const numberIndexes = numberBuilderValues(currentNumberBuilder);

    const minHorizontalIndex = clamp(numberIndexes.reduce<number>((result, value) => Math.min(result, value), Number.MAX_SAFE_INTEGER) - 1, 0, lineLength);
    const maxHorizontalIndex = clamp(numberIndexes.reduce<number>((result, value) => Math.max(result, value), Number.MIN_SAFE_INTEGER) + 1, 0, lineLength);

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

type NumberBuilder = (string | number)[][];
const numberBuilderKeys = (arr: NumberBuilder): string[] => arr.reduce<string[]>((result, value) => { result.push(value[0] as string); return result; }, []);
const numberBuilderValues = (arr: NumberBuilder): number[] => arr.reduce<number[]>((result, value) => { result.push(value[1] as number); return result; }, []);

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

    let result = 0;
    for (let lineIndex = 0; lineIndex < schematic.length; lineIndex++) {
        let currentNumberBuilder: NumberBuilder = [];


        for (let charIndex = 0; charIndex < schematic[lineIndex].length; charIndex++) {
            const char = schematic[lineIndex][charIndex];

            if (!isNaN(parseInt(char))) //Is a number
            {
                currentNumberBuilder.push([char, charIndex]);
                continue;
            }

            const keys = numberBuilderKeys(currentNumberBuilder);
            if (keys.length !== 0 && isNaN(parseInt(char))) {
                const indexesToCheck = getAdjacentIndexes(currentNumberBuilder, lineIndex, schematic.length - 1, schematic[lineIndex].length - 1);

                if (isValidNumber(indexesToCheck, schematic)) {
                    const value = parseInt(keys.join(""));
                    if (value == 755) {
                        console.log("Adding ", value);
                    }
                    result += value;
                }

                currentNumberBuilder = [];
            }
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