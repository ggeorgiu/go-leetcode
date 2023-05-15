## Leetcode problem

[Valid Anagram](https://leetcode.com/problems/valid-anagram/)

## Solution description:

- First check if the length of the two input strings are the same. If they are not the same, return false
- Init two maps holding the number of appearances for every rune in the input list (default count 0).
- For each rune in the first string increase the number of appearances in the first map
- For each rune in the second string increase the number of appearances in the second map
- At the end, check if the first map is equal with the second map. If there is at least one entry in the first map
  that is different from the same entry in the second one return false
- Return true if no difference was found.

