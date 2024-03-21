package baekjoon.prob1915;
// https://www.acmicpc.net/problem/1915
// dp

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.util.StringTokenizer;

public class Main {

    public static void main(String[] args) {
        try (
                BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
                BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out))
        ) {
            StringTokenizer st = new StringTokenizer(br.readLine());
            int row = Integer.parseInt(st.nextToken());
            int column = Integer.parseInt(st.nextToken());

            int maxSize = 0;
            int[][] possibleSize = new int[row + 1][column + 1];
            for (int r = 1; r <= row; r++) {
                char[] rowInfo = br.readLine().toCharArray();
                for (int c = 1; c <= column; c++) {
                    if (rowInfo[c - 1] == '1') {
                        possibleSize[r][c] = Math.min(Math.min(possibleSize[r - 1][c], possibleSize[r][c - 1]), possibleSize[r - 1][c - 1]) + 1;
                        if (maxSize < possibleSize[r][c]) {
                            maxSize = possibleSize[r][c];
                        }
                    }
                }
            }

            bw.write(String.valueOf(maxSize * maxSize));
            bw.newLine();
            bw.flush();

        } catch (IOException e) {
            throw new RuntimeException(e);
        }

    }
}
