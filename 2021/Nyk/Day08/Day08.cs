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

        static List<string> DiscoverNumberConfigs(IEnumerable<string> inputs)
        {
            /*
              0000
             6    1
             6    1
              2222
             5    3
             5    3
              4444
            */

            var letterPositions = new List<string>() { "", "", "", "", "", "", "" };

            var oneConfig = inputs.Where(input => input.Length == 2).ElementAt(0);
            var fourConfig = inputs.Where(input => input.Length == 4).ElementAt(0);
            var sevenConfig = inputs.Where(input => input.Length == 3).ElementAt(0);
            var eightConfig = inputs.Where(input => input.Length == 7).ElementAt(0);

            var sixConfigs = inputs.Where(input => input.Length == 6);
            var fiveConfigs = inputs.Where(input => input.Length == 5);

            letterPositions[0] = sevenConfig.Where(letter => !oneConfig.Contains(letter)).ElementAt(0).ToString(); //Find the odd character out from 7 and 1

            var threeCharConfig = fiveConfigs.Where(config => sevenConfig.Count(letter => config.Contains(letter)) != 3).ElementAt(0); //find the configuration that represents 3
            letterPositions[6] = fourConfig.Where(letter => !threeCharConfig.Contains(letter)).ElementAt(0).ToString(); //Find the odd character out from 3 and 4

            var fiveCharConfig = fiveConfigs.Where(config => config.Contains(letterPositions[6])).ElementAt(0).ToString();
            var twoCharConfig = fiveConfigs.Where(config => !config.Contains(letterPositions[6])).ElementAt(0).ToString();
            var missingFromFive = twoCharConfig.Where(letter => !fiveCharConfig.Contains(letter));
            letterPositions[1] = missingFromFive.Where(letter => oneConfig.Contains(letter)).ElementAt(0).ToString();
            letterPositions[5] = missingFromFive.Where(letter => !oneConfig.Contains(letter)).ElementAt(0).ToString();
            letterPositions[3] = oneConfig.Where(letter => letter.ToString() != letterPositions[2]).ElementAt(0).ToString();

            var zeroCharConfig = sixConfigs
                .Where(config => !config
                .Contains(fiveCharConfig
                .Where(letter => !config
                .Contains(letter))
                .ElementAt(0)
                .ToString()))
                .ElementAt(0)
                .ToString();

            return letterPositions;
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
            var totalSum = 0;

            foreach (var line in fileReader.ReadByLines(inputFile))
            {

                var lineValues = ParseInput(line);
                var inputs = lineValues.ElementAt(0);
                var outputs = lineValues.ElementAt(1);

                var numberConfigs = DiscoverNumberConfigs(inputs);

                var sum = 0;

                foreach (var output in outputs)
                {
                    foreach (var numberConfig in numberConfigs)
                    {
                        if (CheckStringIsNumber(output, numberConfig))
                        {
                            sum *= 10;
                            sum += numberConfigs[1].IndexOf(numberConfig);
                            break;
                        }
                    }

                }

                totalSum += sum;
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
