package baekjoon.prob1110;
// https://www.acmicpc.net/problem/1110

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;

public class Main {

    private static final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
    private static final BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));

    public static void main(String[] args) throws IOException {

        final int initialNumber = Integer.parseInt(br.readLine());
        int count = 0;

        int number = initialNumber;
        do {
            int fn = number / 10;
            int sn = number % 10;
            int temp = fn + sn;
            number = (sn * 10) + (temp % 10);
            count++;
        } while (number != initialNumber);

        bw.write(String.valueOf(count));
        bw.close();
    }
}
