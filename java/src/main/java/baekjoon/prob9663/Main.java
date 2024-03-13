package baekjoon.prob9663;
// https://www.acmicpc.net/problem/9663
// n-queen
// brute force, backtracking

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;

public class Main {

    // queens[row] = column
    private static int[] queens;
    private static int result;
    private static int n;

    private static void backtrack(int row, int column) {
        if (row == n - 1) {
            result++;
            return;
        }
        queens[row] = column;
        for (int nextColumn = 0; nextColumn < n; nextColumn++) {
            if (isAvailable(row + 1, nextColumn)) {
                backtrack(row + 1, nextColumn);
            }
        }
    }

    private static boolean isAvailable(int row, int column) {
        for (int r = 0; r < row; r++) {
            if (column == queens[r] || Math.abs(column - queens[r]) == Math.abs(row - r)) {
                return false;
            }
        }
        return true;
    }

    public static void main(String[] args) {
        try (
                BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
                BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out))
        ) {

            n = Integer.parseInt(br.readLine());
            queens = new int[n];
            for (int c = 0; c < n; c++) {
                backtrack(0, c);
            }

            bw.write(String.valueOf(result));
            bw.newLine();
            bw.flush();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }

    }
}
