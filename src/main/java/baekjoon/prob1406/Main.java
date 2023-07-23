package baekjoon.prob1406;
// https://www.acmicpc.net/problem/1406

import java.io.*;
import java.util.ArrayDeque;
import java.util.Deque;
import java.util.List;
import java.util.Objects;

public class Main {

    private static final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
    private static final BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));

    public static void main(String[] args) throws IOException {

        final String initialText = br.readLine();
        final int num = Integer.parseInt(br.readLine());

        final TextEditor textEditor = new TextEditor(initialText);


        for (int i = 0; i < num; i++) {
            String[] command = br.readLine().split(" ");
            if (command[0].equals("L")) {
                textEditor.moveCursorToLeft();
            } else if (command[0].equals("D")) {
                textEditor.moveCursorToRight();
            } else if (command[0].equals("B")) {
                textEditor.delete();
            } else if (command[0].equals("P")) {
                textEditor.write(command[1]);
            }
        }

        final String contents = textEditor.getContents();
        bw.write(contents);
        bw.close();
    }

    private static class TextEditor {
        private final Deque<String> beforeCursor;
        private final Deque<String> afterCursor;

        public TextEditor(String initialText) {
            beforeCursor = new ArrayDeque<>();
            afterCursor = new ArrayDeque<>();

            beforeCursor.addAll(List.of(initialText.split("")));
        }

        public void moveCursorToLeft() {
            String temp = beforeCursor.pollLast();
            if(Objects.nonNull(temp)){
                afterCursor.addLast(temp);
            }
        }

        public void moveCursorToRight() {
            String temp = afterCursor.pollLast();
            if(Objects.nonNull(temp)) {
                beforeCursor.addLast(temp);
            }
        }

        public void write(String s) {
            beforeCursor.addLast(s);
        }

        public void delete() {
            beforeCursor.pollLast();
        }

        public String getContents() {
            while (true) {
                String s = afterCursor.pollLast();
                if (Objects.isNull(s)) {
                    break;
                }
                beforeCursor.addLast(s);
            }

            return String.join("", beforeCursor);
        }
    }
}
