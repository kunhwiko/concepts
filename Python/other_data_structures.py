# 1) double ended queues 

# this is basically a stack and a queue at the same time 
import collections

queue = collections.deque([2,3])
queue.append(4)
queue.appendleft(1)
queue.pop()
queue.popleft()


# 2) heaps 

# in Python, the heap package supports min heaps 
# to use a max heap, we simply multiply input values by -1 

import heapq 

# min heap 
heap = []
for num in nums:
    heapq.heappush(heap, num) # O(logn) 
    heapq.heappop(arr)        # O(logn)

# max heap 
heap = []
for num in nums:
    heapq.heappush(heap, -num) # O(logn)
    heapq.heappop(arr)         # O(logn)

heapq.heapify(arr)             # O(n)


# 3) monotonic queues 

# Monotonic queues store elements in increasing or decreasing order while preserving the original sequence
# Monotonic queues are great for finding a maximum or minimum within a given window in O(1) time

# original array: [2,3,1,4]
# --> [1,2,3,4] is in increasing order, but is not in time sequence 
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






