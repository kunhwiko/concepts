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
arr.sort(key=len, reverse=True)  #Sorts by length, from longest to shortest
arr.sort(key=lambda x: len(x), reverse=True) #Same but utilizing lambda
```
