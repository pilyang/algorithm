package baekjoon.prob13335;
// https://www.acmicpc.net/problem/13335

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.Queue;

public class Main {

    private static final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
    private static final BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));

    public static void main(String[] args) throws IOException {

        String[] input = br.readLine().split(" ");
        final int length = Integer.parseInt(input[1]);
        final int maximumLoad = Integer.parseInt(input[2]);
        final Bridge bridge = new Bridge(length, maximumLoad);

        String[] truckInput = br.readLine().split(" ");
        final int[] truckWeights = Arrays.stream(truckInput)
                .mapToInt(Integer::parseInt)
                .toArray();

        for (final int truckWeight : truckWeights) {
            while (!bridge.isLoadable(truckWeight)) {
                bridge.moveTrucksForward();
            }
            bridge.enterBridge(truckWeight);
        }

        while (bridge.getCurrentLoad() != 0) {
            bridge.moveTrucksForward();
        }


        bw.write(String.valueOf(bridge.getMovingCount()));
        bw.close();
    }

    private static class Bridge {
        private final int length;
        private final int maximumLoad;
        private final Queue<Integer> queue;
        private int currentLoad;
        private int movingCount;

        private Bridge(int length, int maximumLoad) {
            this.length = length;
            this.maximumLoad = maximumLoad;
            this.queue = new LinkedList<>();
            this.currentLoad = 0;
            this.movingCount = 0;

            for (int i = 0; i < length; i++) {
                queue.add(0);
            }
        }

        public boolean isLoadable(final int weight) {
            int lastTruckWeight = queue.element();
            return weight + currentLoad - lastTruckWeight <= maximumLoad;
        }

        public void enterBridge(final int truckWeight) {
            Integer exitedWeight = queue.poll();
            currentLoad -= exitedWeight;

            queue.add(truckWeight);
            currentLoad += truckWeight;
            movingCount++;
        }

        public void moveTrucksForward() {
            enterBridge(0);
        }

        public int getCurrentLoad() {
            return currentLoad;
        }

        public int getMovingCount() {
            return movingCount;
        }
    }

}
