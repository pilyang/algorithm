package baekjoon.prob1043;
// https://www.acmicpc.net/problem/1043
// graph, union find

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.util.StringTokenizer;

public class Main {

    // root = -1 : 아는놈
    // root = 0 : 모르는놈
    private static int[] root;

    private static int find(int x) {
        if (root[x] <= 0) {
            return x;
        }
        root[x] = find(root[x]);
        return root[x];
    }

    private static void union(int x, int y) {
        x = find(x);
        y = find(y);
        if (x == y) {
            return;
        }

        // 아는놈이 있으면 아는놈이 루트여야함
        if (root[x] == -1) {
            root[y] = x;
            return;
        }
        root[x] = y;
    }

    private static boolean isKnow(int x) {
        return root[find(x)] == -1;
    }

    public static void main(String[] args) {
        try (
                BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
                BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out))
        ) {
            StringTokenizer st = new StringTokenizer(br.readLine());
            int memberNum = Integer.parseInt(st.nextToken());
            int partyNum = Integer.parseInt(st.nextToken());

            root = new int[memberNum + 1];

            st = new StringTokenizer(br.readLine());
            int initialKnownNum = Integer.parseInt(st.nextToken());
            for (int i = 0; i < initialKnownNum; i++) {
                root[Integer.parseInt(st.nextToken())] = -1;
            }

            int[][] parties = new int[partyNum][];

            for (int n = 0; n < partyNum; n++) {
                st = new StringTokenizer(br.readLine());
                int num = Integer.parseInt(st.nextToken());
                parties[n] = new int[num];

                int prev = Integer.parseInt(st.nextToken());
                parties[n][0] = prev;
                for (int i = 1; i < num; i++) {
                    int current = Integer.parseInt(st.nextToken());
                    parties[n][i] = current;
                    union(prev, current);
                    prev = current;
                }
            }

            int result = 0;

            for (int[] participants : parties) {
                boolean allNotKnow = true;
                for (int participant : participants) {
                    if (isKnow(participant)) {
                        allNotKnow = false;
                        break;
                    }
                }
                if (allNotKnow) {
                    result++;
                }
            }

            bw.write(String.valueOf(result));
            bw.newLine();
            bw.flush();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
