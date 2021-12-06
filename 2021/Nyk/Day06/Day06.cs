using System;
using System.Collections.Generic;
using System.Diagnostics.CodeAnalysis;
using System.Linq;
using Utils;

namespace Day06
{
    public class Day06
    {
        static IEnumerable<long> ParseInput(string line)
        {
            return line.Split(",").Select(timer => long.Parse(timer));
        }

        public static int Part1(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var lanternFish = new List<long>();
            lanternFish.AddRange(ParseInput(fileReader.ReadAllLines(inputFile).ElementAt(0)));

            for (var day = 0; day < 80; day++)
            {
                Console.WriteLine($"Starting day: {day} at {DateTime.Now.TimeOfDay}");
                var newFish = new List<long>();

                for (var fishIndex = 0; fishIndex < lanternFish.Count; fishIndex++)
                {
                    if (lanternFish.ElementAt(fishIndex) == 0)
                    {
                        newFish.Add(8);
                        lanternFish[fishIndex] = 6;
                    }
                    else lanternFish[fishIndex] -= 1;
                }
                lanternFish.AddRange(newFish);
            }

            return lanternFish.Count();
        }

        public static int Part2(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var lanternFish = new List<long>();
            lanternFish.AddRange(ParseInput(fileReader.ReadAllLines(inputFile).ElementAt(0)));

            for (int day = 0; day < 256; day++)
            {
                Console.WriteLine($"Starting day: {day} at {DateTime.Now.TimeOfDay}");
                Console.WriteLine($"Amount of Lantern Fish {lanternFish.Count}");
                var newFish = new List<long>();

                for (var fishIndex = 0; fishIndex < lanternFish.Count; fishIndex++)
                {
                    if (lanternFish.ElementAt(fishIndex) == 0)
                    {
                        newFish.Add(8);
                        lanternFish[fishIndex] = 6;
                    }
                    else lanternFish[fishIndex] -= 1;
                }
                lanternFish.AddRange(newFish);
            }

            return lanternFish.Count();
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
