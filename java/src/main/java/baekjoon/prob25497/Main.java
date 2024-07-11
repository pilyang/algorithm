package baekjoon.prob25497;

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
            String n = br.readLine();
            String skills = br.readLine();

            int count = 0;
            int sCount = 0, lCount = 0;

            for (char skill : skills.toCharArray()) {
                if (skill == 'S') {
                    sCount++;
                } else if (skill == 'L') {
                    lCount++;
                } else if (skill == 'K') {
                    if (sCount > 0) {
                        count++;
                        sCount--;
                    } else {
                        break;
                    }
                } else if (skill == 'R') {
                    if (lCount > 0) {
                        count++;
                        lCount--;
                    } else {
                        break;
                    }
                } else {
                    count++;
                }
            }

            bw.write(String.valueOf(count));
            bw.flush();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
