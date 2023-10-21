import { test, expect, describe, beforeAll } from "bun:test";
import { results } from "../solution";

describe("tests", () => {
    test("part1", async () => {
        expect(await results.part1("{{{path}}}")).toBe(2);
    })

    test("part2", async () => {
        expect(await results.part1("{{{path}}}")).toBe(2);
    })
})
