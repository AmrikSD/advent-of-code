using System;
using Utils;
using System.Linq;
using System.Collections.Generic;

namespace Day01
{

    internal class Program
    {
        static int Part1()
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

        static int Part2()
        {
            var fileReader = new FileReader();
            var fileInput = fileReader.ReadByLines("input.txt").ToList().Select(int.Parse).ToList();

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
            Console.WriteLine($"Part 1: {Part1()}");
            Console.WriteLine($"Part 2: {Part2()}");
            Console.ReadLine();
        }
    }
}
