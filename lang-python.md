# Python Programming
### Basics
---
##### Overview 
1) Strong / Dynamic Typing
2) Interpreted / Object Oriented Language
3) Built from C language (CPython)
4) Pass by Object Reference (pass a copy of the reference), primitive values are Pass by Value 
5) Reference Counting Garbage Collection (deletes objects with no references)
6) Generational Garbage Collection (cycle algorithms to get rid of old unreachable objects)

##### Data Structures 
| Type          | mutable?      | ordered?    |
| ------------- |:-------------:| ----------: |
| List          | mutable       | ordered     |
| Tuple         | immutable     | ordered     |
| Set           | mutable       | not ordered |
| Dictionary    | mutable       | not ordered |
| String / Int  | immutable     |             |


<br />

### [List]
---
##### Sorting
```python
arr.sort(key=len, reverse=True)              # Sorts by length, from longest to shortest
arr.sort(key=lambda x: len(x), reverse=True) # Utilizing lambda

arr.sort(key=lambda x : x[0])                # [(10,15),(3,4),(5,11)] -> [(3,4),(5,11),(10,15)]
```

##### Searching
```python
max(arr,key=len)                             # Find the longest element in list
```

##### Operations
```python
arr.insert(index,element)
del arr[index]
arr.remove(element)

nums = nums[::-1]                 # nums changes its reference
nums.reverse()                    # nums reverses the array its referencing
```

##### Multiple Assignment
```python
results = [1] * 10
results = [1 for _ in range(10)]
results[0:10:2] = [0]*len(results[0:10:2])  # [0,1,0,1,0,1,0,1,0,1]
```

<br />

### {Dictionary}
---
##### Searching
```python
max(counter.keys(), key = counter.get)     # Find the highest value out of keys
```

##### Deleting mappings
```python
del dict[some_item]
```

<br />

### [Other Data Structures]
---
##### Set
```python
set.add(4)
set.union(set2)
```

##### Double Ended Queue
```python
queue = collections.deque([2,3])
queue.append(4)
queue.appendleft(1)
queue.pop()
queue.popleft()
```

##### Min Heap
```python
arr = []
for num in nums:
    heapq.heappush(arr,num)   # O(logn)
    heapq.heappop(arr)        # O(logn)

heapq.heapify(arr)            # O(n)
```

##### Monotonic Queue 
```python
# Monotonic queue preserves increasing or decreasing order such that a[n+1] >= a[n] (or vice versa)
# Monotonic queues are great for finding a maximum or minimum within a given set or area in O(1) time
#[2,3,1,4]
# --> [1,2,3,4] is in increasing order, but is not monotonic
# --> [2,3,1,4] is in time sequence, but not in increasing order 
# --> [2,3,4] is monotonic increasing 

class Monoqueue():
    def __init__(self):
        self.monoqueue = collections.deque()
    
    def push(self,n):
        while self.monoqueue and self.monoqueue[-1] < n:
            self.monoqueue.pop()
        self.monoqueue.append(n)
    
    def peek(self):
        return self.monoqueue[0]
    
    def pop(self,n):
        if self.monoqueue[0] == n:
            return self.monoqueue.popleft()
```

<br />

### "String"
---
##### Replace / Find / Count
```python
s = s.replace("h","w",3)      # Creates a copy of string and replaces first 3 "h"s to "w"
s.find("ll",5)                # Returns the index of where "ll" is in string after index 5
s.lower().count("e")          # Creates a copy as lowercases and then counts number of "e"s
```


##### Operations
```python
s.capitalize()          # First letter is capitalized
s.isalnum()             # Checks alphanumeric character
s.isalpha()             # Checks alphabet character
s.isdecimal()           # Checks decimal character
s.isnumeric()           # Checks numeric
```

<br />

### Python Things
---
##### List Comprehension
```python
l = [i for i in range(5) if i > 3]
```

##### Multiple Inheritance
```python
class ArcticBear(Arctic, Bear, Land)
```

##### Split / Join
```python
strs = "hello there"    
strs = strs.split(" ")              # ["hello","there"]
strs = " ".join(strs)               # "hello there"
```

##### Zip
```python
a = [1,2,3] 
b = [4,5,6]
for i in zip(a,b):
    print(i)     # [(1,4),(2,5),(3,6)]
```

##### Map / Filter 
```python
nums = [1,2,3,4]
k = map(lambda x : x**2, nums)
print(list(k))     # [1,4,9,16]

nums = [1,2,3,4]
k = filter(lambda x: x%2 == 0, nums)
print(list(k)) 
```

##### *
```python
board = [[1,2,3],[4,5,6],[7,8,9]]
*board                               # unpacks values
zip(*board)                          # zip unpacked values
def rotate(board):
    board[:] = zip(*board)           # must dereference as board[:], think about stack
    board[:] = map(list,zip(*board)) # converts the tuples to lists
```

##### Comparator
```python
array.sort(cmp = comparator)                       # python
array.sort(key = functools.cmp_to_key(comparator)) # python3

def comparator(o1,o2):
    ...
```

##### Iterator
```python
nums = [1,2,3,4,5]
it = iter(nums)
print(next(it))
```

##### Generator
```python
allows functions to behave like iterators
use yield keyword to automatically return a generator 

less code than a standard iterator 
avoids storing entire sequences into memory 

def create():
    for _ in range(10):
        yield random.randint(1,10)

for num in create():
    print(num)
```

##### Random
```python
random.shuffle(nums)                         # randomly shuffles array
random.randint(1,3)                          # random number out of 1,2,3
```

<br />

### Logic Operations / Bit Manipulation 
---
##### Boolean Logic
```python
sign = (x>0) - (x<0)                  # Can easily find sign
sign = (x<0) == (y<0)                 # Can easily compare signs 
sign * random_num * (random_num<500)  # Get result only if random_num is below 500
```

##### int vs integer division
```python
-5 // 2 = -3       # looks at left integer on number line 
int(-5//2) = -2    # looks at integer closer to zero
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

##### Changing integer to different formats
```python
bin(x)    # Returns integer x as a string of binary numbers
hex(x)    # Returns integer x as a string of hex (0xffff format)
```
