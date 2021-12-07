using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Utils
{
    public class GridNode
    {
        public int LineCount { get; set; } = 0;
        public GridNode()
        {

        }
    }
    public class Grid
    {
        List<List<GridNode>> Nodes { get; set; } = new();

        public Grid()
        {
            for (var row = 0; row < 999; row++)
            {
                var rows = new List<GridNode>();
                for (var column = 0; column < 999; column++)
                {
                    rows.Add(new GridNode());
                }
                Nodes.Add(rows);
            }
        }

        Vector CalculateDirection(Vector startVector, Vector endVector)
        {
            var direction = new Vector(0, 0);

            if (startVector.X < endVector.X)
                direction.X = 1;
            else if (startVector.X > endVector.X)
                direction.X = -1;

            if (startVector.Y < endVector.Y)
                direction.Y = 1;
            else if (startVector.Y > endVector.Y)
                direction.Y = -1;

            return direction;
        }

        IEnumerable<Vector> CalculateVectorsToTraverse(Vector startVector, Vector endVector)
        {
            var xIndicies = new List<int>();
            var yIndicies = new List<int>();

            var direction = CalculateDirection(startVector, endVector);

            if (direction.X == 1)
            {
                for (var xIndex = startVector.X; xIndex <= endVector.X; xIndex++)
                {
                    xIndicies.Add(xIndex);
                }
            }
            else if (direction.X == -1)
            {
                for (var xIndex = startVector.X; xIndex >= endVector.X; xIndex--)
                {
                    xIndicies.Add(xIndex);
                }
            }

            if (direction.Y == 1)
            {
                for (var yIndex = startVector.Y; yIndex <= endVector.Y; yIndex++)
                {
                    yIndicies.Add(yIndex);
                }
            }
            else if (direction.Y == -1)
            {
                for (var yIndex = startVector.Y; yIndex >= endVector.Y; yIndex--)
                {
                    yIndicies.Add(yIndex);
                }
            }

            return xIndicies.Zip(yIndicies, (x, y) => new Vector(x, y));
        }

        public void TraverseGridDiagonal(Vector startVector, Vector endVector)
        {
            var nodesToTraverse = CalculateVectorsToTraverse(startVector, endVector);
            foreach (var node in nodesToTraverse)
            {
                Nodes.ElementAt(node.X).ElementAt(node.Y).LineCount++;
            }
        }

        public void TraverseGridHorizontalAndVertical(Vector startVector, Vector endVector)
        {
            var direction = CalculateDirection(startVector, endVector);

            if (direction.X == 1 && (direction.Y == 1 || direction.Y == 0))
            {
                for (var xIndex = startVector.X; xIndex <= endVector.X; xIndex++)
                {
                    for (var yIndex = startVector.Y; yIndex <= endVector.Y; yIndex++)
                    {
                        Nodes.ElementAt(xIndex).ElementAt(yIndex).LineCount++;
                    }
                }
            }

            else if (direction.X == -1 && (direction.Y == 1 || direction.Y == 0))
            {
                for (var xIndex = startVector.X; xIndex >= endVector.X; xIndex--)
                {
                    for (var yIndex = startVector.Y; yIndex <= endVector.Y; yIndex++)
                    {
                        Nodes.ElementAt(xIndex).ElementAt(yIndex).LineCount++;
                    }
                }
            }

            else if (direction.Y == 1 && (direction.X == 1 || direction.X == 0))
            {
                for (var xIndex = startVector.X; xIndex <= endVector.X; xIndex++)
                {
                    for (var yIndex = startVector.Y; yIndex <= endVector.Y; yIndex++)
                    {
                        Nodes.ElementAt(xIndex).ElementAt(yIndex).LineCount++;
                    }
                }
            }

            else if (direction.Y == -1 && (direction.X == 1 || direction.X == 0))
            {
                for (var xIndex = startVector.X; xIndex <= endVector.X; xIndex++)
                {
                    for (var yIndex = startVector.Y; yIndex >= endVector.Y; yIndex--)
                    {
                        Nodes.ElementAt(xIndex).ElementAt(yIndex).LineCount++;
                    }
                }
            }

            else if (direction.X == -1 && direction.Y == -1)
            {
                for (var xIndex = startVector.X; xIndex >= endVector.X; xIndex--)
                {
                    for (var yIndex = startVector.Y; yIndex >= endVector.Y; yIndex--)
                    {
                        Nodes.ElementAt(xIndex).ElementAt(yIndex).LineCount++;
                    }
                }
            }
        }

        public int GetNodeGreaterThanTwo()
        {
            return Nodes.Sum(row => row.Count(node => node.LineCount >= 2));
        }
    }
}
