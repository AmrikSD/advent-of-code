namespace Utils
{
    public class Vector
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Aim { get; set; }

        public Vector(int x, int y, int aim = 0)
        {
            X = x;
            Y = y;
            Aim = aim;
        }

        public int MultiplyAttributes()
        {
            return X * Y;
        }
    }
}
