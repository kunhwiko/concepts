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
arr[::-1]                                    # Reversing
```

##### Searching
```python
max(arr,key=len)    # Find the longest element in list
```

##### Boyer Moore's Voting Algorithm
```python
# Keeps track of majority key and value in O(n) time and O(1) space
def majorityElement(nums):
    candidate = None
    counter = 0
        
    for num in nums:
        if counter == 0:
            candidate = num
        counter += (1 if num == candidate else -1)
    return candidate
```

##### Multiple Assignment
```python
i[1],i[2],i[3] = i[3],i[2],i[1]             # i[1] = i[3] ...
val1 = val2 = i[1]                          # val1 and val2 become i[1]

results = [1] * 10                         
results[0:10:2] = [0]*len(results[0:10:2])  # [0,1,0,1,0,1,0,1,0,1]

```

### {Dictionary}
---
##### Searching
```python
max(counter.keys(), key = counter.get)     # Find the highest value out of keys
```


### "String"
---
##### Changing
```python
s = s.replace("h","w",3)      # Creates a copy and replaces first 3 "h"s to "w"
s = s[:1] + "r" + s[2:]  # Since strings are not mutable, s[1] = "r" is not allowed 
```

##### Find / Count
```python
s.find("ll",5)       # Returns the index of where "ll" is in string after index 5
s.lower().count("e")   # Creates a copy as lowercases and then counts number of "e"s
```

##### Operations
```python
s.capitalize()          # First letter is capitalized
s.isalnum()             # Checks alphanumeric character
s.isalpha()             # Checks alphabet character
s.isdecimal()           # Checks decimal character
s.isnumeric()           # Checks numeric
```

### Python Things
---
##### Pythonic
```python
# Step 1
while p1 != p2:
    if p1 != None:
        p1 = p1.next
    else:
        p1 = headB
    if p2 != None:
        p2 = p2.next
    else:
        p2 = headA

# Step 2
while p1 != p2:
    p1 = p1.next if p1 != None else headB
    p2 = p2.next if p2 != None else headA

# Step 3
while p1 != p2:
    p1 = p1.next if p1 else headB   # p1 = headB if not p1 else p1.next
    p2 = p2.next if p2 else headA   # p2 = headA if not p2 else p2.next
```

##### Zip
```python
a = [1,2,3] 
b = [4,5,6]
for i in zip(a,b):
    print(i)     # [(1,4),(2,5),(3,6)]
```

##### Map
```python
nums = [1,2,3,4]
k = map(lambda x : x**2, nums)
print(list(k))     # [1,4,9,16]
```

### Logic Operations / Bit Manipulation 
---
##### Boolean Logic
```python
sign = (x>0) - (x<0)                  # Can easily find sign
sign = (x<0) == (y<0)                 # Can easily compare signs 
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
