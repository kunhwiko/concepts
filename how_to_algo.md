### Sorting Algorithms
---
##### Bubble Sort
```python
# Swaps adjacent elements if they are in wrong order 
# Time Complexity : O(n^2)   Space Complexity : O(1)
def bubbleSort(arr):
    for i in range(len(arr)-1):
        for j in range(len(arr)-1-i):
            if arr[j] > arr[j+1]:
                arr[j],arr[j+1] = arr[j+1],arr[j]
```

##### Selection Sort
```python
# Finds minimum element and places them in the beginning 
# Time Complexity : O(n^2)   Space Complexity : O(1)
def selectionSort(arr):
    for i in range(len(arr)-1):
        minIndex = i
        for j in range(i+1,len(arr)):
            if arr[j] < arr[minIndex]:
                minIndex = j
        arr[i],arr[minIndex] = arr[minIndex],arr[i]      
```

##### Merge Sort
```python
# Time Complexity : O(nlogn)  Space Complexity : O(n) 

def mergeSort(arr,left,right):
    if left < right:
        mid = (left+right)//2
        mergeSort(arr,left,mid)
        mergeSort(arr,mid+1,right)
        merge(arr,left,mid,right)

def merge(arr,left,mid,right):
    # find size of each half
    size1 = mid-left+1
    size2 = right-mid

    # create array for each half
    L = [0]*size1
    R = [0]*size2

    # copy element for each half
    for i in range(size1):
        L[i] = arr[left+i]
    for j in range(size2):
        R[j] = arr[mid+1+j]

    # merge left and right
    i,j,k = 0,0,left
    while i < size1 and j < size2:
        if L[i] <= R[j]:
            arr[k] = L[i]
            i += 1
        else:
            arr[k] = R[j]
            j += 1
        k += 1

    # Copy remaining elements of L
    while i < size1:
        arr[k] = L[i]
        i += 1
        k += 1

    # Copy remaining elements of R
    while j < size2:
        arr[k] = R[j]
        j += 1
        k += 1
```

##### Quick Sort 
```python

# Time Complexity : O(nlogn), worst case -> O(n^2) 
# Space Complexity : O(logn) (in place sorting, space is used for recursive calls) 
def quickSort(arr,left,right):
    if left < right:
        # index to partition
        index = partition(arr,left,right)
        quickSort(arr,left,index-1)
        quickSort(arr,index+1,right)

def partition(arr,left,right):
    # set pivot as the right-most element
    pivot = arr[right]
    i = left
    while left < right:
        if arr[left] < pivot:
            arr[i],arr[left] = arr[left],arr[i]
            i += 1
        left += 1
    arr[i],arr[right] = arr[right],arr[i]
    return i
```



### Search Algorithms
---
##### Binary Search 
```python
# Search with two pointers to converge faster than one pointer
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

##### Quick Select (Modify quick sort)
```python
# A change in quick sort, only sorts needed parts and not the entire array 
def partition(arr,left,right):
    pivot = arr[right]
    i = left

    while left < right:
        if arr[left] < pivot:
            arr[i],arr[left] = arr[left],arr[i]
            i += 1
        left +=1 
    arr[right],arr[i] = arr[i],arr[right]
    return i 

def quickSelect(arr,left,right,k):
    # if l < r condition is not needed as there is a new base case
    index = partition(arr,left,right)
    if index > k:
        return quickSelect(arr,left,index-1,k)
    elif index < k:
        return quickSelect(arr,index+1,right,k)
    else:
        return arr[index]
```

### Tree/Graph Algorithms
---
##### Depth First Search
```python
# Explores as far as possible along the branch and then backtracks
# Recursion
def dfsRecursive(root):
    if root != None:
       print(root.val)
       dfsRecursive(root.left)
       dfsRecursive(root.right)
       
# Iteration
def dfsIteration(root):
    stack = [root]
    
    while stack:
        node = stack.pop()
        print(node.val)
        if node.right:
            stack.append(node.right)
        if node.left:
            stack.append(node.left)
```

##### Breath First Search
```python
# Explores each level first before moving to next level
# Iteration
def bfsIteration(root):
    queue = [root]
    
    while queue:
        node = queue[0]
        queue.remove(node)
        print(node.val)
        if node.left:
            queue.append(node.left)
        if node.right:
            queue.append(node.right)
```

##### Floyd's Cycle Detection 
```python
# Finds beginning of cycle, if one exists
# a = distance from head to start of cycle
# b = distance from start of cycle to where slow,fast meet
# c = distance from where slow,fast meet to start of cycle
# distance fast moved = a + b + c + a, distance slow moved = a + b
# relation of fast and slow : a+b+c+a = 2 * (a+b) -> c = a
def detect(head):
    slow = fast = head 
    while fast != None and fast.next != None:
        fast = fast.next.next
        slow = slow.next 
        if slow == fast:
            break
    if fast == None or fast.next == None:
        return None
    while head != fast:
        head = head.next
        fast = fast.next
    return head    
```

### Subproblems and Dynamic Programming
---
##### Kadane's Algorithm
```python
# Update a continuous subset starting from left to right
# In this problem, find the largest continuous subarray possible
def maxSubArray(nums):
    curr_sum = best_sum = nums[0]
    for num in nums[1:]:
        curr_sum = max(num, curr_sum + num)
        best_sum = max(best_sum, curr_sum)
    return best_sum
```

##### Sliding Windows
```python
# Slides a window of size k through the array starting from left to right
# In this problem, slide a window to find the max value of a size k subarray
def maxSubArray(nums,k):
    best = sum(nums[0:k])
    for i in range(1,len(nums)-k+1):
        best = max(best,sum(nums[i:i+k-1]))
    return best
```

### Other Algorithms
---
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
