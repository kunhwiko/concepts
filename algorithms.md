### Sorting Algorithms
---
##### Bubble Sort (Java)
```java
// Time Complexity : O(n^2)   Space Complexity : O(1)

public void bubbleSort(int[] arr){
    for(int i = 0; i < arr.length-1; i++){
        for(int j = 0; j < arr.length-1-i; j++){
            if(arr[j] > arr[j+1]){
                int tmp = arr[j+1];
                arr[j+1] = arr[j];
                arr[j] = tmp;
            }
        }
    }
}
```

##### Selection Sort (Java)
```java
// Time Complexity : O(n^2)   Space Complexity : O(1)

public void selectionSort(int[] arr){
    for(int i = 0; i < arr.length-1; i++){
        int minIndex = i;
        for(int j = i+1; j < arr.length; j++){
            if(arr[j] < arr[minIndex]){
                minIndex = j;
            }
        }
        int tmp = arr[i];
        arr[i] = arr[minIndex];
        arr[minIndex] = tmp;
    }
}
```

##### Merge Sort (Java)
```java
// Time Complexity : O(nlogn)  Space Complexity : O(n) 

public class MergeSort{
    public void sort(int arr[], int left, int right){
        if(left < right){
            int mid = (left+right)/2;
            sort(arr,left,mid);
            sort(arr,mid+1,right);
            merge(arr,left,mid,right);
        }
    }

    public void merge(int arr[], int left, int mid, int right){
        // find size of each half 
        int size1 = mid-left+1;
        int size2 = right-mid;

        // create array for each half
        int L[] = new int[size1];
        int R[] = new int[size2];

        // copy elements to each half
        for(int i = 0; i < size1; i++){
            L[i] = arr[left+i];
        }
        for(int j = 0; j < size2; j++){
            R[j] = arr[mid+1+j];
        }

        // merge L and R
        int i = 0, j = 0;
        int k = left;
        while(i < size1 && j < size2){
            if(L[i] <= R[j]){
                arr[k] = L[i];
                i++;
            }else{
                arr[k] = R[j];
                j++;
            }
            k++;
        }

        // Copy any remaining elements of L[]
        while(i < size1){
            arr[k] = L[i];
            i++;
            k++;
        }

        // Copy any remaining elements of R[]
        while(j < size2){
            arr[k] = R[j];
            j++;
            k++;
        }
    }
}
```

##### Quick Sort (Java)
```java

// Time Complexity : O(nlogn), worst case -> O(n^2)  Space Complexity : O(n) 
public class QuickSort{
    public void sort(int[] arr, int left, int right){
        if(left < right){
            // partition index
            int index = partition(arr,left,right);

            sort(arr,left,index-1);
            sort(arr,index+1,right);
        }
    }

    public int partition(int arr[], int left, int right){
        // quick sort version that takes "last element" as pivot
        int pivot = arr[right];
        int i = left;
        while(left < right){
            if(arr[left] < pivot){
                int tmp = arr[i];
                arr[i] = arr[left];
                arr[left] = tmp;
                i++;
            }
            left++;
        }

        int tmp = arr[i];
        arr[i] = arr[right];
        arr[right] = tmp;

        return i;
    }
}
```

##### Quick Select (Java)
```java
// A change in quick sort, only sorts needed parts and not the entire array 
public class QuickSelect{
    public int partition(int arr[], int left, int right){
        // quick sort version that takes "last element" as pivot
        int pivot = arr[right];
        int i = left;
        while(left < right){
            if(arr[left] < pivot){
                int tmp = arr[i];
                arr[i] = arr[left];
                arr[left] = tmp;
                i++;
            }
            left++;
        }

        int tmp = arr[i];
        arr[i] = arr[right];
        arr[right] = tmp;

        return i;
    }
 
    // k is the index of the array 
    public int select(int[] arr, int left, int right,int k){
        int index = partition(arr,left,right);
        if(index > k) {return select(arr,left,index-1,k);}
        else if(index < k) {return select(arr,index+1,right,k);}
        else{ return arr[index]; }
    }
}
```



### Search Algorithms
---
##### Binary Search Algorithm / Two Pointers (Python)
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

##### Depth First Search (Python)
```python
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
        if node.left:
            stack.append(node.left)
        if node.right:
            stack.append(node.right)
```
