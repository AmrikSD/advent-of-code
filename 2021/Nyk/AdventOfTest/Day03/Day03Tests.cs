using NUnit.Framework;

namespace AdventOfTest.Day03Tests
{
    internal class Day03Tests
    {
        [SetUp]
        public void Setup()
        {
        }

        [Test]
        public void TestPart1()
        {
            var result = Day03.Day03.Part1("Day03/TestInput.txt");
            Assert.AreEqual(198, result);
        }

        [Test]
        public void TestPart2()
        {
            var result = Day03.Day03.Part2("Day03/TestInput.txt");
            Assert.AreEqual(230, result);
        }
    }
}
