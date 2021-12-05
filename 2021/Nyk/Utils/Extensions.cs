using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Utils
{
    public static class Extensions
    {
        public static void Deconstruct(this string[] stringVectors, out Vector vector1, out Vector vector2)
        {
            var startVector = stringVectors[0].Split(",");
            var endVector = stringVectors[1].Split(",");
            vector1 = new Vector(int.Parse(startVector[0]), int.Parse(startVector[1]));
            vector2 = new Vector(int.Parse(endVector[0]), int.Parse(endVector[1]));
        }
    }
}
