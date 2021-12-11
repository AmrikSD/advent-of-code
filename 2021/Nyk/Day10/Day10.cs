using System;
using System.Collections.Generic;
using System.Diagnostics.CodeAnalysis;
using System.Linq;
using Utils;

namespace Day10
{
    public class Day10
    {

        static bool CheckCorruptString(List<List<char>> startEndCollection, char character)
        {
            switch (character)
            {
                case '(':
                    startEndCollection.ElementAt(0).Add(character);
                    startEndCollection.ElementAt(1).Add(')');
                    return false;
                case '[':
                    startEndCollection.ElementAt(0).Add(character);
                    startEndCollection.ElementAt(1).Add(']');
                    return false;
                case '{':
                    startEndCollection.ElementAt(0).Add(character);
                    startEndCollection.ElementAt(1).Add('}');
                    return false;
                case '<':
                    startEndCollection.ElementAt(0).Add(character);
                    startEndCollection.ElementAt(1).Add('>');
                    return false;
                default:
                    if (startEndCollection.ElementAt(1).Last() == character)
                    {
                        startEndCollection.ElementAt(0).RemoveAt(startEndCollection.ElementAt(0).Count - 1);
                        startEndCollection.ElementAt(1).RemoveAt(startEndCollection.ElementAt(1).Count - 1);
                        return false;
                    }
                    else return true;
            }
        }

        public static int Part1(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var bracketCount = new Dictionary<char, int>()
            {
                {'(', 0 },
                {'[', 0 },
                {'{', 0 },
                {'<', 0 },
            };

            var bracketScore = new Dictionary<char, int>()
            {
                {'(', 3 },
                {'[', 57 },
                {'{', 1197 },
                {'<', 25137 },
            };


            foreach (var line in fileReader.ReadByLines(inputFile))
            {
                var startEndCollection = new List<List<char>>()
                {
                    new List<char> { },
                    new List<char> { },
                };

                foreach (var character in line)
                {
                    var opens = new List<char>() { '(', '{', '[', '<' };
                    var closes = new List<char>() { ')', '}', ']', '>' };
                    var corruptString = CheckCorruptString(startEndCollection, character);
                    if (corruptString)
                    {
                        bracketCount[opens.ElementAt(closes.IndexOf(character))]++;
                        break;
                    }
                }
            }

            return bracketCount.Select(bracket => bracket.Value * bracketScore[bracket.Key]).Sum();
        }

        public static long Part2(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();

            var bracketScore = new Dictionary<char, int>()
            {
                {')', 1 },
                {']', 2 },
                {'}', 3 },
                {'>', 4 },
            };

            var incompleteScores = new List<long>();
            foreach (var line in fileReader.ReadByLines(inputFile))
            {
                var startEndCollection = new List<List<char>>()
                {
                    new List<char> { },
                    new List<char> { },
                };

                var corruptString = false;

                foreach (var character in line)
                {
                    var opens = new List<char>() { '(', '{', '[', '<' };
                    var closes = new List<char>() { ')', '}', ']', '>' };

                    corruptString = CheckCorruptString(startEndCollection, character);
                    if (corruptString)
                    {
                        break;
                    }
                }

                if (!corruptString)
                {
                    var scoreValues = startEndCollection
                        .ElementAt(1) //The ending list
                        .Reverse<char>() //Reverse the list as the next needed was added last
                        .Select(character => bracketScore[character]);

                    long totalScoreForLine = 0;
                    foreach (var score in scoreValues)
                    {
                        totalScoreForLine *= 5;
                        totalScoreForLine += score;
                    }

                    incompleteScores.Add(totalScoreForLine);
                }
            }

            return incompleteScores.OrderByDescending(score => score).ElementAt(incompleteScores.Count / 2);
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
