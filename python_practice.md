# Python Code
---
### Operations 
---
##### 1. Using boolean logic
```sign = (x > 0) - (x < 0)``` #find the sign of real number x

```sign * random_integer * (random_integer < 500)``` #print a number only if conditions are met, otherwise 0  

### [List] 
---
##### 1. Ordering
```l.sort(key=len,reverse=True)``` #sort by string length from longest to shortest 

```l.sort(key=lambda x : len(x),reverse=True)``` #sort by function from most to least

##### 2. Slicing
```l[::-1]``` #from(index inclusive): to(index exclusive): by

```max(l,key=len)``` #return the maximum value based on length

##### 3. Multiple Assignment
```i[1],i[2],i[3] = i[3],i[2],i[1]``` 

### "String" 
---
##### 1. Replace
```s = s.replace("h","w",3)``` #creates a copy of string; indicates from, to, number of times (default -> all strings)

##### 2. Find
```s.find("ll",5)``` #returns the index of where "ll" is in string after index 5

##### 3. Count
```s.lower().count("example")``` #creates a copy of string as lowercases; counts the number of occurrences 

### "Data Types"
---
List : mutable, ordered
Tuple : immutable, ordered
Set : mutable, not ordered
Dictionary: immutable, not ordered
