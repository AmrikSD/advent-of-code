import { cpSync } from "node:fs"
import { exit } from "node:process";

if (process.argv.length !== 3) {
    console.error("You need to enter a day to generate. Usage: bun run generateDay Day01")
    exit(1);
}

const day = `Day${process.argv[2].padStart(2, "0")}`

const templatePath = "dayTemplate/";
const newSolutionPath = `src/${day}/`

cpSync(templatePath, newSolutionPath, { recursive: true });