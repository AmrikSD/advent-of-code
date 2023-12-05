import { test, expect, describe } from "bun:test";
import { results } from "../solution";

describe("Day04", () => {
    test("part1", async () => {
        expect(await results.part1("src/Day04/Nyk/__tests__/input.txt")).toBe(13);
    });

    test("part2", async () => {
        expect(await results.part2("src/Day04/Nyk/__tests__/input.txt")).toBe(30);
    });
});
