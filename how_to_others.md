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

###Object Oriented Programming Basics
-----

#####Inheritance
```
parent class : Dog 
child class : Retriever 

child class can inherit all data fields and methods of parent class, while defining its own fields or methods 
```

#####Polymorphism
```
Triangle shape = new Triangle() 
Shape shape = new Triangle()           class triangle is a part of class shape (polymorphism)
Shape shape = new Circle()             class circle is a part of class shape 
Triangle shape = new Shape()           a triangle is a shape, but a shape is not a triangle (invalid)
```

#####Dependency Injection 
```
Suppose you have a class car and you must also instantiate classes wheel, company, material

Wheel, company, and material are 'dependencies' of car because car uses these classes

Instead of instantiating and setting dependencies within the class, one can create a dependency from the outside 

This exterior dependency (e.g. wheel) can then be injected into the class car  
```
