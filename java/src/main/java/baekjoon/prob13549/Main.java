package baekjoon.prob13549;
// https://www.acmicpc.net/problem/13549
// bfs

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.util.LinkedList;
import java.util.Queue;
import java.util.StringTokenizer;

public class Main {

    private static int[] dx = {-1, 1};

    public static void main(String[] args) {
        try (
                BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
                BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out))
        ) {
            StringTokenizer st = new StringTokenizer(br.readLine());
            int start = Integer.parseInt(st.nextToken());
            int goal = Integer.parseInt(st.nextToken());

            final boolean[] isVisit = new boolean[100_001];

            final Queue<Integer[]> queue = new LinkedList<>();
            queue.add(new Integer[]{start, 0});
            isVisit[start] = true;

            while (!queue.isEmpty()) {
                final Integer[] current = queue.poll();
                if (current[0] == goal) {
                    bw.write(String.valueOf(current[1]));
                    bw.newLine();
                    break;
                }

                // teleport
                if (current[0] * 2 <= 100_000 && !isVisit[current[0] * 2]) {
                    queue.add(new Integer[]{current[0] * 2, current[1]});
                    isVisit[current[0] * 2] = true;
                }
                // go
                for (int offset : dx) {
                    if (current[0] + offset <= 100_000 && current[0] + offset >= 0 && !isVisit[current[0] + offset]) {
                        queue.add(new Integer[]{current[0] + offset, current[1] + 1});
                        isVisit[current[0] + offset] = true;
                    }
                }
            }

            bw.flush();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
