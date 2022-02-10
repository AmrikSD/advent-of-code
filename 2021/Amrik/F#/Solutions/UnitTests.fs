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
        let actual = PartOne "./Input/Samples/Day01.txt"
        let expected = 7
        Assert.AreEqual(expected,actual)