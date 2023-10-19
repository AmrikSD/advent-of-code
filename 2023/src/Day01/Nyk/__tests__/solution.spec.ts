import { test, expect, describe } from "bun:test";
import { readFile } from "../../../utils/readFile"
import { daySolution } from "../solution";
import { resolve } from "path";

describe("Day01", () => {
    test("part1", async () => {
        const filePath = resolve("src/Day01/Nyk/input.txt")
        expect(await daySolution.part1(readFile, filePath)).toBe(1);
    })

    test("part2", async () => {
        const filePath = resolve("src/Day01/Nyk/input.txt")
        expect(await daySolution.part1(readFile, filePath)).toBe(1);
    })
})