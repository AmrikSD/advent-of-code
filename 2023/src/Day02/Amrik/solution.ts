import { SolverFunction, DaySolution } from "../../types"
import { readFile } from "utils/readFile";

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

const part1: SolverFunction = async (filePath: string) => {
    let data: string[] = (await readFile(filePath)).split("\n")

    return data
    .map((str) => {
        const id = str.split(":")[0].split(" ")[1]
        const data = str.split(":")[1]?.split(";")
        return {
            id,
            data
        }
    })
    .filter((game) => game.data != undefined)
    .filter((game) => {
        var keep = true

        game.data.forEach((element) => {
            element.split(",").forEach((hand) => {
                let num = Number(hand.split(" ")[1].split(" ")[0])

                if (hand.includes("red") && num > MAX_RED){
                    keep = false
                }
                if (hand.includes("green") && num > MAX_GREEN){
                    keep = false
                }
                if (hand.includes("blue") && num > MAX_BLUE){
                    keep = false
                }

            })

        })
        return keep
    })
    .map((curr) => Number(curr.id))
    .reduce((acc,curr) => acc + curr, 0)
};

const part2: SolverFunction = async (filePath: string) => {
    let data: string[] = (await readFile(filePath)).split("\n")

    const reduced = data
    .map((str) => {
        const id = str.split(":")[0].split(" ")[1]
        const data = str.split(":")[1]?.split(";")
        return {
            id,
            data
        }
    })
    .filter((game) => game.data != undefined)
    .map((game) => {
        let MIN_RED = 0
        let MIN_GREEN = 0
        let MIN_BLUE = 0

        game.data.forEach((element) => {
            element.split(",").forEach((hand) => {
                let num = Number(hand.split(" ")[1].split(" ")[0])

                if (hand.includes("red") && num > MIN_RED){
                    MIN_RED = num
                }
                if (hand.includes("green") && num > MIN_GREEN){
                    MIN_GREEN = num
                }
                if (hand.includes("blue") && num > MIN_BLUE){
                    MIN_BLUE = num
                }

            })

        })
        return {MIN_RED, MIN_GREEN, MIN_BLUE}
    })
    .reduce((acc,val) => acc + (val.MIN_RED * val.MIN_GREEN * val.MIN_BLUE), 0)
    console.log(reduced)
    return 3
};

const results: DaySolution = {
    part1,
    part2,
}

export { results }
