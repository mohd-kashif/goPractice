import java.util.HashMap;
import java.util.Map;

class Solution {
    public static int minimumPushes(String word) {
        int result = 0;
        if (word.length() == 0){
            return result;
        }
        if (word.length() <= 7){
            return result;
        }
        int[] numberMap = new int[7];
        Map<Character, Integer> charMap = new HashMap<>();
        for (int i = 0; i < word.length(); i++){
            if (charMap.containsKey(word.charAt(i))) {
                int press = charMap.get(word.charAt(i));
                result += press;
            } else{
                int minPress = 4;
                int minKey = getMinKey(numberMap);
                minPress = numberMap[minKey] + 1 ;
                numberMap[minKey] = numberMap[minKey] + 1;
                charMap.put(word.charAt(i), minPress);
                result += minPress;
            }
        }
        return result;
    }

    public static int getMinKey(int[] numberMap){
        int minKey = 0;
        int minVal = numberMap[0];
        for (int i = 0; i < numberMap.length; i++){
            if (numberMap[i] < minVal){
                minKey = i;
            }
        }
        return minKey;
    }

    public static void main(String[] args) {
        String word = "aabbccddeeffgghhiiiiii";
        int result = minimumPushes(word);
        System.out.println(result);
    }
}
