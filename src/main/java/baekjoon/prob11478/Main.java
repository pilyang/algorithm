package baekjoon.prob11478;
// https://www.acmicpc.net/problem/11478

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.HashSet;
import java.util.Set;

public class Main {

    private static final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

    public static void main(String[] args) throws IOException {
        final String str = br.readLine();
        final int length = str.length();

        final Set<String> set = new HashSet<>();

        for (int splitStrLenght = 1; splitStrLenght <= length; splitStrLenght++) {
            for (int startIndex = 0; startIndex + splitStrLenght <= length; startIndex++) {
                set.add(str.substring(startIndex, startIndex+splitStrLenght));
            }
        }

        System.out.println(set.size());
    }
}
