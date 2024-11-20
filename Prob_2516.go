package main

import (
	"fmt"
)

func takeCharacters(s string, k int) int {
	n := len(s)
	j := 0
	ans := n
	window := 0
	count := map[byte]int{}

	for i := 0; i < n; i++ {
		count[s[i]]++
	}

	if count['a'] < k || count['b'] < k || count['c'] < k {
		return -1
	}

	for i := 0; i < n; i++ {
		count[s[i]]--
		window++

		for count[s[i]] < k {
			count[s[j]]++
			j++
			window--
		}

		if n-window < ans {
			ans = n - window
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		s        string
		k        int
		expected int
	}{
		{s: "aabaaaacaabc", k: 2, expected: 8},
		{s: "a", k: 1, expected: -1},
	}

	for _, tc := range testCases {
		fmt.Printf("Input: s = %v, k = %d\n", tc.s, tc.k)
		result := takeCharacters(tc.s, tc.k)
		fmt.Printf("Output: %v\n", result)
		fmt.Printf("Expected: %v\n\n", tc.expected)
	}
}

/*
	Pseudocode :

function minTimeToTakeKCharacters(s, k):
    if count('a' in s) < k or count('b' in s) < k or count('c' in s) < k:
        return -1

    left_window = { 'a': 0, 'b': 0, 'c': 0 }
    right_window = { 'a': 0, 'b': 0, 'c': 0 }
    n = length(s)

    for i from 0 to k-1:
        left_window[s[i]] += 1
        right_window[s[n-1-i]] += 1

    min_minutes = inf

    for i from k to n:
        if left_window['a'] >= k and left_window['b'] >= k and left_window['c'] >= k:
            min_minutes = min(min_minutes, i)
        if right_window['a'] >= k and right_window['b'] >= k and right_window['c'] >= k:
            min_minutes = min(min_minutes, i)

        left_window[s[i]] += 1
        right_window[s[n-1-i]] += 1

    return min_minutes if min_minutes != inf else -1


Intuition
The problem requires determining the minimum number of operations needed to take at least k occurrences of each character
('a', 'b', 'c') from both ends of the string. The key insight is to minimize the portion of the string that remains after
taking characters from both ends. This can be achieved by using a sliding window approach to identify the smallest segment
of the string that can be excluded while still ensuring at least k occurrences of each character are present in the remaining string.

Count Character Frequencies:

First, count the total occurrences of each character ('a', 'b', 'c') in the string.
If any character occurs less than k times, it is impossible to satisfy the condition, and the result is -1.
Sliding Window:

Use a sliding window to determine the smallest segment of the string that can be excluded.
Traverse the string while reducing the count of the current character in the window.
If the count of any character outside the window falls below k, shrink the window from the left to restore the required counts.
Update Result:

Calculate the operations required to exclude the current window, which is equal to the total string length minus the size of the window.
Keep track of the minimum operations across all valid windows.
Return Result:

Return the minimum number of operations, or -1 if it is not possible to satisfy the condition.
*/
