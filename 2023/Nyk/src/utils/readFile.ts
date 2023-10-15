//This sucks, stupid JS.
const readLines = async function* (filePath: string) {
    const fileLines = (await readFile(filePath)).split("/n");

    for (let line of fileLines) {
        yield line;
    }
}

const readFile = async (filePath: string) => await Bun.file(filePath).text();

export { readFile, readLines }