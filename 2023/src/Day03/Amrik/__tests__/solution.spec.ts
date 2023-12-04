import { test, expect, describe } from "bun:test";
import { results } from "../solution";

describe("Day03", () => {
    test("part1", async () => {
        expect(await results.part1("src/Day03/Amrik/__tests__/input.txt")).toBe(2);
    })

    test("part2", async () => {
        expect(await results.part2("src/Day03/Amrik/__tests__/input2.txt")).toBe(2);
    })
})
