package baekjoon.prob11724;
// https://www.acmicpc.net/problem/11724

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.LinkedList;
import java.util.Queue;

public class Main {

    private static boolean[][] graph;
    private static boolean[] isVisited;

    public static void main(String[] args) throws IOException {
        try (
                final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        ) {
            // read graph info
            final String[] input = br.readLine().split(" ");
            final int eNum = Integer.parseInt(input[0]);
            final int vNum = Integer.parseInt(input[1]);

            // init graph and is visited
            graph = new boolean[eNum][eNum];
            isVisited = new boolean[eNum];

            // read vertex info
            for (int i = 0; i < vNum; i++) {
                final String[] vertexInput = br.readLine().split(" ");
                final int a = Integer.parseInt(vertexInput[0]) - 1;
                final int b = Integer.parseInt(vertexInput[1]) - 1;
                graph[a][b] = true;
                graph[b][a] = true;
            }

            int componentCount = 0;
            for (int i=0; i<eNum; i++) {
                if (!isVisited[i]) {
                    componentCount++;
                    bfsFrom(i);
                }
            }

            System.out.println(componentCount);
        }
    }

    private static void bfsFrom(int start) {

        final Queue<Integer> queue = new LinkedList<>();
        queue.add(start);
        isVisited[start] = true;

        while (!queue.isEmpty()) {
            int current = queue.poll();
            for (int i=0; i<graph.length; i++) {
                if (graph[current][i] && !isVisited[i]) {
                    queue.add(i);
                    isVisited[i] = true;
                }
            }
        }
    }

}
