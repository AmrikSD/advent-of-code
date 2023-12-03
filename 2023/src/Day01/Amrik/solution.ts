import { _resolveFilename } from "node:module";
import { SolverFunction, DaySolution } from "../../types"
import { readFile } from "utils/readFile";

const tokens = [
    "1", "one",
    "2", "two",
    "3", "three",
    "4", "four",
    "5", "five",
    "6", "six",
    "7", "seven",
    "8", "eight",
    "9", "nine",
]

const token_to_number = new Map<string,number>([
    ["one", 1],
    ["two", 2],
    ["three", 3],
    ["four", 4],
    ["five", 5],
    ["six", 6],
    ["seven", 7],
    ["eight", 8],
    ["nine", 9],
])
const sanitize = (numberOrString: string): number => {
    if(token_to_number.has(numberOrString)) {
        return token_to_number.get(numberOrString) ?? 0
    }
    return Number.parseInt(numberOrString)
}

const part1: SolverFunction = async (filePath: string) => {
    let data: string[]  = (await readFile(filePath)).split("\n").slice(0,-1)

    return data.map((str) => {
        let hits = []
        for(let i = 0; i<str.length; i++){
            if (Number.isInteger(Number(str[i]))){
                hits.push(Number.parseInt(str[i]))
            }
        }
        return {high: hits[0], low: hits[hits.length-1]}
    }).reduce((acc,curr) => acc+curr.high*10+curr.low, 0)

};

const part2: SolverFunction = async (filePath: string) => {
    let data: string[] = (await readFile(filePath)).split("\n").slice(0,-1)


    return data.map((str) => {
        let min_hm = new Map<number, number>();
        let max_hm = new Map<number, number>();
        tokens.forEach((token) => {
            min_hm.set(str.indexOf(token), sanitize(token))
            max_hm.set(str.lastIndexOf(token), sanitize(token))
        })

        min_hm.delete(-1)
        max_hm.delete(-1)
        return {min_hm,max_hm}

    }).map(({min_hm, max_hm}) => {


        let max_key = Number.MIN_SAFE_INTEGER
        let min_key = Number.MAX_SAFE_INTEGER

        min_hm.forEach((_, key) => {
            if (key < min_key){
                min_key = key
            }
        })
        max_hm.forEach((_, key) => {
            if (key > max_key){
                max_key = key
            }
        })

        return {high: min_hm.get(min_key) ?? 0, low: max_hm.get(max_key) ?? 0}

    }).map(({high,low}) => {
        return high * 10 + low
    }).reduce((acc,curr) => acc + curr, 0)
};

const results: DaySolution = {
    part1,
    part2,
}

export { results }
