## Leetcode problem

[Group Anagrams](https://leetcode.com/problems/group-anagrams/)

## Solution description

- We'll init a map that will hold the groupings for the anagrams where an entry is this:
    - key: sorted letters of the string
    - value: list of words that are anagrams
- We then parse the input list of strings and do the following:
    - Sort the letters in the current string (this is efficient since the string length is <= 100)
    - Check if there is an entry in the map for the sorted string.
    - If there is an entry append the current string to the map entry, and continue the parsing
    - If there is no entry init a new map entry with the key as the sorted string and the value
      as a list containing the current string
- At the end the grouped anagrams should be present in the map values.