package kakao2024winter.gift;

import java.util.HashMap;
import java.util.Map;

public class Solution {
    public int solution(String[] friends, String[] gifts) {
        // init nameIdx map for table index search
        Map<String, Integer> nameIdx = new HashMap<>();
        for (int idx = 0; idx < friends.length; idx++) {
            nameIdx.put(friends[idx], idx);
        }

        // make gift table
        // 선물 지수
        // 선물 주며녀 ++, 선물 받으면 --
        int[] giftScore = new int[friends.length];
        // from -> to table
        int[][] giftTable = new int[friends.length][friends.length];
        for (String gift : gifts) {
            String[] g = gift.split(" ");
            giftScore[nameIdx.get(g[0])]++;
            giftScore[nameIdx.get(g[1])]--;
            giftTable[nameIdx.get(g[0])][nameIdx.get(g[1])]++;
        }

        // calculate gift number for each person
        int[] giftCounts = new int[friends.length];
        for (int from = 0; from < friends.length; from++) {
            for (int to = from + 1; to < friends.length; to++) {
                if (giftTable[from][to] == giftTable[to][from]) {
                    if (giftScore[from] == giftScore[to]) {
                        continue;
                    }
                    if (giftScore[from] > giftScore[to]) {
                        giftCounts[from]++;
                    } else {
                        giftCounts[to]++;
                    }
                } else if (giftTable[from][to] > giftTable[to][from]) {
                    giftCounts[from]++;
                } else {
                    giftCounts[to]++;
                }
            }
        }

        int answer = 0;
        for (int i = 0; i < friends.length; i++) {
            if (answer < giftCounts[i]) {
                answer = giftCounts[i];
            }
        }
        return answer;
    }
}
