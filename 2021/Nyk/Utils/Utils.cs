using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;

namespace Utils
{
    public class FileReader
    {
        public IEnumerable<string> ReadByLines(string path)
        {
            foreach (var line in File.ReadLines(ConvertPath(path)))
            {
                yield return line;
            }
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
