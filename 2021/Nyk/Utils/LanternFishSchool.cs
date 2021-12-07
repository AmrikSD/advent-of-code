using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Utils
{
	//Basically Amrik's solution, Lifted and Shifted. Thanks Amrik
    public class LanternFishSchool
    {
        Dictionary<long, long> fishSchool;

        public LanternFishSchool(IEnumerable<int> startingFish)
        {
            fishSchool = new Dictionary<long, long>()
            {
                {0, 0},
                {1, 0},
                {2, 0},
                {3, 0},
                {4, 0},
                {5, 0},
                {6, 0},
                {7, 0},
                {8, 0},
            };

            foreach (var fish in startingFish)
            {
                fishSchool[fish]++;
            }
        }

        public void IterateDay()
        {
            var nextDay = new Dictionary<long, long>(fishSchool);

            //Thanks Amrik, this makes sense why we reverse it.
            var reverseKeys = fishSchool.Keys.Reverse();

            foreach (var key in reverseKeys)
            {
                if (key == 0)
                {
                    nextDay[8] = fishSchool[0];
                    nextDay[6] += fishSchool[0];
                }
                else
                {
                    nextDay[key - 1] = fishSchool[key];
                }
            }

            fishSchool = nextDay;
        }

        public long CountFish()
        {
            return fishSchool.Select((keyValuePair, _index) => keyValuePair.Value).Sum();
        }
    }
}
