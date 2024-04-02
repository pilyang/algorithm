package kakao2024winter.donutgraph;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.Queue;
import java.util.Set;

public class Solution {
    public Map<Integer, List<Integer>> graph = new HashMap<>();

    // 1 for donut
    // 2 for line
    // 3 for 8
    public int findType(int startNode) {
        Set<Integer> vertexes = new HashSet<>();
        Set<String> edges = new HashSet<>();

        // visit all connected edges
        Queue<Integer> queue = new LinkedList<>();
        queue.add(startNode);
        while (!queue.isEmpty()) {
            int current = queue.poll();
            vertexes.add(current);

            for (Integer next : graph.getOrDefault(current, List.of())) {

                String edge = current + "," + next;
                if (edges.contains(edge)) {
                    continue;
                }
                edges.add(edge);
                queue.add(next);
            }
        }

        // check for donut
        if (vertexes.size() == edges.size()) {
            return 1;
        }

        // check for line
        if (vertexes.size() - 1 == edges.size()) {
            return 2;
        }

        // check for 8
        if (vertexes.size() == edges.size() - 1) {
            return 3;
        }
        return 0;
    }

    public int[] solution(int[][] edges) {
        int[] parentsIn = new int[edges.length + 2];
        int[] childsIn = new int[edges.length + 2];
        // init graph
        for (int[] edge : edges) {
            parentsIn[edge[0]]++;
            childsIn[edge[1]]++;
            List<Integer> to = graph.getOrDefault(edge[0], new ArrayList<Integer>());
            to.add(edge[1]);
            graph.put(edge[0], to);
        }

        int[] answer = {0, 0, 0, 0};

        // find root (added node)
        int root = 0;
        for (int i = 1; i < parentsIn.length; i++) {
            if (parentsIn[i] > 1 && childsIn[i] == 0) {
                root = i;
                break;
            }
        }
        answer[0] = root;
        for (Integer node : graph.get(root)) {
            int type = findType(node);
            answer[type]++;
        }

        return answer;
    }
}
