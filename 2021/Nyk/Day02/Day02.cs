using System;
using Utils;

namespace Day02
{
    internal class Day02
    {
        static int Part1()
        {
            var fileReader = new FileReader();
            var vector = new Vector(0, 0);

            foreach (var line in fileReader.ReadByLines("input.txt"))
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

        static int Part2()
        {
            var fileReader = new FileReader();
            var vector = new Vector(0, 0, 0);

            foreach (var line in fileReader.ReadByLines("input.txt"))
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

        static void Main(string[] args)
        {
            var result1 = Part1();
            var result2 = Part2();

            Console.WriteLine($"Part 1: {result1}");
            Console.WriteLine($"Part 2: {result2}");
        }
    }
}
