using System;
using System.Collections.Generic;
using System.Diagnostics.CodeAnalysis;
using System.Linq;
using Utils;

namespace Day08
{
    public class Day08
    {
        static IEnumerable<IEnumerable<string>> ParseInput(string line)
        {
            return line.Split(" | ").Select(element => element.Trim().Split(' '));
        }

        public static long Part1(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var sumOf1748 = 0;

            foreach (var line in fileReader.ReadByLines(inputFile))
            {
                var lineValues = ParseInput(line);
                var outputs = lineValues.ElementAt(1);

                foreach (var output in outputs)
                {
                    sumOf1748 +=
                        output.Length == 2 ||
                        output.Length == 4 ||
                        output.Length == 3 ||
                        output.Length == 7 ? 1 : 0;
                }
            }

            return sumOf1748;
        }

        static List<List<string>> GenerateNumberConfigs()
        {
            return new List<List<string>>(){
                new List<string>{
                "bacefg",
                 "cf",
                 "acdeg",
                "acdfg",
                 "bdcf",
                 "abdfg",
                 "abdefg",
                 "acf",
                 "acedgfb",
                 "abcdfg",
                },
                new List<string>{
                "cagedb",
                 "ab",
                 "gcdfa",
                "fbcad",
                 "eafb",
                 "cdfbe",
                 "cdfgeb",
                 "dab",
                 "acedgfb",
                 "cefabd",
                }
            };
        }

        static bool CheckStringIsNumber(string output, string numberConfig)
        {
            foreach (var letter in output)
            {
                if (!numberConfig.Contains(letter))
                    return false;
            }

            return true;
        }

        public static long Part2(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var numberConfigs = GenerateNumberConfigs();
            var totalSum = 0;

            foreach (var line in fileReader.ReadByLines(inputFile))
            {
                var lineValues = ParseInput(line);
                var inputs = lineValues.ElementAt(0);
                var outputs = lineValues.ElementAt(1);
                var badLine = false;

                foreach (var input in inputs)
                {
                    foreach (var numberConfig in numberConfigs[1])
                    {
                        if (!CheckStringIsNumber(input, numberConfig))
                        {
                            badLine = true;
                            break;
                        }
                    }

                }

                var stringRepresentation = "";

                if (!badLine)
                {
                    foreach (var output in outputs)
                    {
                        foreach (var numberConfig in numberConfigs[1])
                        {
                            if (CheckStringIsNumber(output, numberConfig))
                            {
                                stringRepresentation += numberConfigs[1].IndexOf(numberConfig);
                                break;
                            }
                        }

                    }

                    totalSum += int.Parse(stringRepresentation);
                }

            }

            return totalSum;
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
