package baekjoon.prob18258;
// https://www.acmicpc.net/problem/18258

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;

public class Main {

    private static final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
    private static final BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));

    public static void main(String[] args) throws IOException {

        final int actionNum = Integer.parseInt(br.readLine());
        final MyQueue queue = new MyQueue();

        final StringBuilder sb = new StringBuilder();

        for (int i = 0; i < actionNum; i++) {
            final String[] command = br.readLine().split(" ");
            switch (command[0]) {
                case "push":
                    queue.push(Integer.parseInt(command[1]));
                    break;
                case "pop":
                    int pop = queue.pop();
                    sb.append(pop).append("\n");
                    break;
                case "size":
                    int size = queue.size();
                    sb.append(size).append("\n");
                    break;
                case "empty":
                    int empty = queue.empty();
                    sb.append(empty).append("\n");
                    break;
                case "front":
                    int front = queue.front();
                    sb.append(front).append("\n");
                    break;
                case "back":
                    int back = queue.back();
                    sb.append(back).append("\n");
                    break;
            }
        }
        bw.write(sb.toString());
        bw.flush();
        bw.close();
    }


    private static class MyQueue {

        private Node first;
        private Node last;
        private int size;

        public MyQueue() {
            this.first = null;
            this.last = null;
            size = 0;
        }

        public void push(final int num) {
            size++;
            final Node node = new Node(num);
            if (first == null) {
                first = node;
                last = node;
                return;
            }
            last.setNext(node);
            last = node;
        }

        public int pop() {
            if (first == null) {
                return -1;
            }

            size--;
            final int value = first.value;
            first = first.next;
            if (first == null) {
                last = null;
            }
            return value;
        }

        public int size() {
            return size;
        }

        public int empty() {
            return (first == null) ? 1 : 0;
        }

        public int front() {
            if (first == null) {
                return -1;
            }
            return first.value;
        }

        public int back() {
            if (last == null) {
                return -1;
            }
            return last.value;
        }

        private static class Node {
            private final int value;
            private Node next;

            private Node(final int value) {
                this.value = value;
                this.next = null;
            }

            public void setNext(final Node next) {
                this.next = next;
            }
        }

    }
}
