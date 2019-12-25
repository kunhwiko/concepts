# Python Tips 
### Data Structures
---
| Type          | mutable?      | ordered?    |
| ------------- |:-------------:| ----------: |
| List          | mutable       | ordered     |
| Tuple         | immutable     | ordered     |
| Set           | mutable       | not ordered |
| Dictionary    | mutable       | not ordered |
| String / Int  | immutable     |             |

### [List]
---
##### Sorting
```python
arr.sort(key=len, reverse=True)              # Sorts by length, from longest to shortest
arr.sort(key=lambda x: len(x), reverse=True) # Same but utilizing lambda
```

##### Searching
```python
max(arr,key=len)    # Find the longest element in list
```

##### Binary Search Algorithm
```python
# Binary search uses two pointers to converge to an answer faster than one pointer
class Solution:
    def mySqrt(self, x):
        left, right = 0, x
        
        while(left<=right):
            mid = (left+right)//2
            if mid**2 <= x < (mid+1)**2:
                return mid 
            elif x < mid**2:
                right = mid 
            else:
                left = mid + 1 
```                

##### Kadane's Algorithm
```python
# Kadane's keeps track of current and best sum of subarrays going from left to right
class Solution:
    def maxSubArray(self, nums):
        curr_sum = best_sum = nums[0]
        for num in nums[1:]:
            curr_sum = max(num, curr_sum + num)
            best_sum = max(best_sum, curr_sum)
        return best_sum
```


##### Reversing
```python
arr[::-1]
```

##### Enumerating
```python
for ind, num in enumerate(nums):
  print(ind, num)
```

##### Multiple Assignment
```python
i[1],i[2],i[3] = i[3],i[2],i[1]    # i[1] = i[3] ...
val1 = val2 = i[1]                 # val1 and val2 become i[1]
```

### "String"
---
##### Replace
```python
s = s.replace("h","w",3)      # creates a copy and replaces first 3 "h"s to "w"
```

##### Find / Count
```python
s.find("ll",5)       # returns the index of where "ll" is in string after index 5
s.lower.count("e")   # creates a copy as lowercases and then counts number of "e"s
```

##### Changing
```python
s = s[:1] + "r" + s[2:]  # since strings are not mutable, s[1] = "r" is not allowed 
```


### Logic Operations / Bit Manipulation 
---
##### Boolean Logic
```python
sign = (x>0) - (x<0)                  # Can easily find sign 
sign * random_num * (random_num<500)  # Get result only if random_num is below 500
```
 
##### Bitwise Comparison
```python
~5                       # Not
5&3                      # And
5|3                      # Or 
5^3                      # Xor 
5^3^5 = 5^5^3 = 0^3 = 3  # Commutative
a&0xffffffff             # Mask to unsigned integer
```

##### Changing integer to binary
```python
bin(x)    # Returns integer x as a string of binary numbers
```
