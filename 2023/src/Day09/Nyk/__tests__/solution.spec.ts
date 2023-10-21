import { test, expect, describe, beforeAll } from "bun:test";
import { results } from "../solution";

let inputPath = ""

describe("tests", () => {
    beforeAll(() => {
        const currentPathArray = process.argv[1].split("/").slice(0, -2);
        console.log(currentPathArray)
        inputPath = `${currentPathArray.join("/")}/input.txt`
    })

    test("part1", async () => {
        expect(await results.part1(inputPath)).toBe(2);
    })

    test("part2", async () => {
        expect(await results.part1(inputPath)).toBe(2);
    })
})