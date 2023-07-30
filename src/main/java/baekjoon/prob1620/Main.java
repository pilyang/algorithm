package baekjoon.prob1620;
// https://www.acmicpc.net/problem/1620

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class Main {

    private static final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

    public static void main(String[] args) throws IOException {
        final String[] firstLineInput = br.readLine().split(" ");
        final int pockNum = Integer.parseInt(firstLineInput[0]);
        final int probNum = Integer.parseInt(firstLineInput[1]);

        final Map<String, String> nameToNum = new HashMap<>();
        final Map<Integer, String> numToName = new HashMap<>();

        for (int number = 1; number <= pockNum; number++) {
            final String pockName = br.readLine();
            nameToNum.put(pockName, String.valueOf(number));
            numToName.put(number, pockName);
        }

        final String[] answers = new String[probNum];

        for (int i = 0; i < probNum; i++) {
            final String prob = br.readLine();
            answers[i] = isNumber(prob) ? numToName.get(Integer.parseInt(prob)) : nameToNum.get(prob);
        }

        Arrays.stream(answers).forEach(System.out::println);
    }

    private static boolean isNumber(final String value) {
        try {
            Integer.parseInt(value);
            return true;
        } catch (NumberFormatException e) {
            return false;
        }
    }
}
