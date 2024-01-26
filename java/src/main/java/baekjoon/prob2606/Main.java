package baekjoon.prob2606;
// https://www.acmicpc.net/problem/2606

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.util.Deque;
import java.util.LinkedList;

public class Main {

    private static final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
    private static final BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));
    private static final int START_NUM = 1;

    public static void main(String[] args) throws IOException {

        final int vertexCount = Integer.parseInt(br.readLine());
        final int edgeCount = Integer.parseInt(br.readLine());

        final boolean[][] graph = new boolean[vertexCount + 1][vertexCount + 1];
        final boolean[] isVisited = new boolean[vertexCount + 1];

        for (int i = 0; i < edgeCount; i++) {
            final String[] input = br.readLine().split(" ");
            final int v1 = Integer.parseInt(input[0]);
            final int v2 = Integer.parseInt(input[1]);

            graph[v1][v2] = true;
            graph[v2][v1] = true;
        }

        final Deque<Integer> dfsStack = new LinkedList<>();

        dfsStack.push(START_NUM);

        while (!dfsStack.isEmpty()) {
            final Integer current = dfsStack.pop();
            isVisited[current] = true;

            for (int i = 1; i < vertexCount + 1; i++) {
                if (!isVisited[i] && graph[current][i]) {
                    dfsStack.push(i);
                }
            }
        }

        int infectedCount = 0;
        for (int i = 2; i < isVisited.length; i++) {
            if (isVisited[i]) {
                infectedCount++;
            }
        }

        bw.write(infectedCount + "\n");
        bw.flush();
        bw.close();
    }
}
