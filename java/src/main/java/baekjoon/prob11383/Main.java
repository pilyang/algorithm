package baekjoon.prob11383;
// https://www.acmicpc.net/problem/11383

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.util.Arrays;

public class Main {

    private static final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
    private static final BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));

    private static final String EYFA = "Eyfa";
    private static final String NOT_EYFA = "Not Eyfa";

    public static void main(String[] args) throws IOException {

        final int[] size = Arrays.stream(br.readLine().split(" "))
                .mapToInt(Integer::parseInt)
                .toArray();

        String[][] original = new String[size[0]][size[1]];
        String[][] n2 = new String[size[0]][size[1] * 2];

        for (int row = 0; row < size[0]; row++) {
            original[row] = br.readLine().split("");
        }
        for (int row = 0; row < size[0]; row++) {
            n2[row] = br.readLine().split("");
        }

        for (int row = 0; row < size[0]; row++) {
            for (int colum = 0; colum < size[1]; colum++) {
                int targetColum = colum * 2;
                final String originalStr = original[row][colum];
                if (!originalStr.equals(n2[row][targetColum]) || !originalStr.equals(n2[row][targetColum + 1])) {
                    bw.write(NOT_EYFA);
                    bw.close();
                    return;
                }
            }
        }

        bw.write(EYFA);
        bw.close();
    }
}
