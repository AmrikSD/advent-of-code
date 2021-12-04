using NUnit.Framework;
using Day01;

namespace AdventOfTest.Day01Tests
{
    internal class Day01Tests
    {
        [SetUp]
        public void Setup()
        {
        }

        [Test]
        public void TestPart1()
        {
            var result = Day01.Day01.Part1("Day01/TestInput.txt");
            Assert.AreEqual(7, result);
        }

        [Test]
        public void TestPart2()
        {
            var result = Day01.Day01.Part2("Day01/TestInput.txt");
            Assert.AreEqual(5, result);
        }
    }
}
