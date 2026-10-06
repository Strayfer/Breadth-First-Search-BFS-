# Review of Mahendra's Hanoi Tower

I ran the code and it works, sorting [5, 2, 8, 1, 4] into [1, 2, 4, 5, 8] in 4 passes. The algorithm uses nested loops where the inner loop compares adjacent numbers and swaps them if the left is bigger, so the biggest number bubbles to the end each pass.

However, there is a bug. The code does not stop early if the array is already sorted, so if we run [1, 2, 3, 4, 5] it will still do 4 passes for nothing. We should add a swapped boolean flag to break out early, making the best case O(n) instead of O(n²).

I also suggest testing edge cases like empty array, single element, and duplicates. And please add the README with the 5 steps: predict, run, observe, change, explain.

