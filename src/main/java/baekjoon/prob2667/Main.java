package baekjoon.prob2667;
// https://www.acmicpc.net/problem/2667

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

public class Main {

    private static final int[] dx = {0, 0, 1, -1};
    private static final int[] dy = {1, -1, 0, 0};
    private static final int COORDINATE_NUM_VALUE = 100;

    private static boolean[][] map;
    private static boolean[][] isVisited;

    public static void main(String[] args) throws IOException {
        try (
                final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        ) {
            // read map size and allocate the map and isVisited with default value false
            final int mapSize = Integer.parseInt(br.readLine());

            map = new boolean[mapSize][mapSize];
            isVisited = new boolean[mapSize][mapSize];

            // read and init the map
            for (int r = 0; r < mapSize; r++) {
                char[] charArray = br.readLine().toCharArray();
                for (int c = 0; c < mapSize; c++) {
                    if (charArray[c] == '1') {
                        map[r][c] = true;
                    }
                }
            }

            // bfs for find the adjacent area size
            final List<Integer> sizes = new LinkedList<>();
            for (int x = 0; x < mapSize; x++) {
                for (int y = 0; y < mapSize; y++) {
                    // do bfs if there is points in map and not visited yet
                    if (map[x][y] && !isVisited[x][y]) {
                        sizes.add(bfsSize(x, y));
                    }
                }
            }

            System.out.println(sizes.size());
            sizes.stream()
                    .sorted()
                    .forEach(System.out::println);
        }

    }

    private static int bfsSize(int x, int y) {

        final Queue<Integer> queue = new LinkedList<>();
        queue.add(coordinateToNum(x, y));
        isVisited[x][y] = true;

        int size = 1;

        while (!queue.isEmpty()) {
            final int[] current = numToCoordinate(queue.poll());

            for (int i = 0; i < 4; i++) {
                final int adjacentX = current[0] + dx[i];
                final int adjacentY = current[1] + dy[i];
                if (isInMap(adjacentX, adjacentY) && map[adjacentX][adjacentY] && !isVisited[adjacentX][adjacentY]) {
                    queue.add(coordinateToNum(adjacentX, adjacentY));
                    isVisited[adjacentX][adjacentY] = true;
                    size++;
                }
            }
        }

        return size;
    }

    private static int coordinateToNum(int x, int y) {
        return x * COORDINATE_NUM_VALUE + y;
    }

    private static int[] numToCoordinate(int num) {
        return new int[]{num / COORDINATE_NUM_VALUE, num % COORDINATE_NUM_VALUE};
    }

    private static boolean isInMap(int x, int y) {
        return x >= 0 && x < map.length && y >= 0 && y < map[0].length;
    }
}
