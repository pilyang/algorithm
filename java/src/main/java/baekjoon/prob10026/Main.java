package baekjoon.prob10026;
// https://www.acmicpc.net/problem/10026

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.LinkedList;
import java.util.Queue;

public class Main {

    private static final int[] dx = {0, 0, 1, -1};
    private static final int[] dy = {1, -1, 0, 0};
    private static final int CONV_NUM = 100;

    private static char[][] drawing;

    public static void main(String[] args) throws IOException {
        try (final BufferedReader br = new BufferedReader(new InputStreamReader(System.in))) {

            // init map info
            int size = Integer.parseInt(br.readLine());
            drawing = new char[size][size];
            boolean[][] isVisited = new boolean[size][size];
            boolean[][] isVisitedBlind = new boolean[size][size];

            for (int i = 0; i < size; i++) {
                String input = br.readLine();
                char[] charArray = input.toCharArray();
                for (int j = 0; j < size; j++) {
                    drawing[i][j] = charArray[j];
                }
            }

            int count = 0;
            int blindCount = 0;

            for (int x = 0; x < size; x++) {
                for (int y = 0; y < size; y++) {
                    if (!isVisited[x][y]) {
                        bfsFrom(x, y, isVisited, (v1, v2) -> v1 == v2);
                        count++;
                    }
                    if (!isVisitedBlind[x][y]) {
                        bfsFrom(x, y, isVisitedBlind,
                                (v1, v2) -> {
                                    if (v1 == 'G') v1 = 'R';
                                    if (v2 == 'G') v2 = 'R';
                                    return v1 == v2;
                                });
                        blindCount++;
                    }
                }
            }

            System.out.printf("%d %d\n", count, blindCount);
        }
    }

    private static void bfsFrom(int x, int y, boolean[][] isVisited, AdjacentStrategy adjacentStrategy) {

        Queue<Integer> queue = new LinkedList<>();
        queue.add(coordinateToNum(x, y));
        isVisited[x][y] = true;

        while (!queue.isEmpty()) {
            int[] current = numToCoordinate(queue.poll());
            for (int i = 0; i < 4; i++) {
                int nx = current[0] + dx[i];
                int ny = current[1] + dy[i];
                if (nx >= 0 && nx < isVisited.length && ny >= 0 && ny < isVisited.length && !isVisited[nx][ny] &&
                        adjacentStrategy.isSame(drawing[current[0]][current[1]], drawing[nx][ny])
                ) {
                    queue.add(coordinateToNum(nx, ny));
                    isVisited[nx][ny] = true;
                }
            }
        }

    }

    private static int coordinateToNum(int x, int y) {
        return x * CONV_NUM + y;
    }

    private static int[] numToCoordinate(int num) {
        return new int[]{num / CONV_NUM, num % CONV_NUM};
    }

    @FunctionalInterface
    private interface AdjacentStrategy {
        boolean isSame(char v1, char v2);
    }

}
