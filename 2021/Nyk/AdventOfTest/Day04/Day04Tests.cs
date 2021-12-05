using NUnit.Framework;
using System;
using System.Collections.Generic;
using System.Linq;
using Utils;

namespace AdventOfTest.Day04Tests
{
    internal class Day04Tests
    {
        [SetUp]
        public void Setup()
        {
        }

        [Test]
        public void TestCalculateUncalledSum()
        {
            var filereader = new FileReader();
            var lines = filereader.ReadAllLines("Day04/TestInput.txt");
            var boardStrings = lines.Skip(2);

            var boards = new List<Board>();
            var rows = new List<List<string>>();

            foreach (var row in boardStrings)
            {
                if (row.Equals(""))
                {
                    boards.Add(new Board(rows));
                    rows.Clear();
                }
                else
                {
                    var line = row.Split(new char[0], StringSplitOptions.RemoveEmptyEntries).ToList();
                    rows.Add(line);
                }
            }

            var board = boards.ElementAt(2);

            board.UpdateValue(7);
            board.UpdateValue(4);
            board.UpdateValue(9);
            board.UpdateValue(5);
            board.UpdateValue(11);
            board.UpdateValue(17);
            board.UpdateValue(23);
            board.UpdateValue(2);
            board.UpdateValue(0);
            board.UpdateValue(14);
            board.UpdateValue(21);
            board.UpdateValue(24);

            Assert.AreEqual(188, board.CalculateUncalledSum());
        }

        [Test]
        public void TestPart1()
        {
            var result = Day04.Day04.Part1("Day04/TestInput.txt");
            Assert.AreEqual(4512, result);
        }

        [Test]
        public void TestPart2()
        {
            var result = Day04.Day04.Part2("Day04/TestInput.txt");
            Assert.AreEqual(1924, result);
        }
    }
}
