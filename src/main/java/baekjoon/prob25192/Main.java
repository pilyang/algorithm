package baekjoon.prob25192;
// https://www.acmicpc.net/problem/25192

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.HashSet;
import java.util.Set;

public class Main {

    private static final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

    private static final String ENTER = "ENTER";

    public static void main(String[] args) throws IOException {

        final int actionNumber = Integer.parseInt(br.readLine());
        final Set<String> helloSet = new HashSet<>();
        int helloCount = 0;

        for (int i = 0; i < actionNumber; i++) {
            final String command = br.readLine();
            if (command.equals(ENTER)) {
                helloSet.clear();
                continue;
            }
            if (helloSet.contains(command)) {
                continue;
            }
            helloSet.add(command);
            helloCount++;
        }

        System.out.println(helloCount);
    }
}
