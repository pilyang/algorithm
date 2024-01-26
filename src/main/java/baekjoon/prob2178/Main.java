package baekjoon.prob2178;
// https://www.acmicpc.net/problem/2178

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.LinkedList;
import java.util.Objects;
import java.util.Queue;

public class Main {

    private static final int[] dx = {0, 0, 1, -1};
    private static final int[] dy = {1, -1, 0, 0};

    private static boolean[][] graph;
    private static boolean[][] isVisited;

    public static void main(String[] args) throws IOException {
        try (final BufferedReader br = new BufferedReader(new InputStreamReader(System.in))) {

            // read info and init arrays
            final String[] input = br.readLine().split(" ");
            final int c = Integer.parseInt(input[0]);
            final int r = Integer.parseInt(input[1]);

            graph = new boolean[c][r];
            isVisited = new boolean[c][r];

            // init map
            for (int x = 0; x < c; x++) {
                char[] rowInfo = br.readLine().toCharArray();
                for (int y = 0; y < r; y++) {
                    if (rowInfo[y] == '1') {
                        graph[x][y] = true;
                    }
                }
            }

            int result = bfsShortDistance(0, 0, c - 1, r - 1);
            System.out.println(result);
        }
    }

    private static int bfsShortDistance(int startX, int startY, int desX, int desY) {

        final Queue<Point> bfsQueue = new LinkedList<>();
        bfsQueue.add(new Point(startX, startY, 1));
        isVisited[startX][startY] = true;

        while (!bfsQueue.isEmpty()) {
            final Point current = bfsQueue.poll();

            for (int i = 0; i < 4; i++) {
                int nextX = current.x + dx[i];
                int nextY = current.y + dy[i];
                if (isInGraph(nextX, nextY) && graph[nextX][nextY] && !isVisited[nextX][nextY]) {
                    // if next point is same with distance, return the distance to goal and end bfs
                    if (nextX == desX && nextY == desY) {
                        return current.distance + 1;
                    }
                    // if not continue the bfs
                    bfsQueue.add(new Point(nextX, nextY, current.distance + 1));
                    isVisited[nextX][nextY] = true;
                }
            }

        }

        return 0;
    }

    private static boolean isInGraph(int x, int y) {
        return x >= 0 && x < graph.length && y >= 0 && y < graph[0].length;
    }

    private static class Point {

        int x;
        int y;
        int distance;

        public Point(int x, int y, int distance) {
            this.x = x;
            this.y = y;
            this.distance = distance;
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            Point point = (Point) o;
            return x == point.x && y == point.y;
        }

        @Override
        public int hashCode() {
            return Objects.hash(x, y);
        }
    }
}
