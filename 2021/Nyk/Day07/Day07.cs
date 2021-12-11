using System;
using System.Collections.Generic;
using System.Diagnostics.CodeAnalysis;
using System.Linq;

using Utils;

namespace Day07
{
    public class Day07
    {
        static Dictionary<int, int> ParseInput(string line)
        {
            var crabInput = line.Split(",").Select(crab => crab.ToInt());
            var crabs = new Dictionary<int, int>();

            foreach (var crab in crabInput)
            {
                if (crabs.ContainsKey(crab))
                    crabs[crab]++;
                else crabs.Add(crab, 1);
            }

            return crabs;
        }

        static int CalculateShortestDistancePart1(Dictionary<int, int> crabs)
        {
            var distanceForCrabs = new Dictionary<int, int>();

            foreach (var crab in crabs)
            {
                distanceForCrabs.Add(crab.Key, 0);

                foreach (var crab2 in crabs)
                {
                    distanceForCrabs[crab.Key] += Math.Abs(crab.Key - crab2.Key) * crab2.Value;
                }
            }

            return distanceForCrabs.Select(crabDistance => crabDistance.Value).Min();
        }

        static int CalculateShortestDistancePart2(Dictionary<int, int> crabs)
        {
            var distanceForCrabs = new Dictionary<int, int>();
            var minPosition = crabs.Select(crabPosition => crabPosition.Key).Min();
            var maxPosition = crabs.Select(crabPosition => crabPosition.Key).Max();

            for (var pos = minPosition; pos < maxPosition; pos++)
            {
                distanceForCrabs.Add(pos, 0);

                foreach (var crab2 in crabs)
                {
                    var distance = Math.Abs(pos - crab2.Key);
                    distanceForCrabs[pos] += ((distance * (distance + 1)) / 2) * crab2.Value;
                }
            }

            //foreach (var dist in distanceForCrabs)
            //{
            //    Console.WriteLine($"{dist.Key}: {dist.Value}");
            //}

            return distanceForCrabs.Select(crabDistance => crabDistance.Value).Min();
        }

        public static long Part1(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var crabs = ParseInput(fileReader.ReadByLines(inputFile).ElementAt(0));

            return CalculateShortestDistancePart1(crabs);
        }

        public static long Part2(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var crabs = ParseInput(fileReader.ReadByLines(inputFile).ElementAt(0));

            return CalculateShortestDistancePart2(crabs);
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
