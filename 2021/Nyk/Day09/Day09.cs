using System;
using System.Collections.Generic;
using System.Diagnostics.CodeAnalysis;
using System.Linq;
using Utils;

namespace Day09
{
    public class Day09
    {
        static IEnumerable<IEnumerable<int>> ParseInput(IEnumerable<string> lines)
        {
            return lines.Select(line => line.Select(numChar => int.Parse(numChar.ToString())));
        }

        static (int Value, Vector Position) GetAdjacentWithCoords(IEnumerable<IEnumerable<int>> heightMap, int rowIndex, int colIndex)
        {
            if (rowIndex < 0 || rowIndex >= heightMap.Count() ||
                colIndex < 0 || colIndex >= heightMap.ElementAt(0).Count())
                return (-1, new Vector(rowIndex, colIndex));
            else return (heightMap.ElementAt(rowIndex).ElementAt(colIndex), new Vector(rowIndex, colIndex));
        }

        static int GetAdjacent(IEnumerable<IEnumerable<int>> heightMap, int rowIndex, int colIndex)
        {
            if (rowIndex < 0 || rowIndex >= heightMap.Count() ||
                colIndex < 0 || colIndex >= heightMap.ElementAt(0).Count())
                return -1;
            else return heightMap.ElementAt(rowIndex).ElementAt(colIndex);
        }

        static List<(int Value, Vector Position)> CalculateLowPoints(IEnumerable<IEnumerable<int>> heightMap)
        {
            var lowPoints = new List<(int, Vector)>();
            for (var row = 0; row < heightMap.Count(); row++)
            {
                for (var col = 0; col < heightMap.ElementAt(row).Count(); col++)
                {
                    var lowestPointRow = new List<int>();

                    lowestPointRow.Add(GetAdjacent(heightMap, row - 1, col)); //up
                    lowestPointRow.Add(GetAdjacent(heightMap, row + 1, col)); //down
                    lowestPointRow.Add(GetAdjacent(heightMap, row, col - 1)); //left
                    lowestPointRow.Add(GetAdjacent(heightMap, row, col + 1)); //right

                    lowestPointRow.RemoveAll(ints => ints == -1);
                    var min = lowestPointRow.Min();

                    if (min > heightMap.ElementAt(row).ElementAt(col))
                        lowPoints.Add((heightMap.ElementAt(row).ElementAt(col) + 1, new Vector(row, col)));
                }
            }
            return lowPoints;
        }

        public static int Part1(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var heightMap = ParseInput(fileReader.ReadAllLines(inputFile));

            var lowPoints = CalculateLowPoints(heightMap);

            return lowPoints.Select(tuple => tuple.Value).Sum();
        }

        static List<Vector> DiscoverBasin(IEnumerable<IEnumerable<int>> heightMap, Vector lowPoint, List<Vector> visitedPositions = null)
        {
            var pointsInBasin = new List<Vector>();
            var neighbours = new List<(int Value, Vector Position)>();

            neighbours.Add(GetAdjacentWithCoords(heightMap, lowPoint.X - 1, lowPoint.Y)); //up
            neighbours.Add(GetAdjacentWithCoords(heightMap, lowPoint.X + 1, lowPoint.Y)); //down
            neighbours.Add(GetAdjacentWithCoords(heightMap, lowPoint.X, lowPoint.Y - 1)); //left
            neighbours.Add(GetAdjacentWithCoords(heightMap, lowPoint.X, lowPoint.Y + 1)); //right

            neighbours.RemoveAll(neighbour => neighbour.Item1 == -1);

            if (visitedPositions == null)
            {
                visitedPositions = new List<Vector>();
            }

            pointsInBasin.Add(lowPoint);
            visitedPositions.Add(lowPoint);

            foreach (var neighbour in neighbours)
            {
                var visited = visitedPositions
                    .Select(position => position == neighbour.Position)
                    .Where(isVisited => isVisited == true).Count() > 0;

                if (!visited)
                {
                    var neighbourPosition = neighbour.Position;
                    if (heightMap.ElementAt(neighbourPosition.X).ElementAt(neighbourPosition.Y) != 9)
                    {
                        pointsInBasin.AddRange(DiscoverBasin(heightMap, neighbourPosition, visitedPositions));
                    }
                }
            }

            return pointsInBasin;
        }

        static List<List<Vector>> CalculateBasins(IEnumerable<IEnumerable<int>> heightMap, List<Vector> lowPoints)
        {
            var basins = new List<List<Vector>>();
            foreach (var lowPoint in lowPoints)
            {
                basins.Add(DiscoverBasin(heightMap, lowPoint));
            }

            return basins;
        }

        public static long Part2(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var heightMap = ParseInput(fileReader.ReadAllLines(inputFile));

            var lowPoints = CalculateLowPoints(heightMap);
            var lowPointPositions = lowPoints.Select(tuple => tuple.Position).ToList();
            var basins = CalculateBasins(heightMap, lowPointPositions).OrderByDescending(basin => basin.Count());

            var topThreeBasins = basins.Take(3);

            return topThreeBasins.Select(basin => basin.Count).Aggregate((basin1, basin2) => basin1 * basin2);
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
