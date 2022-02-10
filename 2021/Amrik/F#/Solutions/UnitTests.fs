namespace Solutions

open NUnit.Framework
open Day_01

module UnitTests = 
    [<SetUp>]
    let Setup () =
        ()

    [<Test>]
    let Test1 () =
        Assert.Pass()
    
    [<Test>]
    let ``Day 1 Part 1`` () = 
        Assert.AreEqual(7, PartOne "./Input/Samples/Day01.txt")
    
    [<Test>]
    let ``Day 1 Part 2`` () =
        Assert.AreEqual(5, PartTwo "./Input/Samples/Day01.txt")