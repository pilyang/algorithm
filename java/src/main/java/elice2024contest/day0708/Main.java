package elice2024contest.day0708;

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;

public class Main {

    public static void main(String[] args) {
        try (
                BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
                BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out))
        ) {
            char[] input = br.readLine().toCharArray();
            int[] nums = new int[input.length];
            for (int i = 0; i < input.length; i++) {
                nums[i] = input[i] - '0';
            }

            int[] swapIndex = findSwapIndex(nums);
            if (swapIndex == null) {
                bw.write("0");
                bw.flush();
                return;
            }

            // swap
            int temp = nums[swapIndex[0]];
            nums[swapIndex[0]] = nums[swapIndex[1]];
            nums[swapIndex[1]] = temp;

            // reverse after index swapIndex[0]
            for (int i = swapIndex[0] + 1, j = nums.length - 1; i < j; i++, j--) {
                temp = nums[i];
                nums[i] = nums[j];
                nums[j] = temp;
            }

            for (int num : nums) {
                bw.write(String.valueOf(num));
            }
            bw.flush();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }

    }

    // int{x, y}, x < y
    private static int[] findSwapIndex(int[] nums) {
        for (int i = nums.length - 2; i >= 0; i--) {
            // skip if not possible
            if (nums[i] >= nums[i + 1]) {
                continue;
            }

            // find index to swap
            // nums after index i is ordered by descending
            for (int j = i + 2; j < nums.length; j++) {
                if (nums[i] >= nums[j]) {
                    return new int[]{i, j - 1};
                }
            }
            return new int[]{i, nums.length - 1};
        }
        return null;
    }
}
