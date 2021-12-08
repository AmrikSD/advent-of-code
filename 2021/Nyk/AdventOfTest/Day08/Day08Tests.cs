using NUnit.Framework;


namespace AdventOfTest.Day08Tests
{
    internal class Day08Tests
    {
        [SetUp]
        public void Setup()
        {
        }

        [Test]
        public void TestPart1()
        {
            var result = Day08.Day08.Part1("Day08/TestInput.txt");
            Assert.AreEqual(26, result);
        }

        [Test]
        public void TestPart2()
        {
            var result = Day08.Day08.Part2("Day08/TestInput.txt");
            Assert.AreEqual(61229, result);
        }
    }
}
