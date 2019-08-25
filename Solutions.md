# Problem: there is an array A in which every value appears twice except one value. Try to get the unique value
where :
  1. A: name of array
  2. n: len of array (n > 0)

# Solutions:
### Here we have two solutions. In this text, we discuss Time complexity and Space complexity of them

### Solution 1: using Set/ Hashtable

First, we loop to get all elems of array by using for, so this step, we have O(n) time and space complexity
In addition, we calculate sum of all elems:

` int sum = 0  // here, to make it simple, we assume that elems are interger `  
` for int i = 0, i < n, i++ { get A[i]; Set(A[i]); sum += A[i] }`

End of this step, we have a Set() which has all value in array, and sum of all elems in array A (sumA).  
(Note: in Set, one value appears once)

Second, we using for loop to calculate sum of all elem in Set(). This step, we get O(~n/2) ~ O(n) Space complexity and Time complexity.
now we have **sumA - sumSet x 2 = unique value**

Let make it clear, O(n) ~ O(n/2) so **Time complexity is O(n)**. 

Let talk about **space complexity**: you need n/2 unit of memory for Set(), n unit of memiory for array A.
And if GC hasn't freed memory of A array yet (Ex: Strong reference / a global variable or just you wanna keeping it alive - 
in C, you use delete instead of waiting for stop the world, ... maybe If you want this) 
so now you have 3/2n unit memory.
Yes, I know I know, here you can free A's memory but in C, this is very easy by using delete. How about in java, python, go, when you
must wait for gc (don't use System.gc() or something else to call direct gc because I "google" and found that this is ... [suicide](https://stackoverflow.com/questions/2414105/why-is-it-bad-practice-to-call-system-gc) )

So conclusion: 
  * Using Hashtable (like Set) / Set: I think It is fastest because this is a random array. 
  But if array is too big (or just your memory is too small) don't use it

### Solution 2: using sort (a-sort-like-externalSort)
Here, for common sorts, we have nlogn time complexsity
but I don't want to talk about common sorts, I mean all sort which must load all array to memory.
I talked a bout it. And when you read the #1 solution, maybe you can ask :" so how to sovle if array is too big or mmemory is too small ?"

So I talk about external sort, which don't need loading all array to memory.
I am sorry now, because I don't implement it. I focus on how and why we use it first, about "what is it ? " I think we can talk
later because it is simple and I set its priority lowest so I can focus on my simple Chat App


External Sort = two step : 
  1. split/separate array into many small arrays (which can be loaed on my memory)
  2. sort all small arrays
  3. use merge sort to merge all small arrays. in the end, array A is sorted
  4. compare every pair elem in array to find the unique value

We need consider this case because the problem said that array A can be too big.

If we use solution 2: 
space complexity is n/k, where k is the number of small arrays
time complexity in worst case: sort(a small array) *k + merge sort() + for(compare every pair) = nlogn*k + nlogn + n/2
= ~ O(nlogn)

... something happens here
return to step 3 "user merge ..."
what if we don't need any more sort here ?! what if we don't need merge sort ?! 
ok, in step 3: don't use merge sort ! we get every first elems in k array and pair them. (first elem in k sorted array is the smallest values)
so now by get first elem in k array we can pair it. And delete pair of elems ulti we catch a elem which cannot pair. This is unique number.

and time complexity is : sort(small array) * k + pair(n elem - in worst case: the unique elem is in tail) ~ O(nlogn)

Conlusion: solution 2 get time complexity: O(nlogn) and space complexity (n/k) which k > 0, k is long and k <= n

