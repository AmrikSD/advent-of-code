using System;
using Utils;

namespace Day01
{

    internal class Program
    {
        static void Main(string[] args)
        {
            var fileReader = new FileReader();
            var increases = 0;
            var previousLine = 0;

            foreach (var line in fileReader.ReadAllLines("input.txt"))
            {
                var lineNum = Int32.Parse(line);
                if (previousLine != 0)
                {
                    if (previousLine < lineNum)
                    {
                        increases++;
                    }
                }
                previousLine = lineNum;
            }

            Console.WriteLine(increases);
            Console.ReadLine();
        }
    }
}
