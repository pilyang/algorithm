package baekjoon.prob1966;
//https://www.acmicpc.net/problem/1966

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.util.LinkedList;
import java.util.Objects;
import java.util.Queue;

public class Main {

    private static final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
    private static final BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));

    public static void main(String[] args) throws IOException {
        final int caseNum = Integer.parseInt(br.readLine());

        for (int i = 0; i < caseNum; i++) {
            Queue<Doc> queue = new LinkedList<>();
            String[] input = br.readLine().split(" ");
            int num = Integer.parseInt(input[0]);
            int target = Integer.parseInt(input[1]);
            String[] priorities = br.readLine().split(" ");

            for (int index = 0; index < num; index++) {
                queue.add(new Doc(index, Integer.parseInt(priorities[index])));
            }

            int index = findIndex(queue, target);
            bw.write(String.valueOf(index + 1));
            bw.write("\n");
        }
        bw.flush();
        bw.close();
    }

    private static int findIndex(Queue<Doc> queue, int target) {
        Doc targetDoc = new Doc(target, -1);
        int size = queue.size();

        for (int answer = 0; answer < size; answer++) {
            Doc doc = queue.poll();
            int pr = doc.priority;
            while (isNotPintable(queue, pr)) {
                queue.add(doc);
                doc = queue.poll();
                pr = doc.priority;
            }
            if (doc.equals(targetDoc)) {
                return answer;
            }
        }
        return -1;
    }

    private static boolean isNotPintable(Queue<Doc> queue, int pr) {
        return queue.stream().anyMatch(d -> d.priority > pr);
    }


    private static class Doc {
        private final int num;
        private final int priority;

        public Doc(int num, int priority) {
            this.num = num;
            this.priority = priority;
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            Doc doc = (Doc) o;
            return num == doc.num;
        }

        @Override
        public int hashCode() {
            return Objects.hash(num);
        }
    }
}
