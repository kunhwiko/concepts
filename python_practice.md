# Python Tips 
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
##### 1. Sorting
##### #sorts from longest to shortest length 

```l.sort(key=len,reverse=True)``` 

#using lambda

```l.sort(key=lambda x : len(x),reverse=True)``` 

#reverse list

```l[::-1]``` 


##### 2. Search
#maximum value based on length

```max(l,key=len)``` 



##### 3. Multiple Assignment
```i[1],i[2],i[3] = i[3],i[2],i[1]``` 



### {Dict} 
---
##### 1. Search
#return the maximum value based on dictionary values

```max(d.keys(), key = d.get)```




### "String" 
---
##### 1. Replace
#creates a copy of string and replaces h to w 3 times

```s = s.replace("h","w",3)``` 


##### 2. Find / Count
#returns the index of where "ll" is in string after index 5

```s.find("ll",5)``` 

#creates a copy of string as lowercases and counts number of e 

```s.lower().count("e")``` 


##### 3. Appending
#add a string in the middle of existing string

```s[:1] + "something" + s[1:]``` 





### Logic Operations / Bit Manipulation 
---
##### 1. Using boolean logic
#find the sign of number

```sign = (x > 0) - (x < 0)``` 

#print a number only if conditions are met, otherwise 0  

```sign * random_integer * (random_integer < 500)``` 


##### 2. Bitwise Comparison
```~5``` #NOT operation, 0101 -> 1010

```5&3``` #AND operation,0101 & 0011 -> 0001

```5|3``` #OR operation, 0101 | 0011 -> 7

```5^3``` #XOR operation, 0101^0011 -> 6

```a^b^a``` = ```a^a^b``` = ```0^b``` = ```b``` #commutative

```mask = 0xffffffff``` #32bits of 1's can be used as ```a&mask``` to make 'a' an unsigned integer 

 

##### 3. Changing integer to binary
```bin(x)``` #returns integer x as a string of binarys
