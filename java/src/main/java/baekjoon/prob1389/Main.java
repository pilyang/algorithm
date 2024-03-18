package baekjoon.prob1389;
// https://www.acmicpc.net/problem/1389
// graph, bfs, floyd-warshall

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.util.StringTokenizer;

public class Main {

    public static int solve(int[][] graph) {
        for (int k = 1; k < graph.length; k++) {
            for (int from = 1; from < graph.length; from++) {
                for (int to = 1; to < graph.length; to++) {
                    if (graph[from][k] != 0 && graph[k][to] != 0) {
                        if (graph[from][to] == 0) {
                            graph[from][to] = graph[from][k] + graph[k][to];
                        } else {
                            graph[from][to] = Math.min(graph[from][to], graph[from][k] + graph[k][to]);
                        }
                    }
                }
            }
        }

        int res = 0;
        int max = Integer.MAX_VALUE;
        for (int i = 1; i < graph.length; i++) {
            int sum = 0;
            for (int j = 1; j < graph.length; j++) {
                sum += graph[i][j];
            }
            if (sum < max) {
                max = sum;
                res = i;
            }
        }

        return res;
    }

    public static void main(String[] args) {
        try (
                BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
                BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out))
        ) {
            StringTokenizer st = new StringTokenizer(br.readLine());
            int v = Integer.parseInt(st.nextToken());
            int e = Integer.parseInt(st.nextToken());
            int[][] graph = new int[v + 1][v + 1];
            for (int i = 0; i < e; i++) {
                st = new StringTokenizer(br.readLine());
                int a = Integer.parseInt(st.nextToken());
                int b = Integer.parseInt(st.nextToken());
                graph[a][b] = 1;
                graph[b][a] = 1;
            }

            int answer = solve(graph);
            bw.write(String.valueOf(answer));
            bw.newLine();
            bw.flush();

        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
