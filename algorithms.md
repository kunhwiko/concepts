### Search Algorithms
---
##### Binary Search Algorithm / Two Pointers Method
```python
# Binary search uses two pointers to converge to an answer faster than one pointer
def mySqrt(x):
    left, right = 0, x
        
    while left <= right:
        mid = (left + right)//2
        if mid**2 <= x < (mid+1)**2:
            return mid 
        elif x < mid**2:
            right = mid 
        else:
            left = mid + 1 
```              
