import { exit } from "node:process";

if (process.argv.length !== 3) {
    console.error("You need to enter a day to generate. Usage: bun run solution Day01")
    exit(1);
}

const day = `Day${process.argv[2].padStart(2, "0")}`

const runSolutionPath = `../${day}/index`

const run = async () => {
    let { runSolutions } = await import(runSolutionPath);
    await runSolutions(day)
}

await run();