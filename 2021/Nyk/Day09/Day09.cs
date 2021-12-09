using System;
using System.Collections.Generic;
using System.Diagnostics.CodeAnalysis;
using System.Linq;
using Utils;

namespace Day09
{
    public class Day09
    {
        static IEnumerable<IEnumerable<int>> ParseInput(IEnumerable<string> lines)
        {
            return lines.Select(line => line.Select(numChar => int.Parse(numChar.ToString())));
        }

        static int GetAdjacent(IEnumerable<IEnumerable<int>> heightMap, int rowIndex, int colIndex)
        {
            if (rowIndex < 0 || rowIndex >= heightMap.Count() ||
                colIndex < 0 || colIndex >= heightMap.ElementAt(0).Count())
                return -1;
            else return heightMap.ElementAt(rowIndex).ElementAt(colIndex);
        }

        public static int Part1(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var heightMap = ParseInput(fileReader.ReadAllLines(inputFile));
            var lowestPoints = new List<int>();

            for (var row = 0; row < heightMap.Count(); row++)
            {
                for (var col = 0; col < heightMap.ElementAt(row).Count(); col++)
                {
                    var lowestPointRow = new List<int>();

                    lowestPointRow.Add(GetAdjacent(heightMap, row - 1, col)); //up
                    lowestPointRow.Add(GetAdjacent(heightMap, row + 1, col)); //down
                    lowestPointRow.Add(GetAdjacent(heightMap, row, col - 1)); //left
                    lowestPointRow.Add(GetAdjacent(heightMap, row, col + 1)); // right

                    lowestPointRow.RemoveAll(ints => ints == -1);
                    var min = lowestPointRow.Min();

                    if (min > heightMap.ElementAt(row).ElementAt(col))
                        lowestPoints.Add(heightMap.ElementAt(row).ElementAt(col) + 1);
                }
            }

            return lowestPoints.Sum();
        }

        public static long Part2(string inputFile = "input.txt")
        {
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
