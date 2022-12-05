package xyz.amrik.adventofcode.Day;

import java.util.Stack;

public class DayFive extends Day {

    Stack<Character>[] stacks;

    Stack reverseStack(Stack stack) {
        Stack newStack = new Stack();
        while (!stack.empty()) {
            newStack.add(stack.pop());
        }
        return newStack;
    }

    public DayFive(String filePathString) {
        super(filePathString);
    }

    @Override
    public void partOne() {

        stacks = new Stack[(stringList.get(0).length() + 1) / 4];
        for (int i = 0; i < stacks.length; i++) {
            stacks[i] = new Stack();
        }

        int start = 1;
        int offset = 4;

        for (String string : stringList) {
            if (string.startsWith(" 1")) {
                break;
            }

            int curr = 0;
            char[] charArr = string.toCharArray();
            while (curr * offset + start < string.length()) {
                char c = charArr[curr * offset + start];
                if (c != ' ') {
                    stacks[curr].add(c);
                }
                curr++;

            }
        }

        for (int i = 0; i < stacks.length; i++) {
            stacks[i] = reverseStack(stacks[i]);
        }

        for (String string : stringList) {
            if (!string.startsWith("move")) {
                continue;
            }
            int count = Integer.valueOf(string.split(" ")[1]);
            int from = Integer.valueOf(string.split(" ")[3]) - 1; // need to -1 because arrays are 0 based indexed....
            int to = Integer.valueOf(string.split(" ")[5]) - 1;
            int curr = 0;

            while (curr < count) {
                stacks[to].push(stacks[from].pop());
                curr++;
            }

        }

        StringBuilder sb = new StringBuilder();

        for (Stack<Character> stack : stacks) {
            sb.append(stack.peek());
        }

        answerOne = sb.toString();

    }

    @Override
    public void partTwo() {
        stacks = new Stack[(stringList.get(0).length() + 1) / 4];
        for (int i = 0; i < stacks.length; i++) {
            stacks[i] = new Stack();
        }

        int start = 1;
        int offset = 4;

        for (String string : stringList) {
            if (string.startsWith(" 1")) {
                break;
            }

            int curr = 0;
            char[] charArr = string.toCharArray();
            while (curr * offset + start < string.length()) {
                char c = charArr[curr * offset + start];
                if (c != ' ') {
                    stacks[curr].add(c);
                }
                curr++;

            }
        }

        for (int i = 0; i < stacks.length; i++) {
            stacks[i] = reverseStack(stacks[i]);
        }

        for (String string : stringList) {
            if (!string.startsWith("move")) {
                continue;
            }
            int count = Integer.valueOf(string.split(" ")[1]);
            int from = Integer.valueOf(string.split(" ")[3]) - 1; // need to -1 because arrays are 0 based indexed....
            int to = Integer.valueOf(string.split(" ")[5]) - 1;
            int curr = 0;

            Stack<Character> charStack = new Stack();

            while (curr < count) {
                charStack.push(stacks[from].pop());
                curr++;
            }

            charStack = reverseStack(charStack);
            for (Character character : charStack) {
                stacks[to].push(character);
            }

        }

        StringBuilder sb = new StringBuilder();

        for (Stack<Character> stack : stacks) {
            sb.append(stack.peek());
        }

        answerTwo = sb.toString();

    }

}
