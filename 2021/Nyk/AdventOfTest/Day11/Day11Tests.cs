using NUnit.Framework;


namespace AdventOfTest.Day11Tests
{
    internal class Day11Tests
    {
        [SetUp]
        public void Setup()
        {
        }

        [Test]
        public void TestPart1()
        {
            var result = Day11.Day11.Part1("Day11/TestInput.txt");
            Assert.AreEqual(1656, result);
        }

        [Test]
        public void TestPart2()
        {
            var result = Day11.Day11.Part2("Day11/TestInput.txt");
            Assert.AreEqual(195, result);
        }
    }
}
