using System;
using System.Collections.Generic;
using System.Diagnostics.CodeAnalysis;
using System.Linq;

using Utils;

namespace Day11
{
    public class Day11
    {
        public static List<List<DumboOctopus>> ParseInput(IEnumerable<string> lines)
        {
            return lines.Select(line => line.Select(energy => new DumboOctopus(energy.ToInt())).ToList()).ToList();
        }

        static List<List<DumboOctopus>> IncreaseEnergyForStep(List<List<DumboOctopus>> octopuses)
        {
            foreach (var row in octopuses)
            {
                foreach (var octopus in row)
                {
                    octopus.IncreaseEnergy();
                }
            }

            return octopuses;
        }

        static List<List<DumboOctopus>> PerformFlash(List<List<DumboOctopus>> octopuses, DumboOctopus octopus, int row, int col)
        {
            if (octopus.Energy > 9 && octopus.FlashedThisStep == false)
            {
                var neighbours = new List<(DumboOctopus Octopus, Vector Position)>();

                octopus.IncreaseFlashes();
                octopus.SetFlashedThisStep(true);

                neighbours.Add(UsefulFuncs.GetAdjacentWithCoords(octopuses, row - 1, col)); //up
                neighbours.Add(UsefulFuncs.GetAdjacentWithCoords(octopuses, row + 1, col)); //down
                neighbours.Add(UsefulFuncs.GetAdjacentWithCoords(octopuses, row, col - 1)); //left
                neighbours.Add(UsefulFuncs.GetAdjacentWithCoords(octopuses, row, col + 1)); //right
                neighbours.Add(UsefulFuncs.GetAdjacentWithCoords(octopuses, row - 1, col - 1)); //topLeft
                neighbours.Add(UsefulFuncs.GetAdjacentWithCoords(octopuses, row - 1, col + 1)); //topRight
                neighbours.Add(UsefulFuncs.GetAdjacentWithCoords(octopuses, row + 1, col - 1)); //bottomLeft
                neighbours.Add(UsefulFuncs.GetAdjacentWithCoords(octopuses, row + 1, col + 1)); //bottomRight

                neighbours.RemoveAll(neighbour => neighbour.Octopus == null);

                foreach (var neighbour in neighbours)
                {
                    neighbour.Octopus.IncreaseEnergy();
                    PerformFlash(octopuses, neighbour.Octopus, neighbour.Position.X, neighbour.Position.Y);
                }
            }

            return octopuses;
        }

        static List<List<DumboOctopus>> PerformFlashes(List<List<DumboOctopus>> octopuses)
        {
            foreach (var row in octopuses)
            {
                var rowIndex = octopuses.IndexOf(row);
                foreach (var octopus in row)
                {
                    var colIndex = row.IndexOf(octopus);
                    octopuses = PerformFlash(octopuses, octopus, rowIndex, colIndex);
                }
            }

            return octopuses;
        }

        static List<List<DumboOctopus>> ResetFlashedOctopuses(List<List<DumboOctopus>> octopuses)
        {
            foreach (var row in octopuses)
            {
                foreach (var octopus in row)
                {
                    if (octopus.FlashedThisStep)
                    {
                        octopus.ResetEnergy();
                        octopus.SetFlashedThisStep(false);
                    }
                }
            }

            return octopuses;
        }

        public static long Part1(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var octopuses = ParseInput(fileReader.ReadAllLines(inputFile));

            for (var step = 0; step < 100; step++)
            {
                octopuses = IncreaseEnergyForStep(octopuses);
                octopuses = PerformFlashes(octopuses);
                octopuses = ResetFlashedOctopuses(octopuses);
            }

            return octopuses.Select(row => row.Select(octopus => octopus.Flashes).Sum()).Sum();
        }

        static bool CheckAllOctopusesFlashed(IEnumerable<IEnumerable<DumboOctopus>> octopuses)
        {
            return octopuses.Select(row => row.Where(octopus => octopus.FlashedThisStep == true).Count()).Sum() == 100;
        }

        public static int Part2(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var octopuses = ParseInput(fileReader.ReadAllLines(inputFile));

            var foundStep = false;
            var stepCount = 0;

            while (!foundStep)
            {
                octopuses = IncreaseEnergyForStep(octopuses);
                octopuses = PerformFlashes(octopuses);
                if (CheckAllOctopusesFlashed(octopuses))
                {
                    foundStep = true;
                }
                else
                {
                    octopuses = ResetFlashedOctopuses(octopuses);
                }
                stepCount++;
            }

            return stepCount;
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
