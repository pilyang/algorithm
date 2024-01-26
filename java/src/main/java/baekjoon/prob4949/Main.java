package baekjoon.prob4949;
// https://www.acmicpc.net/problem/4949

import java.io.*;
import java.util.ArrayDeque;
import java.util.Deque;
import java.util.Map;

public class Main {

    private static final BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
    private static final BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));
    private static final Map<String, String> BRACKET = Map.ofEntries(
            Map.entry(")", "("),
            Map.entry("]", "[")
    );

    public static void main(String[] args) throws IOException {
        while (true){
            String input = br.readLine();
            if(input.equals(".")){
                break;
            }
            String result =validate(input.split("")) ? "yes" : "no";
            bw.write(result + "\n");
        }
        bw.close();
    }

    private static boolean validate(String[] input) {
        final Deque<String> dq = new ArrayDeque<>(input.length);

        for (int index = 0; index < input.length; index++) {
            String target = input[index];
            if (BRACKET.containsKey(target)) {
                String lastBracket = dq.pollLast();
                if (!BRACKET.get(target).equals(lastBracket)) {
                    return false;
                }
            } else if (BRACKET.containsValue(target)) {
                dq.addLast(target);
            }
        }
        return dq.isEmpty();
    }
}
