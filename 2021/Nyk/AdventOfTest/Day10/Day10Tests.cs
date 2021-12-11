using NUnit.Framework;


namespace AdventOfTest.Day10Tests
{
    internal class Day10Tests
    {
        [SetUp]
        public void Setup()
        {
        }

        [Test]
        public void TestPart1()
        {
            var result = Day10.Day10.Part1("Day10/TestInput.txt");
            Assert.AreEqual(26397, result);
        }

        [Test]
        public void TestPart2()
        {
            var result = Day10.Day10.Part2("Day10/TestInput.txt");
            Assert.AreEqual(288957, result);
        }
    }
}
