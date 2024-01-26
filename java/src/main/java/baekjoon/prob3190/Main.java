package baekjoon.prob3190;
// https://www.acmicpc.net/problem/3190

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.util.Arrays;
import java.util.Deque;
import java.util.HashMap;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.Map;
import java.util.Objects;
import java.util.Set;

public class Main {

    private static final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
    private static final BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));


    public static void main(String[] args) throws IOException {

        final int size = Integer.parseInt(br.readLine());
        final int numberOfApples = Integer.parseInt(br.readLine());
        Set<Coordinate> apples = new HashSet<>();
        for (int i = 0; i < numberOfApples; i++) {
            int[] applePosition = Arrays.stream(br.readLine().split(" "))
                    .mapToInt(Integer::parseInt)
                    .toArray();
            apples.add(new Coordinate(applePosition[0], applePosition[1]));
        }

        final GameMap snakeGame = new GameMap(size, apples);

        final int commandNumber = Integer.parseInt(br.readLine());
        final Map<Integer, String> movingCommandTime = new HashMap<>();

        for (int i = 0; i < commandNumber; i++) {
            final String[] command = br.readLine().split(" ");
            movingCommandTime.put(Integer.parseInt(command[0]), command[1]);
        }

        while (snakeGame.moveForward()) {
            if (movingCommandTime.containsKey(snakeGame.getPlayTime())) {
                snakeGame.changeHeading(movingCommandTime.get(snakeGame.getPlayTime()));
            }
        }

        bw.write(String.valueOf(snakeGame.playTime));
        bw.close();
    }

    private static class GameMap {

        private final int size;
        private final Deque<Coordinate> snakeQueue;
        private final Set<Coordinate> apples;

        private Direction headingDirection;
        private int playTime;

        public GameMap(final int size, final Set<Coordinate> apples) {
            this.size = size;

            snakeQueue = new LinkedList<>();
            this.apples = new HashSet<>();
            this.apples.addAll(apples);

            snakeQueue.addLast(new Coordinate(1, 1));
            headingDirection = Direction.RIGHT;
            playTime = 0;
        }

        public boolean moveForward() {
            playTime++;
            final Coordinate currentHead = snakeQueue.getLast();
            final Coordinate target = getForwardCoordinate(currentHead);
            if (isMovable(target)) {
                snakeQueue.addLast(target);
                processForTail(target);
                return true;
            }
            return false;
        }

        private Coordinate getForwardCoordinate(Coordinate currentHead) {
            return new Coordinate(currentHead.row + headingDirection.rowOffset, currentHead.column + headingDirection.columnOffset);
        }

        public boolean isMovable(final Coordinate target) {
            return !isCrashedToWall(target) && !isCrashedToSnake(target);
        }

        private boolean isCrashedToWall(final Coordinate target) {
            return target.column > size || target.column < 1 || target.row > size || target.row < 1;
        }

        private boolean isCrashedToSnake(final Coordinate target) {
            return snakeQueue.contains(target);
        }

        private void processForTail(Coordinate target) {
            if (apples.contains(target)) {
                apples.remove(target);
            } else {
                snakeQueue.removeFirst();
            }
        }

        public void changeHeading(final String command) {
            switch (command) {
                case "L":
                    headingDirection = headingDirection.leftDirection();
                    break;
                case "D":
                    headingDirection = headingDirection.rightDirection();
                    break;
            }
        }

        public int getPlayTime() {
            return playTime;
        }

        private enum Direction {
            UP(-1, 0),
            DOWN(1, 0),
            LEFT(0, -1),
            RIGHT(0, 1);

            private final int rowOffset;
            private final int columnOffset;

            Direction(int rowOffset, int columnOffset) {
                this.rowOffset = rowOffset;
                this.columnOffset = columnOffset;
            }

            public Direction leftDirection() {
                if (this == UP) return LEFT;
                if (this == DOWN) return RIGHT;
                if (this == LEFT) return DOWN;
                if (this == RIGHT) return UP;
                return null;
            }

            public Direction rightDirection() {
                if (this == UP) return RIGHT;
                if (this == DOWN) return LEFT;
                if (this == LEFT) return UP;
                if (this == RIGHT) return DOWN;
                return null;
            }
        }
    }

    private static class Coordinate {
        private final int row;
        private final int column;

        public Coordinate(int row, int column) {
            this.row = row;
            this.column = column;
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            Coordinate that = (Coordinate) o;
            return row == that.row && column == that.column;
        }

        @Override
        public int hashCode() {
            return Objects.hash(row, column);
        }
    }
}
