package xyz.amrik.adventofcode.Day;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class DayOne extends Day {

    List<Integer> elfCalories() {
        int curr = 0;
        List<Integer> ec = new ArrayList<>();
        for (String str : stringList) {
            if (!str.equals("")) {
                curr += Integer.valueOf(str);
            } else {
                ec.add(curr);
                curr = 0;
            }
        }
        Collections.sort(ec);
        return ec;
    }

    public DayOne(String filePathString) {
        super(filePathString);

    }

    @Override
    public void partOne() {
        answerOne = String.valueOf(Collections.max(elfCalories()));
    }

    @Override
    public void partTwo() {

        // Since we sort in elfCalories() we can just take the last 3.
        List<Integer> topThree = elfCalories().subList(elfCalories().size() - 3, elfCalories().size());

        int sumOfTopThree = topThree.stream().mapToInt(Integer::intValue).sum();

        answerTwo = String.valueOf(sumOfTopThree);

    }

}
