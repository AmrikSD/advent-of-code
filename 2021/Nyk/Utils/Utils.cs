using System.Collections.Generic;

namespace Utils
{
    public class FileReader
    {
        public IEnumerable<string> ReadAllLines(string path)
        {
            string[] lines = System.IO.File.ReadAllLines(path);
            foreach (var line in lines)
            {
                yield return line;
            }
        }
    }
}
