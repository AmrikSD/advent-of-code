using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Diagnostics.CodeAnalysis;
using System.Linq;

namespace Utils
{
    [ExcludeFromCodeCoverage]
    public class PerformanceTester
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
