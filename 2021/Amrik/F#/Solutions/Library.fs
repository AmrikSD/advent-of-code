namespace Solutions

open System.IO

module Lib = 
    let readLines (path:string) =
        File.ReadAllLines(path)

module Day_01 =
    let PartOne (path:string) =
        let input = File.ReadAllLines(path)
        input[0]
        