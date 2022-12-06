package xyz.amrik.adventofcode.Day;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

enum Move {
    ROCK, // Lose
    PAPER, // Draw
    SCISSORS; // Win

    Move losesTo() {
        switch (this) {
            case ROCK:
                return PAPER;
            case PAPER:
                return SCISSORS;
            case SCISSORS:
                return ROCK;
            default:
                return null;
        }
    }

    Move winsAgainst() {
        return this.losesTo().losesTo();
    }

}

class Round {
    Move opponentMove;
    Move playerMove;
}

public class DayTwo extends Day {

    private List<Round> roundsList() {
        List<Round> rounds = new ArrayList<>();

        for (String str : stringList) {

            Round round = new Round();

            String strMoveOpp = str.split(" ")[0];
            String strMovePlayer = str.split(" ")[1];

            switch (strMoveOpp) {
                case "A":
                    round.opponentMove = Move.ROCK;
                    break;
                case "B":
                    round.opponentMove = Move.PAPER;
                    break;
                case "C":
                    round.opponentMove = Move.SCISSORS;
                    break;
                default:
                    break;
            }
            // DO player now
            switch (strMovePlayer) {
                case "X":
                    round.playerMove = Move.ROCK;
                    break;
                case "Y":
                    round.playerMove = Move.PAPER;
                    break;
                case "Z":
                    round.playerMove = Move.SCISSORS;
                    break;
                default:
                    break;
            }
            rounds.add(round);
        }
        return rounds;
    }

    private int scoreForRound(Round round) {

        Move playerMove = round.playerMove;
        Move oppMove = round.opponentMove;

        int choiceScore = playerMove.ordinal() + 1;

        int outcomeScore = 0;

        if (playerMove.equals(oppMove)) {
            outcomeScore = 3;
        } else {
            if (playerWon(round)) {
                outcomeScore = 6;
            } else {
                outcomeScore = 0;
            }
        }

        return choiceScore + outcomeScore;
    }

    // We only call this if playerMove and oppMove are different.
    private boolean playerWon(Round round) {
        switch (round.playerMove) {
            case ROCK:
                return round.opponentMove == Move.SCISSORS;
            case PAPER:
                return round.opponentMove == Move.ROCK;
            case SCISSORS:
                return round.opponentMove == Move.PAPER;
            default:
                return false; // We should not end up in here ever...
        }
    }

    public DayTwo(String filePathString) {
        super(filePathString);
    }

    @Override
    public void partOne() {
        List<Round> rounds = roundsList();
        answerOne = rounds.stream().collect(Collectors.summingInt(this::scoreForRound)).toString();
    }

    @Override
    public void partTwo() {
        List<Round> rounds = roundsList();
        int sum = 0;
        for (Round round : rounds) {

            int outcomeScore = round.playerMove.ordinal() * 3;
            int choiceScore = 0;

            switch (round.playerMove) {
                case PAPER:
                    choiceScore = round.opponentMove.ordinal() + 1;
                    break;
                case ROCK:
                    choiceScore = round.opponentMove.winsAgainst().ordinal() + 1;
                    break;
                case SCISSORS:
                    choiceScore = round.opponentMove.losesTo().ordinal() + 1;
                    break;
                default:
                    System.out.println("fucked up");
                    break;

            }
            sum += choiceScore + outcomeScore;

        }
        answerTwo = String.valueOf(sum);

    }

}
