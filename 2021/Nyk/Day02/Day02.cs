using System;
using System.Diagnostics.CodeAnalysis;
using Utils;

namespace Day02
{
    public class Day02
    {
        public static int Part1(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var vector = new Vector(0, 0);

            foreach (var line in fileReader.ReadByLines(inputFile))
            {
                var directionAmount = line.Split(' ');
                var amount = int.Parse(directionAmount[1]);

                switch (directionAmount[0])
                {
                    case "forward":
                        vector.X += amount;
                        break;
                    case "down":
                        vector.Y += amount;
                        break;
                    case "up":
                        vector.Y -= amount;
                        break;
                }
            }

            return vector.MultiplyAttributes();
        }

        public static int Part2(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var vector = new Vector(0, 0, 0);

            foreach (var line in fileReader.ReadByLines(inputFile))
            {
                var directionAmount = line.Split(' ');
                var amount = int.Parse(directionAmount[1]);

                switch (directionAmount[0])
                {
                    case "forward":
                        vector.X += amount;
                        vector.Y += vector.Aim * amount;
                        break;
                    case "down":
                        vector.Aim += amount;
                        break;
                    case "up":
                        vector.Aim -= amount;
                        break;
                }
            }

            return vector.MultiplyAttributes();
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
