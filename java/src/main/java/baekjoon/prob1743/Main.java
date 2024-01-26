// https://www.acmicpc.net/problem/1743
package baekjoon.prob1743;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Deque;
import java.util.LinkedList;

public class Main {

    // offset for adjacent points in xy coordinate
    private static final int[] dx = {0, 0, 1, -1};
    private static final int[] dy = {-1, 1, 0, 0};

    public static void main(String[] args) throws IOException {
        try (
                final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        ) {
            final String[] firstInput = br.readLine().split(" ");
            final int row = Integer.parseInt(firstInput[0]);
            final int column = Integer.parseInt(firstInput[1]);
            final int numOfFood = Integer.parseInt(firstInput[2]);

            // 복도 지도
            final boolean[][] map = new boolean[row][column];
            final boolean[][] isVisited = new boolean[row][column];
            for (int i = 0; i < numOfFood; i++) {
                final String[] foodInput = br.readLine().split(" ");
                // x for row, y for column
                final int x = Integer.parseInt(foodInput[0]) - 1;
                final int y = Integer.parseInt(foodInput[1]) - 1;
                map[x][y] = true;
            }

            int maxSize = 0;

            for (int x = 0; x < row; x++) {
                for (int y = 0; y < column; y++) {
                    if (map[x][y] && !isVisited[x][y]) {
                        int linkedSize = dfsSize(x, y, map, isVisited);
                        maxSize = Math.max(maxSize, linkedSize);
                    }
                }
            }

            System.out.println(maxSize);
        }
    }

    static int dfsSize(int x, int y, boolean[][] map, boolean[][] isVisited) {
        final int row = map.length;
        final int column = map[0].length;

        Deque<Integer> dfsStack = new LinkedList<>();
        dfsStack.push(coordinateToNum(x, y));
        isVisited[x][y] = true;

        // do dfs and count the passing points
        int count = 0;
        while (!dfsStack.isEmpty()) {
            Integer current = dfsStack.pop();
            int[] coordinate = numToCoordinate(current);
            count++;

            for (int i = 0; i < 4; i++) {
                int cx = coordinate[0] + dx[i];
                int cy = coordinate[1] + dy[i];
                if (cx >= 0 && cx < row && cy >= 0 && cy < column && map[cx][cy] && !isVisited[cx][cy]) {
                    dfsStack.push(coordinateToNum(cx, cy));
                    isVisited[cx][cy] = true;
                }
            }
        }

        return count;
    }

    static int coordinateToNum(int x, int y) {
        return x * 1000 + y;
    }

    static int[] numToCoordinate(int num) {
        return new int[]{num / 1000, num % 1000};
    }

}
