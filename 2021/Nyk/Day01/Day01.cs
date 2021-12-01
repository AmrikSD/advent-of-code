using System;
using Utils;
using System.Linq;
using System.Collections.Generic;

namespace Day01
{

    internal class Day01
    {
        static dynamic Part1(object[] args)
        {
            var fileReader = new FileReader();
            var increases = 0;
            var previousLine = 0;

            foreach (var line in fileReader.ReadByLines("input.txt"))
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

        static dynamic Part2(object[] args)
        {
            var fileReader = new FileReader();
            var fileInput = fileReader.ReadAllLines("input.txt").Select(int.Parse).ToList();

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

        static void Main(string[] args)
        {
            var timer = new Timer();
            var result1 = timer.TimeFunction(Part1);
            var result2 = timer.TimeFunction(Part2);

            Console.WriteLine($"Part 1: {result1.result}, Timings: Min: {result1.min}ms, Max: {result1.max}ms, Average: {result1.avg}ms");
            Console.WriteLine($"Part 2: {result2.result}, Timings: Min: {result2.min}ms, Max: {result2.max}ms, Average: {result2.avg}ms");
        }
    }
}
