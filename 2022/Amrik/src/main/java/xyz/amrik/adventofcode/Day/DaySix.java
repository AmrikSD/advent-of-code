package xyz.amrik.adventofcode.Day;

import java.util.LinkedList;
import java.util.Set;
import java.util.TreeSet;

public class DaySix extends Day {

    public DaySix(String filePathString) {
        super(filePathString);
    }

    @Override
    public void partOne() {

        char[] input = stringList.get(0).toCharArray();
        int uniqueNeeded = 4;
        LinkedList<Character> seen = new LinkedList<>();

        for (int i = 0; i < input.length; i++) {
            seen.add(input[i]);

            if (i >= uniqueNeeded - 1) {
                Set unique = new TreeSet(seen);
                if (unique.size() == seen.size()) {
                    answerOne = String.valueOf(i + 1);
                    return;
                }
                seen.removeFirst();
            }

        }

    }

    @Override
    public void partTwo() {

        char[] input = stringList.get(0).toCharArray();
        int uniqueNeeded = 14;
        LinkedList<Character> seen = new LinkedList<>();

        for (int i = 0; i < input.length; i++) {
            seen.add(input[i]);

            if (i >= uniqueNeeded - 1) {
                Set unique = new TreeSet(seen);
                if (unique.size() == seen.size()) {
                    answerTwo = String.valueOf(i + 1);
                    return;
                }
                seen.removeFirst();
            }

        }

    }

}
