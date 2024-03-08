package baekjoon.prob1700;
// https://www.acmicpc.net/problem/1700
// greedy

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.util.HashSet;
import java.util.Set;
import java.util.StringTokenizer;

public class Main {

    public static void main(String[] args) {
        try (
                BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
                BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out))
        ) {
            StringTokenizer inputSt;
            inputSt = new StringTokenizer(br.readLine(), " ");
            final int n = Integer.parseInt(inputSt.nextToken());
            final int k = Integer.parseInt(inputSt.nextToken());

            inputSt = new StringTokenizer(br.readLine(), " ");
            int[] orders = new int[k];
            Set<Integer> set = new HashSet<>();
            for (int i = 0; i < k; i++) {
                orders[i] = Integer.parseInt(inputSt.nextToken());
            }

            // init
            int c = 0;
            int index = 0;
            while (c < n && index < k) {
                if (!set.contains(orders[index])) {
                    set.add(orders[index]);
                    c++;
                }
                index++;
            }

            int result = 0;

            for (; index < k; index++) {
                if (set.contains(orders[index])) {
                    continue;
                }

                // get the item to remove
                Set<Integer> removeSet = new HashSet<>(set);
                int removeIndex = index + 1;
                while (removeSet.size() > 1 && removeIndex < k) {
                    removeSet.remove(orders[removeIndex]);
                    removeIndex++;
                }
                Integer itemToRemove = removeSet.stream().findAny().get();

                // remove
                set.remove(itemToRemove);
                set.add(orders[index]);
                result++;
            }

            bw.write(String.valueOf(result));
            bw.newLine();
            bw.flush();

        } catch (IOException e) {
            System.out.println(e);
            throw new RuntimeException(e);
        }
    }
}
