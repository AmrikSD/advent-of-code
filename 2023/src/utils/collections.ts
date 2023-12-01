const toWindows = <T>(inputArray: T[], size: number) => {
    return Array.from(
        {length: inputArray.length - (size-1)},
        (_,index) => inputArray.slice(index, index+size)
    )
}

export { toWindows }
