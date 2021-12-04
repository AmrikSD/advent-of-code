using System;
using Utils;
using System.Linq;
using System.Collections.Generic;
using System.Diagnostics.CodeAnalysis;

namespace Day01
{
    public class Day01
    {
        public static int Part1(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var increases = 0;
            var previousLine = 0;

            foreach (var line in fileReader.ReadByLines(inputFile))
            {
                var lineNum = int.Parse(line);
                if (previousLine != 0)
                {
                    if (previousLine < lineNum)
                    {
                        increases++;
                    }
                }
                previousLine = lineNum;
            }

            return increases;
        }

        public static int Part2(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var fileInput = fileReader.ReadAllLines(inputFile).Select(int.Parse).ToList();

            var slidingWindows = new List<int>();
            var increases = 0;

            for (var index = 2; index < fileInput.Count; index++)
            {
                var sum = 0;
                for (var slice = index; slice > index - 3; slice--)
                {
                    sum += fileInput.ElementAt(slice);
                }
                slidingWindows.Add(sum);

                if (slidingWindows.Count > 1)
                {
                    if (slidingWindows.Last() > slidingWindows[slidingWindows.Count - 2])
                    {
                        increases++;
                    }
                }

            }

            return increases;
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
