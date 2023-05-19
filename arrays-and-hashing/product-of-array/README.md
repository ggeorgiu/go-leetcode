## Leetcode problem

[Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/)

## Solution description

This problem can efficiently be solved by parsing the input list of numbers twice, once from start to end and once vice
versa
We'll compute the result in an integer list.

First parsing, from start to end:

- starting with the prefix equal to 1
- store in the result list the current prefix
- compute the product of the prefix with the current number and store it in the prefix

What this will solve is that in the result list will contain at each position i the product of all the values that
have the index < i

Second parsing, from end to start:

- start with the postfix equal to 1
- update the result list by multiplying the value stored in it at the current index with the current postfix
- compute the product of the postfix with the current number and store it in the postfix

What this will solve is that at each step i the postfix will contain the product of all the values that have the
index > i.

At the end of the second parsing the result list should contain the final answer.

