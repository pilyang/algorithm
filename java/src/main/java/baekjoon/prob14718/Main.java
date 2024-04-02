package baekjoon.prob14718;
// https://www.acmicpc.net/problem/14718
// dp, knapsack

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
            int n = Integer.parseInt(st.nextToken());
            int maxTime = Integer.parseInt(st.nextToken());
            int[] times = new int[n];
            int[] scores = new int[n];

            for (int i = 0; i < n; i++) {
                st = new StringTokenizer(br.readLine());
                times[i] = Integer.parseInt(st.nextToken());
                scores[i] = Integer.parseInt(st.nextToken());
            }

            // original
            // dp[n][t] = max(dp[n - 1][t], dp[n - 1][t - times[n - 1]] + scores[n - 1])
            // dp[t] = max(dp[t], dp[t - times[i]] + scores[i]) : for reverse direction
            int[] dp = new int[maxTime + 1];
            for (int i = 0; i < n; i++) {
                for (int t = maxTime; t >= times[i]; t--) {
                    dp[t] = Math.max(dp[t], dp[t - times[i]] + scores[i]);
                }
            }

            bw.write(String.valueOf(dp[maxTime]));
            bw.flush();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
