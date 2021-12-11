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
            vector1 = new Vector(startVector[0].ToInt(), startVector[1].ToInt());
            vector2 = new Vector(endVector[0].ToInt(), endVector[1].ToInt());
        }

        public static int ToInt(this string number)
        {
            return int.Parse(number);
        }

        public static int ToInt(this char number)
        {
            return int.Parse(number.ToString());
        }
    }
}
