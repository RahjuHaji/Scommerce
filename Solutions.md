# Problem: there is an array A in which every value appears twice except one value. Try to get the unique value
where :
  1. A: name of array
  2. n: len of array (n > 0)

# Solutions:
### Here we have two solutions. In this text, we discuss Time complexity and Space complexity of them

### Solution: using Map

We just loop from 0 to n-1 to get all elems of the array A and store if to the Map B, follow the rule:
  1. if Map[elem] is null (elem's value is not in Map) --> sat Map['elem'] = true.
  2. if Map[elem] = true, then remove elem from Map.
Finally, the last `key:value` in Map is : `unique value` : true. Or the unique value is the key of final pair on Map

Time complexity: O(n)
Space complexity: O(3/2n) in worst-case ~ O(n)
