import { test, expect, describe } from "bun:test";
import { results } from "../solution";

describe("Day01", () => {
    test("part1", async () => {
        expect(await results.part1("src/Day01/Nyk/__tests__/input.txt")).toBe(142);
    });

    test("part2", async () => {
        expect(await results.part2("src/Day01/Nyk/__tests__/input2.txt")).toBe(281);
    });
});
