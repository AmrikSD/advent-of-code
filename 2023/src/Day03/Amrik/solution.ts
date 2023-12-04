import { SolverFunction, DaySolution } from "../../types"
import { readFile } from "utils/readFile";


const part1: SolverFunction = async (filePath: string) => {

    let data: string[][]  = (await readFile(filePath)).split("\n").slice(0,-1).map((str: string) => str.split(""))

    const isSymbol = (str: string): boolean => !(str == "." || Number.isInteger(Number(str)))
    const isNumber = (str: string): boolean => Number.isInteger(Number(str))

    let sum = 0

    for (let y = 0; y < data.length; y++){
        for (let x = 0; x < data.length; x++){
            let char = data[y][x]
            if (!isSymbol(char)){
                continue
            }

            for(let i = -1; i <= 1; i++){
                for(let j = -1; j <= 1; j++){
                    let char = data[y+i][x+j]
                    if (!isNumber(char)){
                        continue
                    }
                    let pointer = 0
                    while(isNumber(data[y+i][x+j+pointer+1])){
                        pointer++
                    }
                    let tens = 1
                    let num = 0
                    while(isNumber(data[y+i][x+j+pointer])){
                        num += tens * (Number(data[y+i][x+j+pointer]))
                        data[y+i][x+j+pointer] = "."
                        tens *= 10
                        pointer--
                    }
                    sum += num
                }
            }
        }
    }


    return sum
};


const part2: SolverFunction = async (filePath: string) => {
    let data: string[][]  = (await readFile(filePath)).split("\n").slice(0,-1).map((str: string) => str.split(""))

    const isSymbol = (str: string): boolean => (str == "*")
    const isNumber = (str: string): boolean => Number.isInteger(Number(str))

    let sum = 0

    for (let y = 0; y < data.length; y++){
        for (let x = 0; x < data.length; x++){
            let char = data[y][x]
            if (!isSymbol(char)){
                continue
            }

            let prod = []
            for(let i = -1; i <= 1; i++){
                for(let j = -1; j <= 1; j++){
                    let char = data[y+i][x+j]
                    if (!isNumber(char)){
                        continue
                    }
                    let pointer = 0
                    while(isNumber(data[y+i][x+j+pointer+1])){
                        pointer++
                    }
                    let tens = 1
                    let num = 0
                    while(isNumber(data[y+i][x+j+pointer])){
                        num += tens * (Number(data[y+i][x+j+pointer]))
                        data[y+i][x+j+pointer] = "."
                        tens *= 10
                        pointer--
                    }
                    prod.push(num)
                }
            }
            if(prod.length == 2){
                sum += (prod[0] * prod[1])
                console.log(prod)
            }
        }
    }


    return sum
};

const results: DaySolution = {
    part1,
    part2,
}

export { results }
