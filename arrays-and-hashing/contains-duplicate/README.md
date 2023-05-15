## Leetcode problem
[Contains duplicate](https://leetcode.com/problems/contains-duplicate/)

## Solution description:
- Init a map holding the number of appearances for every number in the input list (default count 0).
- For each number check if it was previously checked. If yes return true.
- If not, increase the number of appearances for the current number with one.
- If the list is exhausted, return false.

