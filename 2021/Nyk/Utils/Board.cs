using System.Collections.Generic;
using System.Linq;

namespace Utils
{
    public class BoardNode
    {
        public bool Called { get; set; }
        public int Value { get; set; }

        public BoardNode(int value)
        {
            Called = false;
            Value = value;
        }
    }

    public class Board
    {
        public List<List<BoardNode>> BoardNodes { get; set; } = new();
        public Board(List<List<string>> boardInput)
        {
            foreach (var row in boardInput)
            {
                var newRow = new List<BoardNode>();
                foreach (var num in row)
                {
                    newRow.Add(new BoardNode(num.ToInt()));
                }
                BoardNodes.Add(newRow);
            }
        }

        public void UpdateValue(int calledNumber)
        {
            foreach (var row in BoardNodes)
            {
                foreach (var node in row)
                {
                    if (node.Value == calledNumber) node.Called = true;
                }
            }
        }

        bool CheckColumns()
        {
            for (var rowIndex = 0; rowIndex < BoardNodes.Count(); rowIndex++)
            {
                var count = 0;
                for (var colIndex = 0; colIndex < BoardNodes.Count(); colIndex++)
                {
                    if (BoardNodes.ElementAt(colIndex).ElementAt(rowIndex).Called == true) count++;
                }

                if (count == BoardNodes.Count()) return true;
            }
            return false;

        }

        bool CheckRows()
        {
            return BoardNodes.Count(row => row.Count(node => node.Called) == BoardNodes.Count()) > 0;
        }

        public int CalculateUncalledSum()
        {
            return BoardNodes.Select(row => row.Select(node => !node.Called ? node.Value : 0).Sum()).Sum();
        }

        public bool CheckBingo()
        {
            return CheckRows() || CheckColumns();
        }
    }
}
