package baekjoon.prob23933;
// https://www.acmicpc.net/problem/23933
// greedy, dp, knapsack

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
            final int tc = Integer.parseInt(br.readLine());
            for (int i = 0; i < tc; i++) {
                int n = Integer.parseInt(br.readLine());
                Stone[] stones = new Stone[n];
                for (int j = 0; j < n; j++) {
                    StringTokenizer st = new StringTokenizer(br.readLine(), " ");
                    stones[j] = new Stone(Integer.parseInt(st.nextToken()), Integer.parseInt(st.nextToken()), Integer.parseInt(st.nextToken()));
                }
                int result = solve(stones);

                bw.write("Case #");
                bw.write(String.valueOf(i + 1));
                bw.write(": ");
                bw.write(String.valueOf(result));
                bw.newLine();
            }

            bw.flush();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    private static int solve(Stone[] stones) {
        int[] dp = new int[stones.length + 1];

        return 0;
    }

    private static class Stone {
        public final int second;
        public final int energy;
        public final int lose;

        public Stone(int second, int energy, int lose) {
            this.second = second;
            this.energy = energy;
            this.lose = lose;
        }
    }
}
