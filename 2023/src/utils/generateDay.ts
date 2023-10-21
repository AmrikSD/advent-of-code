import { exit } from "node:process"
import fs from "fs"
import mustache from "mustache"

const participants = [
    { name: "Amrik"},
    { name: "Binda"},
    { name: "Nyk" },
]

if (process.argv.length !== 3) {
    console.error("You need to enter a day to generate.\n Usage: bun run generateDay Day01")
    exit(1);
}

const day = `Day${process.argv[2].padStart(2, "0")}`
console.log(`Generating ${day}`)
fs.mkdirSync(`src/${day}`, { recursive: true })


const templatePath = "dayTemplate/";

const idxTemplate = fs.readFileSync(`${templatePath}/index.ts.tpl`, "utf-8")
const index = mustache.render(idxTemplate, {participants})

fs.writeFileSync(`src/${day}/index.ts`, index)

participants.forEach(participant => {
    console.log(`Generating ${day} for ${participant.name}`)
    fs.mkdirSync(`src/${day}/${participant.name}`, { recursive: true })
    fs.copyFileSync(`${templatePath}/participant/input.txt`, `src/${day}/${participant.name}/input.txt`)
    fs.copyFileSync(`${templatePath}/participant/solution.ts`, `src/${day}/${participant.name}/solution.ts`)
    fs.mkdirSync(`src/${day}/${participant.name}/__tests__`, { recursive: true })

    const testTemplate = fs.readFileSync(`${templatePath}participant/__tests__/solution.spec.ts.tpl`, "utf-8")
    const path = `src/${day}/${participant.name}/input.txt`
    const test = mustache.render(testTemplate, {path: path})
    console.log(test)
    fs.writeFileSync(`src/${day}/${participant.name}/__tests__/solution.spec.ts`, test)
})
