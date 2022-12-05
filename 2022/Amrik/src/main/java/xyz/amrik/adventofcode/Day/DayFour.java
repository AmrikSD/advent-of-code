package xyz.amrik.adventofcode.Day;

class Area {
    int start;
    int end;

    Boolean contains(Area otherArea) {
        return start <= otherArea.start && end >= otherArea.end;
    }

    Boolean overlaps(Area otherArea) {
        if (start >= otherArea.start && start <= otherArea.end) {
            return true;
        }
        if (end >= otherArea.start && end <= otherArea.end) {
            return true;
        }
        if (otherArea.start >= start && otherArea.start <= end) {
            return true;
        }
        if (otherArea.end >= start && otherArea.end <= end) {
            return true;
        }
        return false;
    }
}

public class DayFour extends Day {

    public DayFour(String filePathString) {
        super(filePathString);
    }

    @Override
    public void partOne() {
        int sum = 0;
        for (String str : stringList) {

            String[] split = str.split(",");

            String areaStrOne = split[0];
            String[] areaOneLowHigh = areaStrOne.split("-");

            String areaStrTwo = split[1];
            String[] areaTwoLowHigh = areaStrTwo.split("-");

            Area areaOne = new Area();
            areaOne.start = Integer.parseInt(areaOneLowHigh[0]);
            areaOne.end = Integer.parseInt(areaOneLowHigh[1]);

            Area areaTwo = new Area();
            areaTwo.start = Integer.parseInt(areaTwoLowHigh[0]);
            areaTwo.end = Integer.parseInt(areaTwoLowHigh[1]);

            if (areaOne.contains(areaTwo) || areaTwo.contains(areaOne)) {
                sum++;
            }

        }

        answerOne = String.valueOf(sum);

    }

    @Override
    public void partTwo() {

        int sum = 0;
        for (String str : stringList) {

            String[] split = str.split(",");

            String areaStrOne = split[0];
            String[] areaOneLowHigh = areaStrOne.split("-");

            String areaStrTwo = split[1];
            String[] areaTwoLowHigh = areaStrTwo.split("-");

            Area areaOne = new Area();
            areaOne.start = Integer.parseInt(areaOneLowHigh[0]);
            areaOne.end = Integer.parseInt(areaOneLowHigh[1]);

            Area areaTwo = new Area();
            areaTwo.start = Integer.parseInt(areaTwoLowHigh[0]);
            areaTwo.end = Integer.parseInt(areaTwoLowHigh[1]);

            if (areaOne.overlaps(areaTwo)) {
                sum++;
            }

        }

        answerTwo = String.valueOf(sum);
    }

}
