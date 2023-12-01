import { SolverFunction, DaySolution } from "../../types";
import { readLines } from "utils/readFile";

const part1: SolverFunction = async (filePath: string) => {
    let result = 0;
    for await (const line of readLines(filePath)) {
        let lineResult = "";

        for (const char of line) {
            if (isNaN(parseInt(char))) continue;

            lineResult += char;
            break;
        }

        for (const char of line.split("").reverse().join()) {
            if (isNaN(parseInt(char))) continue;

            lineResult += char;
            break;
        }

        result += parseInt(lineResult);
    }

    return result;
};


const getMinMaxOccurrence = (digit: string, line: string): [number, number] => {
    let endOfLine = false;
    const occurences: number[] = [];
    let nextIndex = -1;

    while (!endOfLine) {
        const index = line.indexOf(digit, nextIndex + 1);

        if (index == -1) {
            endOfLine = true;
            continue;
        }

        occurences.push(index);
        nextIndex = index;
    }

    return [
        occurences.at(0) ?? Number.MAX_SAFE_INTEGER,
        occurences.at(-1) ?? Number.MIN_SAFE_INTEGER
    ];

};

const convertToStringLiteral = (value: string): string => {
    const map: { [index: string]: string; } = {
        one: "1",
        two: "2",
        three: "3",
        four: "4",
        five: "5",
        six: "6",
        seven: "7",
        eight: "8",
        nine: "9",
    };

    return map[value] ?? value;
};

const getNumValue = (line: string): number => {
    const valuesToFind: string[] = [
        "one",
        "two",
        "three",
        "four",
        "five",
        "six",
        "seven",
        "eight",
        "nine",
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7",
        "8",
        "9",
    ];

    const occurences = valuesToFind.reduce<Record<string, Record<string, number>>>((accumulator, digit) => {
        const [min, max] = getMinMaxOccurrence(digit, line);

        accumulator[digit] = { min, max };
        return accumulator;
    }, {});

    let minIndex = Number.MAX_SAFE_INTEGER;
    let maxIndex = Number.MIN_SAFE_INTEGER;

    let firstNum = "";
    let lastNum = "";

    for (const [digit, minMax] of Object.entries(occurences)) {
        if (minMax["min"] < minIndex) {
            firstNum = digit;
            minIndex = minMax["min"];
        }

        if (minMax["max"] > maxIndex) {
            lastNum = digit;
            maxIndex = minMax["max"];
        }

    }

    return parseInt(convertToStringLiteral(firstNum) + convertToStringLiteral(lastNum));
};

const part2: SolverFunction = async (filePath: string) => {
    let result = 0;
    for await (const line of readLines(filePath)) {
        result += getNumValue(line);
    }

    return result;
};

const results: DaySolution = {
    part1,
    part2,
};

export { results };