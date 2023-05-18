## Leetcode problem

[Two Sum](https://leetcode.com/problems/two-sum/)

## Solution description

### Naive solution (O(n)*log(n))

- Start at the beginning of the input array
- Check every value after the current one
- If no solution is found move to the next value from the array and repeat
- When the solution is found return the indexes

### Efficient solution(O(n))

- Init a map having the input array values as keys and their indexes as values
- Parse the input array starting at the first value
- At each step compute the difference between the target number and the current value
- If the result is present in the map return the current index and the one stored in the map
- Otherwise, populate the map with the current value/index pair
