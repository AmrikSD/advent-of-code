using System;
using System.Collections.Generic;
using System.Diagnostics.CodeAnalysis;
using System.Linq;
using Utils;

namespace Day10
{
    public class Day10
    {
        public static T ParseInput<T>(IEnumerable<string> lines)
        {
            throw new NotImplementedException();
        }

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

        public static T Part1<T>(string inputFile = "input.txt")
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

            var sum = bracketCount.Select(bracket => bracket.Value * bracketScore[bracket.Key]).Sum();
            return (T)Convert.ChangeType(sum, typeof(T));
        }

        public static T Part2<T>(string inputFile = "input.txt")
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
                        .Reverse<char>() //Reverse the last as the next needed was added last
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

            var result = incompleteScores.OrderByDescending(score => score).ElementAt(incompleteScores.Count / 2);

            return (T)Convert.ChangeType(result, typeof(T));
        }

        [ExcludeFromCodeCoverage]
        static void Main(string[] args)
        {
            var result1 = Part1<int>();
            var result2 = Part2<long>();

            Console.WriteLine($"Part 1: {result1}");
            Console.WriteLine($"Part 2: {result2}");
        }
    }
}
