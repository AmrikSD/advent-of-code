using NUnit.Framework;


namespace AdventOfTest.Day05Tests
{
    internal class Day05Tests
    {
        [SetUp]
        public void Setup()
        {
        }

        [Test]
        public void TestPart1()
        {
            var result = Day05.Day05.Part1("Day05/TestInput.txt");
            Assert.AreEqual(5, result);
        }

        [Test]
        public void TestPart2()
        {
            var result = Day05.Day05.Part2("Day05/TestInput.txt");
            Assert.AreEqual(12, result);
        }
    }
}
