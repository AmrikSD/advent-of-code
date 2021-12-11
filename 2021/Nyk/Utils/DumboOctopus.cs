using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Utils
{
    public class DumboOctopus
    {
        public int Energy { get; set; } = 0;
        public bool FlashedThisStep { get; set; } = false;
        public long Flashes { get; set; } = 0;

        public DumboOctopus() { }
        public DumboOctopus(int energyValue)
        {
            Energy = energyValue;
        }

        public void IncreaseEnergy()
        {
            Energy++;
        }

        public void ResetEnergy()
        {
            Energy = 0;
        }

        public void SetFlashedThisStep(bool status)
        {
            FlashedThisStep = status;
        }

        public void IncreaseFlashes()
        {
            Flashes++;
        }


    }
}
