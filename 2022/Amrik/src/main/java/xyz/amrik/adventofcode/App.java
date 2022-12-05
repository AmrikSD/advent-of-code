package xyz.amrik.adventofcode;

import xyz.amrik.adventofcode.Day.*;

public class App {
    public static void main(String[] args) {
        Day[] days = new Day[] {
                new DayOne("./resources/day01.txt"),
                new DayTwo("./resources/day02.txt"),
                new DayThree("./resources/day03.txt"),
                new DayFour("./resources/day04.txt")
        };
        System.out.format("%3s%15s%15s%n", "Day", "Part 1", "Part 2");
        for (int i = 0; i < days.length; i++) {
            System.out.format("%3s%15s%15s%n", i + 1, days[i].answerOne, days[i].answerTwo);
        }
    }
}
