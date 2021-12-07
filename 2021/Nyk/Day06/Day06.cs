using System;
using System.Collections.Generic;
using System.Diagnostics.CodeAnalysis;
using System.Linq;
using Utils;

namespace Day06
{
    public class Day06
    {
        static IEnumerable<int> ParseInput(string line)
        {
            return line.Split(",").Select(timer => int.Parse(timer));
        }

        public static long Part1(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var lanternFish = new LanternFishSchool(ParseInput(fileReader.ReadAllLines(inputFile).ElementAt(0)));

            for (var day = 0; day < 80; day++)
            {
                lanternFish.IterateDay();
            }

            return lanternFish.CountFish();
        }

        public static long Part2(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var lanternFish = new LanternFishSchool(ParseInput(fileReader.ReadAllLines(inputFile).ElementAt(0)));

            for (var day = 0; day < 256; day++)
            {
                lanternFish.IterateDay();
            }

            return lanternFish.CountFish();
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
