namespace Solutions

open System.IO

module Day_01 =

// https://docs.microsoft.com/en-us/dotnet/fsharp/language-reference/sequences

    let PartOne (path:string) =
        File.ReadAllLines(path)
            |> Seq.map int
            |> Seq.pairwise
            |> Seq.filter (fun (a,b) -> b>a)
            |> Seq.length
    
    let PartTwo (path) =
        File.ReadAllLines path
            |> Seq.map int
            |> Seq.windowed 3
            |> Seq.pairwise
            |> Seq.filter (fun (a, b) -> Seq.sum a < Seq.sum b)
            |> Seq.length