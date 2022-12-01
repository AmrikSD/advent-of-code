package xyz.amrik.adventofcode.Day;

import java.util.ArrayList;
import java.util.List;
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public abstract class Day implements ProblemSolver {

    List<String> stringList;
    public String answerOne;
    public String answerTwo;

    protected Day(String filePathString){
        popluateStringList(filePathString);
        partOne();
        partTwo();
    }

    // Give it a file and it will give you a list of strings, one for each line.
    // Idea being that you compose it with parseList() to get an even more useful list, but don't need to do this boring boilerplate each time.
    void popluateStringList(String fileName){
        List<String> tmpList = new ArrayList<>();
        try (BufferedReader reader = new BufferedReader(new FileReader(fileName))) {
            String line = reader.readLine();
            while(line!=null){
                tmpList.add(line);
                line = reader.readLine();
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
        stringList = tmpList;
    }

}
