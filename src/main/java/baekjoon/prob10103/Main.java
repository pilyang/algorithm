package baekjoon.prob10103;

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.util.Arrays;

public class Main {

    private static final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
    private static final BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));

    private static final int INITIAL_SCORE = 100;

    public static void main(String[] args) throws IOException {

        final int totalRound = Integer.parseInt(br.readLine());

        int changScore = INITIAL_SCORE;
        int sangScore = INITIAL_SCORE;

        for (int i = 0; i < totalRound; i++) {
            final int[] numbers = Arrays.stream(br.readLine().split(" "))
                    .mapToInt(Integer::parseInt)
                    .toArray();

            changScore = (numbers[0] < numbers[1]) ? changScore-numbers[1] : changScore;
            sangScore = (numbers[0] > numbers[1]) ? sangScore-numbers[0] : sangScore;
        }

        bw.write(String.valueOf(changScore));
        bw.newLine();
        bw.write(String.valueOf(sangScore));
        bw.close();
    }
}
