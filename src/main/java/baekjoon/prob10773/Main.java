package baekjoon.prob10773;
// https://www.acmicpc.net/problem/10773

import java.io.*;
import java.util.ArrayDeque;
import java.util.Deque;

public class Main {

    private static final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
    private static final BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));

    public static void main(String[] args) throws IOException {

        final int num = Integer.parseInt(br.readLine());

        final Deque<Integer> bankAccount = new ArrayDeque<>();

        for (int i = 0; i < num; i++) {
            int amount = Integer.parseInt(br.readLine());
            if (amount == 0) {
                bankAccount.pollLast();
            } else {
                bankAccount.addLast(amount);
            }
        }

        final int sum = bankAccount.stream()
                .mapToInt(Integer::intValue)
                .sum();
        bw.write(String.valueOf(sum));
        bw.close();
    }
}
