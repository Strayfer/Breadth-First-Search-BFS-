# Bubble-Sort and Tower Hanoi Review

I tested Kevin's Bubble-Sort code, and reviewed Mahendra's Tower of Hanoi.

# Bubble Sort

I predicted that the array will sort to [1, 2, 3, 5, 7, 9]. Since the index length is now 6, it will take 5 passes to fully sort. The largest number (9) will bubble to the end in Pass 1, then 7, then 5.

I changed the values and the index length (from 5 to 6). Because the array is longer, Bubble Sort needed 5 passes instead of 4. The largest numbers (9, 7, 5) bubbled to the end one by one in the first few passes. The final sorted result is [1, 2, 3, 5, 7, 9].

# Hanoi Tower

I ran the code and it works, sorting [5, 2, 8, 1, 4] into [1, 2, 4, 5, 8] in 4 passes. The algorithm uses nested loops where the inner loop compares adjacent numbers and swaps them if the left is bigger, so the biggest number bubbles to the end each pass.

However, there is a bug. The code does not stop early if the array is already sorted, so if we run [1, 2, 3, 4, 5] it will still do 4 passes for nothing. We should add a swapped boolean flag to break out early, making the best case O(n) instead of O(n²).

I also suggest testing edge cases like empty array, single element, and duplicates. And please add the README with the 5 steps: predict, run, observe, change, explain.
