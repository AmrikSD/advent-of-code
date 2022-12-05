package xyz.amrik.adventofcode.Day;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.TreeSet;

public class DayThree extends Day {

    public DayThree(String filePathString) {
        super(filePathString);
    }

    int priority(char c) {
        if (c < 97) {
            return c - 64 + 26;
        }
        return c - 96;
    }

    int priorityOfItemInBothBags(String str) {
        Set<Character> charSet = new HashSet<>();
        final int mid = str.length() / 2;
        char[] firstHalf = str.substring(0, mid).toCharArray();
        char[] secondHalf = str.substring(mid).toCharArray();

        for (char c : firstHalf) {
            charSet.add(c);
        }

        for (char c : secondHalf) {
            if (charSet.contains(c)) {
                return priority(c);
            }

        }

        System.err.println("something went bigly badly");
        return 0; // should never hit this

    }

    int priorityOfItemInAllGroupsBags(List<String> elvesBags) {

        Set<Character> charSet = new TreeSet<>();

        for (char c = 'a'; c <= 'z'; c++) {
            charSet.add(c);
        }
        for (char c = 'A'; c <= 'Z'; c++) {
            charSet.add(c);
        }

        for (String str : elvesBags) {
            char[] charArr = str.toCharArray();
            Set<Character> elvSet = new TreeSet();
            for (Character c : charArr) {
                elvSet.add(c);
            }
            charSet.retainAll(elvSet);
        }

        Character[] charArr = new Character[1];
        charArr = charSet.toArray(charArr);

        return priority(charArr[0]);

    }

    @Override
    public void partOne() {
        answerOne = stringList.stream().map(this::priorityOfItemInBothBags).reduce(0, Integer::sum).toString();
    }

    @Override
    public void partTwo() {
        int c = 0;
        int sum = 0;
        List<String> group = new ArrayList<>();
        for (String str : stringList) {
            c++;
            group.add(str);
            if (c == 3) {
                sum += priorityOfItemInAllGroupsBags(group);
                group.removeAll(group);
                c = 0;
            }
        }
        answerTwo = String.valueOf(sum);

    }
}
