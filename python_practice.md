### Data Types
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
##### 1. Ordering
```l.sort(key=len,reverse=True)``` #sorts from longest to shortest length 

```l.sort(key=lambda x : len(x),reverse=True)``` #using lambda

##### 2. Slicing
```l[::-1]``` #reverse list

```max(l,key=len)``` #maximum value based on length

##### 3. Multiple Assignment
```i[1],i[2],i[3] = i[3],i[2],i[1]``` 



### {Dict} 
---
##### 1. Search
```max(d.keys(),key = d.get)```#return the maximum value based on dictionary values



### "String" 
---
##### 1. Replace
```s = s.replace("h","w",3)``` #creates a copy of string and replaces h to w 3 times

##### 2. Find
```s.find("ll",5)``` #returns the index of where "ll" is in string after index 5

##### 3. Count
```s.lower().count("e")``` #creates a copy of string as lowercases and counts number of e 

##### 4. Appending
```s[:1] + "something" + s[1:]``` #add a string in the middle of existing string


# Python Code

### Logic Operations / Bit Manipulation 
---
##### 1. Using boolean logic
```sign = (x > 0) - (x < 0)``` #find the sign of number

```sign * random_integer * (random_integer < 500)``` #print a number only if conditions are met, otherwise 0  

##### 2. Bitwise Comparison
```~5``` #NOT operation, 0101 -> 1010

```5&3``` #AND operation,0101 & 0011 -> 0001

```5|3``` #OR operation, 0101 | 0011 -> 7

```5^3``` #XOR operation, 0101^0011 -> 6

```a^b^a``` = ```a^a^b``` = ```0^b``` = ```b``` #commutative

```mask = 0xffffffff```#32bits of 1's can be used as ```a&mask``` to make 'a' an unsigned integer 

##### 3. Changing integer to binary
```bin(x)``` #returns binary string of integer x
