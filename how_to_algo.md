### Master Theorem :
```
T(n) = a T(n/b) + O(n^c)

if log_b{a} < c --> O(n^c)
if log_b{a} = c --> O(n^c * logn)
if log_b{a} > c --> O(n^log_b{a})
```

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

##### Counting Sort
```python
```

### Search Algorithms
---
##### Binary Search 
```python
# Search with two pointers to converge faster than one pointer
def binary(arr):
    left, right = 0, len(arr)-1

    while l <= r:
        mid = (l+r)//2
        if arr[mid] < target:
            left = mid + 1
        elif arr[mid] > target:
            right = mid - 1
        else:
            return True
    return False
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

##### Breadth First Search
```python
# Explores each level first before moving to next level
# Iteration
def bfsIteration(root):
    queue = collections.deque([root])
    
    while queue:
        node = queue.popleft()
        print(node.val)
        if node.left:
            queue.append(node.left)
        if node.right:
            queue.append(node.right)
```

### Useful Algorithms
---
##### Kadane's Algorithm
```python
# Update the largest continguous sum/product of a subarray
def maxSubArray(nums):
    curr = best = nums[0]
    for i in range(1,len(nums)):
        curr = max(nums[i], curr + nums[i])
        best = max(curr, best)
    return best
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
    while fast and fast.next:
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

### Question Specific Algorithms
---
##### Boyer Moore's Voting Algorithm
```python
# The number of majority will be larger than the rest  
def majorityElement(nums):
    candidate = None
    counter = 0
        
    for num in nums:
        if counter == 0:
            candidate = num
        if num == candidate:
            counter += 1
        else:
            counter -= 1
    return candidate
```

##### Dutch National Flag
```python
# Sort a series of 0,1,2s in order   
def sortColor(nums):
    p0, p2 = 0, len(nums)-1
    curr = 0
        
    while curr <= p2:
        if nums[curr] == 0:
            nums[curr],nums[p0] = nums[p0],nums[curr]
            p0 += 1
            curr += 1
        elif nums[curr] == 2:
            nums[curr],nums[p2] = nums[p2],nums[curr]
            p2 -= 1
        else:
            curr += 1
```

##### Tarjan's Algorithm
```python
# low array will show the lowest rank the node can get to by 2 or more non-overlapping means
# if discovery[curr] < low[neighbor], curr-neighbor has no other means of connecting back to the graph 
class Solution:
    def criticalConnections(self, n, connections):
        self.graph = [[] for _ in range(n)]
        self.discovery = [0] * n
        self.low = [0] * n
        self.visited = [False] * n
        self.res = []
        
        for con in connections:
            self.graph[con[0]].append(con[1])
            self.graph[con[1]].append(con[0])
        
        self.dfs(0,-1,0)
        return self.res
        
    def dfs(self,curr,parent,time):
        self.visited[curr] = True
        self.discovery[curr] = self.low[curr] = time 
        
        for neighbor in self.graph[curr]:
            if neighbor == parent:
                continue 
            
            if not self.visited[neighbor]:
                self.dfs(neighbor,curr,time+1)
                
            self.low[curr] = min(self.low[curr],self.low[neighbor])
            if self.discovery[curr] < self.low[neighbor]:
                self.res.append([curr,neighbor])
```
