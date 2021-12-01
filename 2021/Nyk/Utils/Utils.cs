using System.Collections.Generic;

namespace Utils
{
    public class FileReader
    {
        public IEnumerable<string> ReadByLines(string path)
        {
            foreach (var line in System.IO.File.ReadLines(path))
            {
                yield return line;
            }
        }
    }
}
