using NUnit.Framework;


namespace AdventOfTest.Day06Tests
{
    internal class Day06Tests
    {
        [SetUp]
        public void Setup()
        {
        }

        [Test]
        public void TestPart1()
        {
            var result = Day06.Day06.Part1("Day06/TestInput.txt");
            Assert.AreEqual(5934, result);
        }

        [Test]
        public void TestPart2()
        {
            var result = Day06.Day06.Part2("Day06/TestInput.txt");
            Assert.AreEqual(26984457539, result);
        }
    }
}
