# Java Programming 

### Basics 
---
##### Overview
```
  1. Pass by value 
  2. Compiler transforms code to bytecode, which is then read by JVM
  3. Garbage collection removes unreachable items 
  4. All references for a JVM will have same size regardless of object 
```

### Variables 
---
##### Key Notes 
```
  1. Boolean and integers are incompatible 
  2. Local variables do not have a default value
```

##### Pre and Post Increment 
```java
int x = 0; int z = x++;       // z is 0 
int x = 0; int z = ++x;       // z is 1  
```

##### Random
```java
Math.random() * 5      // number from 0 to 4.999..
```


### {Array}
---
##### Initializing Arrays
```java
// Declaration  = Initialization 
String[] greetings = {"Hello","Hi","How are you"};
String[] greetings = new String[5]; 
String[] greetings = new String[]{"Hello","Hi","How are you"}; 

// at times, there is no need to declare the array 
return new int[]{0,1};
```

##### Basic Operations
```java
arr.length                    // length
Arrays.copyOfRange(arr,1,3)   // copies subarray from index 1 to 3
```

### {ArrayList}
---
##### Initializing ArrayLists
```java
// Arrays.asList(arr) changes an array to an arraylist
ArrayList<String> arrlist = new ArrayList<>(Arrays.asList(greetings));
```

##### Basic Operations
```java
arrlist.size()                            // length
arrlist.get(1)                            // find element 
arrlist.add(5,"new element")              // add new element to index 5
arrlist.remove(1)                         // remove index 1
arrlist.indexOf("Hello")                  // find the index of where "Hello" is
arrlist.contains("Hello")                 // check to see if "Hello" exists  
arrlist.toArray(new arr[arrlist.size()])  // change arraylist to array type
```


### {Hash Maps}
---
##### Initializing Hash/Sorted Maps
```java
Map<Integer, Integer> map = new HashMap<>();      // HashMap
SortedMap<String,Integer> map = new TreeMap<>();  // SortedMap
```

##### Traversing through keys and values
```java
// All key values in map
for(String k : hashmap.keySet()) 
	System.out.print(k);

// All value values in map
for(String v : hashmap.values()) 
	System.out.print(v);
```

##### Basic Operations
``` java
map.put(k,v);            // add new keys and values to hashmap
map.get(k);              // get value corresponding to key
map.remove(k);           // remove key and values
map.replace(k,i,j)       // replace value i of key k to value j
map.containsKey(k);      // check if map contains the key
map.containsValue(v);    // check if map contains the value
map.getOrDefault(k,0);   // get value or if not existant 0 
```

<br />

### "String and StringBuilder" 
---
##### String Multiplication 
```java
// Stringbuilder creates a resizeable array to mutate strings

// Create "printprintprintprint"
StringBuilder sb = new StringBuilder();
for(int i = 0; i < 4; i++){
	sb.append("print");
}
sb.toString();
```

##### String Manipulation
```java
// Change "Herlo" to "Hello"

// 1st method using substrings
String s1 = "Herlo";
s1 = s1.substring(0,2) + "l" + s1.substring(3);

// 2nd method using stringbuilder
String s1 = "Herlo";
StringBuilder sb = new StringBuilder(s1);
sb.setCharAt(2,'l');
s1 = sb.toString();
```

##### String(Object) Operations
```java
s1.compareTo(s2)                           // compare order
s1.equals(s2)                              // equality check
s1.charAt(2)                               // finds letter at index 2
s1.length()                                // find string length
Character.isLetterOrDigit(s1.charAt(0));   // determine if alphabet or number
Character.toLowerCase(s1.charAt(0));       // char to lower case, does not mutate string
s1.toCharArray()                           // converts string into a character array
s1.contains(s2)                            // see if s1 contains s2
s1.indexOf(s2)				   // check index of s2 inside s1
```

##### StringBuilder Operations
```java
StringBuilder sb = new StringBuilder("abbcd");
for (int i = 0; i < sb.length(); i++) {
    sb.delete(1,3);                        // delete index 1 and 2
    System.out.println(i);                 // unlike Python, sb.length() will change to 4 during the for loop
}
```

<br />

### Other Topics
---
##### Casting
```
float f = 5/2f;
(int)f

Integer.toString(someint);           // changes integer to string
Integer.parseInt(stringint);         // changes string to integer
Double.parseDouble(stringdouble);    // changes double to integer
String.valueOf(5);                   // changes to string
```

##### Switch
```java
public int getValue(char c) {
    switch(c){
        case 'I' : return 1;
        case 'V' : return 5;
        case 'X' : return 10;
        case 'L' : return 50;
        case 'C' : return 100;
        case 'D' : return 500;
        case 'M' : return 1000;
    }
    throw new IllegalArgumentException("Letter does not exist");
}
```

##### Iterator
```java
LinkedList<String> list = new LinkedList<>();
Iterator<String> iter = list.iterator();             // starting from head
Iterator<String> iter2 = list.descendingIterator();  // starting from tail
while(iter.hasNext()){
    String s = iter.next();
    System.out.println(s);
}
```

##### Synchronized
```java
// only allows 1 thread to use the function at a given time 
public synchronized void main () {

}   
```

### I/O Controls
---
##### Input
```java
Scanner s1 = new Scanner(System.in);
Scanner s2 = new Scanner(System.in);
length = s1.nextInt();
word = s2.next();
```

