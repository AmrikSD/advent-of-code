using System;
using System.Collections.Generic;
using System.Diagnostics.CodeAnalysis;
using System.Linq;
using Utils;

namespace Day04
{
    public class Day04
    {
        static (IEnumerable<Board>, IEnumerable<int>) ParseInput(IEnumerable<string> input)
        {
            var calledNumbers = input.ElementAt(0).Split(',').ToList().Select(value => int.Parse(value));
            var boardStrings = input.Skip(2);

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

            return (boards, calledNumbers);
        }

        public static int Part1(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var input = fileReader.ReadAllLines(inputFile);

            (var boards, var numInput) = ParseInput(input);

            foreach (var num in numInput)
            {
                foreach (var board in boards)
                {
                    board.UpdateValue(num);
                    if (board.CheckBingo()) return board.CalculateUncalledSum() * num;
                }
            }

            return 0;
        }
        public static int Part2(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var input = fileReader.ReadAllLines(inputFile);

            (var boards, var numInput) = ParseInput(input);

            var winningBoards = new List<Board>();

            foreach (var num in numInput)
            {
                foreach (var board in boards)
                {
                    board.UpdateValue(num);
                    if (board.CheckBingo())
                    {
                        if (winningBoards.Count() == boards.Count() - 1 && !winningBoards.Contains(board))
                            return board.CalculateUncalledSum() * num;
                        else if (!winningBoards.Contains(board))
                            winningBoards.Add(board);
                    }
                }
            }
            return 0;
        }

        [ExcludeFromCodeCoverage]
        static void Main(string[] args)
        {
            var result1 = Part1();
            var result2 = Part2();

            Console.WriteLine($"Part 1: {result1}");
            Console.WriteLine($"Part 2: {result2}");
        }
    }
}
