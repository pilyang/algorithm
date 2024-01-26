package baekjoon.prob14713;
// https://www.acmicpc.net/problem/14713

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

public class Main {

    private static final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
    private static final BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));

    public static void main(String[] args) throws IOException {

        final int num = Integer.parseInt(br.readLine());
        final List<Queue<String>> queues = new ArrayList<>();

        for (int i = 0; i < num; i++) {
            final Queue queue = new LinkedList();
            String[] words = br.readLine().split(" ");
            for (int j = 0; j < words.length; j++) {
                queue.add(words[j]);
            }
            queues.add(queue);
        }

        String[] sentence = br.readLine().split(" ");

        for (int i = 0; i < sentence.length; i++) {
            if (!isInFirst(queues, sentence[i])) {
                bw.write("Impossible");
                bw.close();
                return;
            }
        }
        if (queues.stream().anyMatch(q -> !q.isEmpty())) {
            bw.write("Impossible");
            bw.close();
            return;
        }
        bw.write("Possible");
        bw.close();
    }

    private static boolean isInFirst(List<Queue<String>> queues, String target) {
        for (int i = 0; i < queues.size(); i++) {
            Queue<String> queue = queues.get(i);
            if (!queue.isEmpty() && queue.element().equals(target)) {
                queue.poll();
                return true;
            }
        }
        return false;
    }
}
