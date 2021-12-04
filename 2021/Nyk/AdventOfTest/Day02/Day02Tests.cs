using NUnit.Framework;

namespace AdventOfTest.Day02Tests
{
    internal class Day02Tests
    {
        [SetUp]
        public void Setup()
        {
        }

        [Test]
        public void TestPart1()
        {
            var result = Day02.Day02.Part1("Day02/TestInput.txt");
            Assert.AreEqual(150, result);
        }

        [Test]
        public void TestPart2()
        {
            var result = Day02.Day02.Part2("Day02/TestInput.txt");
            Assert.AreEqual(900, result);
        }
    }
}
