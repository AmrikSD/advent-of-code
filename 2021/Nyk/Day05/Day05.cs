using System;
using System.Diagnostics.CodeAnalysis;
using System.Linq;
using System.Collections.Generic;
using Utils;

namespace Day05
{
    public class Day05
    {
        static (Vector startVector, Vector endVector) ParseLine(string line)
        {
            var startEndList = line.Split(" -> ");
            return (_, _) = startEndList;
        }

        public static int Part1(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var grid = new Grid();

            foreach (var line in fileReader.ReadByLines(inputFile))
            {
                (var startVector, var endVector) = ParseLine(line);
                if (startVector.X == endVector.X || startVector.Y == endVector.Y)
                    grid.TraverseGridHorizontalAndVertical(startVector, endVector);
            }

            return grid.GetNodeGreaterThanTwo();
        }

        public static int Part2(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var grid = new Grid();

            foreach (var line in fileReader.ReadByLines(inputFile))
            {
                (var startVector, var endVector) = ParseLine(line);
                if (startVector.X == endVector.X || startVector.Y == endVector.Y)
                    grid.TraverseGridHorizontalAndVertical(startVector, endVector);
                else
                    grid.TraverseGridDiagonal(startVector, endVector);
            }

            return grid.GetNodeGreaterThanTwo();
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
