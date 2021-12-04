﻿using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Diagnostics.CodeAnalysis;
using System.IO;
using System.Linq;

namespace Utils
{
    public class FileReader
    {
        public IEnumerable<string> ReadByLines(string path)
        {
            return File.ReadLines(ConvertPath(path));
        }

        public IEnumerable<string> ReadAllLines(string path)
        {
            return File.ReadAllLines(ConvertPath(path)).ToList();
        }

        string ConvertPath(string path)
        {
            return "../../../" + path;
        }
    }

    public class Vector
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Aim { get; set; }

        public Vector(int x, int y, int aim = 0)
        {
            X = x;
            Y = y;
            Aim = aim;
        }

        public int MultiplyAttributes()
        {
            return X * Y;
        }
    }

    [ExcludeFromCodeCoverage]
    public class Timer
    {
        private Stopwatch _stopwatch = new();
        public (dynamic result, double min, double max, double avg) TimeFunction(Func<object[], dynamic> func, params object[] funcArgs)
        {
            var times = new List<double>();
            dynamic value = null;

            for (var _index = 0; _index < 10; _index++)
            {
                _stopwatch.Reset();
                _stopwatch.Start();
                value = func(funcArgs);
                _stopwatch.Stop();
                times.Add(_stopwatch.Elapsed.TotalMilliseconds);
            }

            return (value, times.Min(), times.Max(), times.Average());
        }
    }
}
