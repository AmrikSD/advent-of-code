using System.Collections.Generic;
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
}
