﻿using System;
using System.Collections.Generic;
using System.Diagnostics.CodeAnalysis;
using System.Linq;
using Utils;

namespace Day03
{
    public class Day03
    {
        static IEnumerable<IEnumerable<char>> ConvertInput(IEnumerable<string> input)
        {
            return input.ElementAt(0)
                .Select(
                    (_bit, index) => input
                    .Select(
                        byteString => byteString.ElementAt(index)
                    )
                );
        }

        static int ConvertBinaryStringToInt(string binary)
        {
            var binaryValues = binary.Select((bit, index) => index == 0 ? 1 : ((int)Math.Pow(2, index)));
            var reversedString = new string(binary.ToCharArray().Reverse().ToArray());

            return reversedString.Select((bit, index) => bit == '1' ? binaryValues.ElementAt(index) : 0).Sum();
        }

        static string GenerateBinaryString(IEnumerable<IEnumerable<char>> values)
        {
            return string.Join("",
                values.Select(line => line.Count(bit => bit == '1') > line.Count(bit => bit == '0') ? "1" : "0")
            );
        }

        static int CalculateGammaRate(IEnumerable<IEnumerable<char>> values, out string gammaString)
        {
            gammaString = GenerateBinaryString(values);
            return ConvertBinaryStringToInt(gammaString);
        }

        static string InvertBinaryString(string binary)
        {
            return new string(binary.Select(bit => bit == '1' ? '0' : '1').ToArray());
        }

        static int CalculateEpsilonRate(string gammaString)
        {
            var epsilonString = InvertBinaryString(gammaString);
            return ConvertBinaryStringToInt(epsilonString);
        }

        public static int Part1(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var lines = fileReader.ReadAllLines(inputFile);

            var values = ConvertInput(lines);

            var gammaValue = CalculateGammaRate(values, out var gammaString);
            var epsilonValue = CalculateEpsilonRate(gammaString);

            return gammaValue * epsilonValue;
        }

        static int CalculateOxygenRate(IEnumerable<string> lines, int binaryIndex)
        {
            var values = ConvertInput(lines);

            var count1Bits = values.ElementAt(binaryIndex).Count(bit => bit == '1');
            var count0Bits = values.ElementAt(binaryIndex).Count(bit => bit == '0');

            char mostCommonBit;

            if (count1Bits == count0Bits)
            {
                mostCommonBit = '1';
            }
            else
            {
                mostCommonBit = count1Bits > count0Bits ? '1' : '0';
            }

            var newLines = lines.Where(binary => binary.ElementAt(binaryIndex) == mostCommonBit);

            if (newLines.Count() > 1)
            {
                return CalculateOxygenRate(newLines, binaryIndex + 1);
            }

            return ConvertBinaryStringToInt(newLines.ElementAt(0));
        }

        static int CalculateCO2Rate(IEnumerable<string> lines, int binaryIndex)
        {
            var values = ConvertInput(lines);

            var count1Bits = values.ElementAt(binaryIndex).Count(bit => bit == '1');
            var count0Bits = values.ElementAt(binaryIndex).Count(bit => bit == '0');

            char leastCommonBit;

            if (count1Bits == count0Bits)
            {
                leastCommonBit = '0';
            }
            else
            {
                leastCommonBit = count1Bits < count0Bits ? '1' : '0';
            }

            var newLines = lines.Where(binary => binary.ElementAt(binaryIndex) == leastCommonBit);

            if (newLines.Count() > 1)
            {
                return CalculateCO2Rate(newLines, binaryIndex + 1);
            }

            return ConvertBinaryStringToInt(newLines.ElementAt(0));
        }

        public static int Part2(string inputFile = "input.txt")
        {
            var fileReader = new FileReader();
            var lines = fileReader.ReadAllLines(inputFile);

            var oxygenValue = CalculateOxygenRate(lines, 0);
            var co2Value = CalculateCO2Rate(lines, 0);

            return oxygenValue * co2Value;
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
