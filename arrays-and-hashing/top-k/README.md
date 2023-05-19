## Leetcode problem

[Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/)

## Solution description

- We can solve this problem by using two maps:
    - One that will count the number of appearances for each entry in the input list
    - One that will group the numbers based on their frequency in the list
- We then do a reverse parse starting at the length of the input list and at each step we'll do the following:
    - Check if our frequency map contains the current index.
        - If yes append the values in the map to our result list
        - if the length of the result list is equal with the required number we'll stop the processing
    - If the map doesn't contain it we move to the next index

Note: It is safe to do this approach since the constraints specify that there is a single/unique solution