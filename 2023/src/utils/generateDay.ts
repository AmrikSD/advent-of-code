import { exit } from 'node:process'
import fs from 'fs'
import mustache from 'mustache'

if (process.argv.length !== 4) {
    console.error('You need to enter a day to generate, and a participant.\n Usage: bun run generateDay Day01 Greg')
    exit(1);
}

//Skip the first 2 values, because we don't care about them.
let [, , day] = process.argv
const [, , , participant] = process.argv

day = `Day${day.padStart(2, "0")}`

console.log(`Generating ${day} for ${participant}`)
fs.mkdirSync(`src/${day}`, { recursive: true })

const templatePath = "dayTemplate/";

const idxTemplate = fs.readFileSync(`${templatePath}/index.ts.tpl`, "utf-8")
const index = mustache.render(idxTemplate, { participant })

console.log(`Generating file: src/${day}/index.ts`)
fs.writeFileSync(`src/${day}/index.ts`, index)

fs.mkdirSync(`src/${day}/${participant}`, { recursive: true })
console.log(`Generating file: src/${day}/${participant}/input.txt`)
fs.copyFileSync(`${templatePath}/participant/input.txt`, `src/${day}/${participant}/input.txt`)

console.log(`Generating file: src/${day}/${participant}/solution.ts`)
fs.copyFileSync(`${templatePath}/participant/solution.ts`, `src/${day}/${participant}/solution.ts`)


fs.mkdirSync(`src/${day}/${participant}/__tests__`, { recursive: true })
console.log(`Generating file: src/${day}/${participant}/__tests__/input.txt`)
fs.copyFileSync(`${templatePath}/participant/__tests__/input.txt`, `src/${day}/${participant}/__tests__/input.txt`)
console.log(`Generating file: src/${day}/${participant}/__tests__/input2.txt`)
fs.copyFileSync(`${templatePath}/participant/__tests__/input.txt`, `src/${day}/${participant}/__tests__/input2.txt`)

const testTemplate = fs.readFileSync(`${templatePath}participant/__tests__/solution.spec.ts.tpl`, "utf-8")
const path = `src/${day}/${participant}/__tests__/input.txt`
const path2 = `src/${day}/${participant}/__tests__/input2.txt`

const test = mustache.render(testTemplate, { path, path2, day })

console.log(`Generating file: src/${day}/${participant}/__tests__/solution.spec.ts`)
fs.writeFileSync(`src/${day}/${participant}/__tests__/solution.spec.ts`, test)

