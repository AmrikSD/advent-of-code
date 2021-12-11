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

        static List<char> CalculateLetterPositions(IEnumerable<string> inputs)
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

            var letterPositions = new List<char>() { ' ', ' ', ' ', ' ', ' ', ' ', ' ' };

            var oneConfig = inputs.Where(input => input.Length == 2).ElementAt(0); //1
            var fourConfig = inputs.Where(input => input.Length == 4).ElementAt(0); //4
            var sevenConfig = inputs.Where(input => input.Length == 3).ElementAt(0); //7
            var eightConfig = inputs.Where(input => input.Length == 7).ElementAt(0); //8

            var sixConfigs = inputs.Where(input => input.Length == 6);
            var fiveConfigs = inputs.Where(input => input.Length == 5);

            letterPositions[0] = sevenConfig.Where(letter => !oneConfig.Contains(letter)).ElementAt(0); //Find the odd character out from 7 and 1
            var threeConfig = fiveConfigs.Where(config => config.Where(letter => sevenConfig.Contains(letter)).Count() == 3).ElementAt(0); //Find the string representing 3
            fiveConfigs = fiveConfigs.Where(config => config != threeConfig); //Remove 3 string so we can get 5 and 2 from remaining

            letterPositions[6] = fourConfig.Where(letter => !threeConfig.Contains(letter)).ElementAt(0); //Find the odd character out from 4 and 3

            var twoConfig = fiveConfigs.Where(config => !config.Contains(letterPositions[6])).ElementAt(0); //Find the string representing 2
            var fiveConfig = fiveConfigs.Where(config => config.Contains(letterPositions[6])).ElementAt(0); //Find the string representing 5

            letterPositions[1] = sevenConfig.Where(letter => !fiveConfig.Contains(letter)).ElementAt(0); //Find the odd character between 7 and 5
            letterPositions[3] = oneConfig.Where(letter => letter != letterPositions[1]).ElementAt(0); //Find the other letter position of 1

            var foundLettersSoFar = new List<char>() { letterPositions[1], letterPositions[6], letterPositions[3] }; //Known character for 4 config

            letterPositions[2] = fourConfig.Where(letter => !foundLettersSoFar.Contains(letter)).ElementAt(0); //Find letter missing from 4 config

            foundLettersSoFar = new List<char>() { letterPositions[0], letterPositions[1], letterPositions[2], letterPositions[3] }; //Known characters for 3 config

            letterPositions[4] = threeConfig.Where(letter => !foundLettersSoFar.Contains(letter)).ElementAt(0); //Find letter missing from 3 config

            foundLettersSoFar = new List<char>(){
                letterPositions[0], letterPositions[1], letterPositions[2],
                letterPositions[3], letterPositions[4], letterPositions[6]
            }; //All known charcters

            letterPositions[5] = eightConfig.Where(letter => !foundLettersSoFar.Contains(letter)).ElementAt(0);//Find last letter missing 

            return letterPositions;
        }

        static List<string> TranslatePositionsToStringConfigs(List<char> letterPositions)
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
            var zero = new List<char>() { letterPositions[0], letterPositions[1], letterPositions[3], letterPositions[4], letterPositions[5], letterPositions[6] };
            var one = new List<char>() { letterPositions[1], letterPositions[3] };
            var two = new List<char>() { letterPositions[0], letterPositions[1], letterPositions[2], letterPositions[5], letterPositions[4] };
            var three = new List<char>() { letterPositions[0], letterPositions[1], letterPositions[2], letterPositions[3], letterPositions[4] };
            var four = new List<char>() { letterPositions[6], letterPositions[2], letterPositions[1], letterPositions[3] };
            var five = new List<char>() { letterPositions[0], letterPositions[6], letterPositions[2], letterPositions[3], letterPositions[4] };
            var six = new List<char>() { letterPositions[0], letterPositions[6], letterPositions[2], letterPositions[3], letterPositions[4], letterPositions[5] };
            var seven = new List<char>() { letterPositions[0], letterPositions[1], letterPositions[3] };
            var eight = new List<char>() { letterPositions[0], letterPositions[1], letterPositions[2], letterPositions[3], letterPositions[4], letterPositions[5], letterPositions[6] };
            var nine = new List<char>() { letterPositions[6], letterPositions[0], letterPositions[1], letterPositions[2], letterPositions[3], letterPositions[4] };

            return new List<string>()
            {
                new string(zero.ToArray()),
                new string(one.ToArray()),
                new string(two.ToArray()),
                new string(three.ToArray()),
                new string(four.ToArray()),
                new string(five.ToArray()),
                new string(six.ToArray()),
                new string(seven.ToArray()),
                new string(eight.ToArray()),
                new string(nine.ToArray()),
            };
        }

        static List<string> DiscoverNumberConfigs(IEnumerable<string> inputs)
        {

            var letterPositions = CalculateLetterPositions(inputs);
            return TranslatePositionsToStringConfigs(letterPositions);
        }

        static bool CheckStringIsNumber(string output, string numberConfig)
        {
            var found = true;
            foreach (var letter in output)
            {
                if (!numberConfig.Contains(letter))
                    found = false;
            }

            return found && output.Length == numberConfig.Length;
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
                            sum += numberConfigs.IndexOf(numberConfig);
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
