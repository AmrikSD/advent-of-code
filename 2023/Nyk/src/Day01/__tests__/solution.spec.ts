import { test, expect } from "bun:test";
import { readFile } from "../../utils/readFile";
import { solver } from "../solution";

test("solution result", async () => {
    expect(await solver(readFile, "input.txt")).toBe(1);
})