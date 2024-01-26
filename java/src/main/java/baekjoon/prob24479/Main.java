package baekjoon.prob24479;
// https://www.acmicpc.net/problem/24479

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.Comparator;
import java.util.Deque;
import java.util.LinkedList;
import java.util.List;

public class Main {

    private static final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
    private static final BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));

    public static void main(String[] args) throws IOException {

        final String[] infoInput = br.readLine().split(" ");
        final int vertexCount = Integer.parseInt(infoInput[0]);
        final int edgeCount = Integer.parseInt(infoInput[1]);
        final int startNodeNum = Integer.parseInt(infoInput[2]);

        final boolean[] isVisited = new boolean[vertexCount + 1];

        final List<Integer>[] graph = new List[vertexCount + 1];
        for (int i=0; i<graph.length; i++) {
            graph[i] = new ArrayList<>();
        }

        final int[] result = new int[vertexCount + 1];
        int visitedOrder = 1;

        // init graph
        for (int i = 0; i < edgeCount; i++) {
            final String[] edgeInput = br.readLine().split(" ");
            final int vertex1 = Integer.parseInt(edgeInput[0]);
            final int vertex2 = Integer.parseInt(edgeInput[1]);

            graph[vertex1].add(vertex2);
            graph[vertex2].add(vertex1);
        }

        Arrays.stream(graph).forEach(edges -> Collections.sort(edges, Comparator.reverseOrder()));

        final Deque<Integer> dfsStack = new LinkedList<>();
        dfsStack.push(startNodeNum);

        while (!dfsStack.isEmpty()) {
            final Integer currentVertex = dfsStack.pop();
            if (!isVisited[currentVertex]) {
                result[currentVertex] = visitedOrder++;
                isVisited[currentVertex] = true;
            }

            List<Integer> adjacentVertex = graph[currentVertex];
            for (final int vertex : adjacentVertex) {
                if (!isVisited[vertex]) {
                    dfsStack.push(vertex);
                }
            }
        }

        for (int i = 1; i < graph.length; i++) {
            bw.write(result[i] + "\n");
        }
        bw.flush();
        bw.close();
    }
}
