using NUnit.Framework;


namespace AdventOfTest.Day09Tests
{
    internal class Day09Tests
    {
        [SetUp]
        public void Setup()
        {
        }

        [Test]
        public void TestPart1()
        {
            var result = Day09.Day09.Part1("Day09/TestInput.txt");
            Assert.AreEqual(15, result);
        }

        [Test]
        public void TestPart2()
        {
            var result = Day09.Day09.Part2("Day09/TestInput.txt");
            Assert.AreEqual(1134, result);
        }
    }
}
