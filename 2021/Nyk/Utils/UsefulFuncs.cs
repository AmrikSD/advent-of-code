using System;
using System.Collections.Generic;
using System.Linq;

namespace Utils
{
    static public class UsefulFuncs
    {
        static public (DumboOctopus octopus, Vector Position) GetAdjacentWithCoords(IEnumerable<IEnumerable<DumboOctopus>> collection, int rowIndex, int colIndex)
        {
            if (rowIndex < 0 || rowIndex >= collection.Count() ||
                colIndex < 0 || colIndex >= collection.ElementAt(0).Count())
                return (null, new Vector(rowIndex, colIndex));
            else
                return (collection.ElementAt(rowIndex).ElementAt(colIndex), new Vector(rowIndex, colIndex));
        }

        static public (int, Vector Position) GetAdjacentWithCoords(IEnumerable<IEnumerable<int>> collection, int rowIndex, int colIndex)
        {
            if (rowIndex < 0 || rowIndex >= collection.Count() ||
                colIndex < 0 || colIndex >= collection.ElementAt(0).Count())
                return (-1, new Vector(rowIndex, colIndex));
            else
                return (collection.ElementAt(rowIndex).ElementAt(colIndex), new Vector(rowIndex, colIndex));
        }

        static public int GetAdjacent(IEnumerable<IEnumerable<int>> collection, int rowIndex, int colIndex)
        {
            if (rowIndex < 0 || rowIndex >= collection.Count() ||
                colIndex < 0 || colIndex >= collection.ElementAt(0).Count())
                return -1;
            else return collection.ElementAt(rowIndex).ElementAt(colIndex);
        }
    }
}
