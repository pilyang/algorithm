package baekjoon.prob21771;
// https://www.acmicpc.net/problem/21771

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.util.Arrays;

public class Main {

    private static final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
    private static final BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));

    public static void main(String[] args) throws IOException {
        final int[] mapSize = Arrays.stream(br.readLine().split(" "))
                .mapToInt(Integer::parseInt)
                .toArray();

        final int[] objectSizeInput = Arrays.stream(br.readLine().split(" "))
                .mapToInt(Integer::parseInt)
                .toArray();

        final int[] gaSize = {objectSizeInput[0], objectSizeInput[1]};
        final int[] piSize = {objectSizeInput[2], objectSizeInput[3]};

        final int requiredPCount = piSize[0]*piSize[1];
        int actualPCount = 0;

        for (int rowIndex = 0; rowIndex < mapSize[0]; rowIndex++) {
            final String columInput = br.readLine();
            final String[] colum = columInput.split("");
            final long count = Arrays.stream(colum)
                    .filter(tile -> tile.equals("P"))
                    .count();
            actualPCount += count;
        }

        int result = (actualPCount == requiredPCount) ? 0 : 1;
        bw.write(String.valueOf(result));
        bw.close();
    }
}
