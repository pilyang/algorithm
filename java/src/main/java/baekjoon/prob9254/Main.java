package baekjoon.prob9254;
// https://www.acmicpc.net/problem/9252
// dp, lcs

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;

public class Main {

    public static void main(String[] args) {
        try (
                BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
                BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out))
        ) {
            String str1 = br.readLine();
            String str2 = br.readLine();

            int[][] dp = new int[str1.length() + 1][str2.length() + 1];
            for (int s1 = 1; s1 <= str1.length(); s1++) {
                for (int s2 = 1; s2 <= str2.length(); s2++) {
                    if (str1.charAt(s1 - 1) == str2.charAt(s2 - 1)) {
                        dp[s1][s2] = dp[s1 - 1][s2 - 1] + 1;
                    } else {
                        dp[s1][s2] = Math.max(dp[s1 - 1][s2], dp[s1][s2 - 1]);
                    }
                }
            }

            int resLength = dp[str1.length()][str2.length()];
            char[] res = new char[resLength];

            int row = str1.length();
            int col = str2.length();
            int idx = resLength - 1;
            while (row > 0 && col > 0) {
                if(dp[row][col] == dp[row-1][col]) {
                    row--;
                } else if(dp[row][col] == dp[row][col-1]) {
                    col--;
                } else {
                    res[idx] = str1.charAt(row - 1);
                    row--;
                    col--;
                    idx--;
                }
            }

            bw.write(String.valueOf(resLength));
            bw.newLine();
            for(char c : res) {
                bw.write(c);
            }
            bw.newLine();
            bw.flush();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
