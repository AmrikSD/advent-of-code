using NUnit.Framework;


namespace AdventOfTest.Day07Tests
{
    internal class Day07Tests
    {
        [SetUp]
        public void Setup()
        {
        }

        [Test]
        public void TestPart1()
        {
            var result = Day07.Day07.Part1("Day07/TestInput.txt");
            Assert.AreEqual(37, result);
        }

        [Test]
        public void TestPart2()
        {
            var result = Day07.Day07.Part2("Day07/TestInput.txt");
            Assert.AreEqual(168, result);
        }
    }
}
