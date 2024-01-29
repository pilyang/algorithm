package baekjoon.prob1753;
// https://www.acmicpc.net/problem/1753

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.PriorityQueue;
import java.util.Queue;

public class Main {

    private static final long INF = Long.MAX_VALUE;

    private static long[] dist;
    private static List<Edge>[] graph;
    private static boolean[] isVisited;

    public static void main(String[] args) throws IOException {
        try (final BufferedReader br = new BufferedReader(new InputStreamReader(System.in))) {

            // read data and init
            final String[] input = br.readLine().split(" ");
            final int vertexCount = Integer.parseInt(input[0]);
            final int edgeCount = Integer.parseInt(input[1]);

            graph = new List[vertexCount];
            for (int i = 0; i < vertexCount; i++) {
                graph[i] = new ArrayList<>();
            }
            isVisited = new boolean[vertexCount];

            dist = new long[vertexCount];
            Arrays.fill(dist, INF);

            final int startIndex = Integer.parseInt(br.readLine()) - 1;

            for (int i = 0; i < edgeCount; i++) {
                final String[] edgeInput = br.readLine().split(" ");
                graph[Integer.parseInt(edgeInput[0]) - 1].add(new Edge(Integer.parseInt(edgeInput[1]) - 1, Integer.parseInt(edgeInput[2])));
            }

            // dijkstra
            dist[startIndex] = 0;
            int current = startIndex;

            // Priority Queue of Edge(target, dist[startIndex, Edge.target])
            final Queue<Edge> pq = new PriorityQueue<>();
            pq.add(new Edge(current, 0));

            while (!pq.isEmpty()) {
                Edge currentEdge;
                do {
                    currentEdge = pq.poll();
                } while (isVisited[currentEdge.target] && !pq.isEmpty());

                if (isVisited[currentEdge.target]) {
                    break;
                }

                current = currentEdge.target;
                isVisited[current] = true;

                for (Edge nextEdge : graph[current]) {
                    long newDist = dist[current] + nextEdge.cost;
                    if (newDist < dist[nextEdge.target]) {
                        dist[nextEdge.target] = newDist;
                        pq.add(new Edge(nextEdge.target, newDist));
                    }
                }
            }

            for (long cost : dist) {
                if (cost == INF) {
                    System.out.println("INF");
                } else {
                    System.out.println(cost);
                }
            }
        }
    }

    private static class Edge implements Comparable {
        int target;
        long cost;

        public Edge(int target, long cost) {
            this.target = target;
            this.cost = cost;
        }

        @Override
        public int compareTo(Object o) {
            Edge e = (Edge) o;
            return Long.compare(this.cost, e.cost);
        }
    }
}
