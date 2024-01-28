package baekjoon.prob7576;
// https://www.acmicpc.net/problem/7576

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

public class Main {

    private static final int[] dx = {0, 0, 1, -1};
    private static final int[] dy = {1, -1, 0, 0};

    private static byte[][] map;

    public static void main(String[] args) throws IOException {
        try (final BufferedReader br = new BufferedReader(new InputStreamReader(System.in))) {

            // init bfs queue
            Queue<Integer[]> bfsQueue = new LinkedList<>();

            // init input data
            final String[] input = br.readLine().split(" ");
            final int column = Integer.parseInt(input[0]);
            final int row = Integer.parseInt(input[1]);

            map = new byte[row][column];
            for (int i = 0; i < row; i++) {
                final String[] lineInput = br.readLine().split(" ");
                for (int j = 0; j < column; j++) {
                    byte current = Byte.parseByte(lineInput[j]);
                    // add to bfsQueue to reduce initial searching time
                    if (current == 1) {
                        bfsQueue.add(new Integer[]{i, j});
                    }

                    map[i][j] = current;
                }
            }

            int dayCount = 0;
            while (!bfsQueue.isEmpty()) {

                List<Integer[]> adjacentPoints = new LinkedList<>();
                while (!bfsQueue.isEmpty()) {
                    adjacentPoints.addAll(getNextAdjacentPoints(bfsQueue.poll()));
                }

                if (!adjacentPoints.isEmpty()) {
                    dayCount++;
                    bfsQueue.addAll(adjacentPoints);
                }
            }

            if (!isAllTomatoGood()) {
                dayCount = -1;
            }
            System.out.println(dayCount);
        }
    }

    private static List<Integer[]> getNextAdjacentPoints(Integer[] point) {
        List<Integer[]> adjacentPoints = new LinkedList<>();
        for (int i = 0; i < 4; i++) {
            int nx = point[0] + dx[i];
            int ny = point[1] + dy[i];
            if (nx >= 0 && nx < map.length && ny >= 0 && ny < map[0].length && map[nx][ny] == 0) {
                adjacentPoints.add(new Integer[]{nx, ny});
                map[nx][ny] = 1;
            }
        }
        return adjacentPoints;
    }

    private static boolean isAllTomatoGood() {
        for (byte[] bytes : map) {
            for (byte value : bytes) {
                if (value == 0) return false;
            }
        }
        return true;
    }
}
