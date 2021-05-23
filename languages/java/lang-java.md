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

##### Packages
```
Every class in Java library belongs to a package 
  1. Help organize project or library by grouping into different functionalities 
  2. Name-scoping to prevent same class name collisions in a project 
  3. Provide security where only other classes in the same package can access written code  
```

##### Security
```
Preventing subclasses / inheritance
  1. non-public classes (no such thing as a private class) can be subclassed only by classes in the same package
  2. final classes to make methods always work the way they are intended to (no method overriding)
  3. classes with only private constructors (this prevents instantiating the class)  

Preventing method overrides 
  1. declare method as final 
```

##### Data Structures 
```
TreeSet : keep elements sorted, remove duplicates 
LinkedList : makes it easy to build stacks / queues 
LinkedHashMap : hash map that remembers the order in which elements were inserted 
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

##### Short Circuits 
```java
&&               // if left side is false, the right side will not be evaluated 
&                // both sides must be evaluated, even if the left side is false
```

##### Random
```java
Math.random() * 5      // number from 0 to 4.999..
```

##### Casting
```java
float f = 5/2f;
(int)f;

Integer.toString(someint);           // changes integer to string
Integer.parseInt(stringint);         // changes string to integer
Double.parseDouble(stringdouble);    // changes double to integer
String.valueOf(5);                   // changes to string
```

##### Format
```java
String s = String.format("%,d", 1000000000);           // change to 1,000,000,000
s = String.format("%,.2f bugs to fix", 4501.123);      // change to 4,501.12
```

##### Static 
```java
/* 
   Allows for a method to run without an instance of the class
   Static means behavior is not dependent on an instance variable 
 */

// Math class is a good example of "private constructor" with "static methods" 
Math m = new Math();                  // this is impossible because Math only has private constructors 

m.random();                           // assuming Math can be instantiated, we can call random the normal way
Math.random();                        // or even like this 
```

```
Static methods CANNOT refer to any instance variables or non-static method.

However, non-static methods can still access static methods or static variables 
```

```java
// static variables are initialized when the class is loaded, not when a new instance is made
// one value per class, rather than one value per instance 

public class Rabbit {
  // all Rabbit instances share the count variable 
  private static int count = 0;

  public Rabbit() { count++; }

  public static int getCount() { return Rabbit.count; }
}
```

##### Final 
```java
// final means the variable cannot change 
public class Foo {
  // final static variables must be initialized in one of the two ways below
  // final static variables have the following naming convention 
  public static final double PI = 3.1415;
  public static final int PI_INT;

  static {
    PI_INT = 3;
  }
}

/*
   final variable: cannot change value 
   final method: cannot override the method 
   final class: cannot be extended 
 */

public final class Bar {
  // final variables must be initialized in one of the two ways below
  final int x = 3;
  final int y;

  public Bar() { y = 4; }

  public final int getX() { return x; }
    
  // can't modify str within the method
  public final int doStuff(final String str) { return y; }
}
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
##### Autoboxing
```
Arraylists and hash maps support non-primitive objects. Post Java 5.0, primitive objects are automatically converted to wrapper objects. 

ArrayList<Integer> arr = new ArrayList<Integer>();
arr.add(3); --> primitive int type is converted to Integer type
```
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

##### Exception
```
For any method that could throw an exception, you must handle the exception (e.g. wrap a try/catch clause).

The only exception to this rule is RuntimeException (e.g. NullPointerException, NumberFormatException, DivideByZero). You may or may not decide to handle these errors. 

This is because "unchecked exceptions" are typically faults in code logic, rather than things like missing files, server failure etc.  
```

```
Superclass of an exception can catch all subclasses of the exception. 

This means superclass "Exception" can be used to catch all exceptions, but just using this isn't the best convention. 

Order your exceptions from those lower in the inheritance tree, up to those higher in the tree. 
```

```java
// exception throwing code (method that throws)
public void turnOven() throws StartFireException {
	if (thisStartsFire) {
		throw new StartFireException();
	}
}

// method that uses exception throwing code (method that catches)
public String cook() {
	try {
    System.out.println("This will still run");  
		this.turnOven();
    System.out.println("This will not run"); 
    return "Done here";                       
	} catch (StartFireException ex) {
    // all exceptions inherit this method, at least use this when you can't recover from an exception 
		ex.printStackTrace();
    return "Ouch";                      
	} finally {
    // this is where you put code that must run regardless of exception 
    // important: finally will still run even if try and catch block has a return statement! 
    this.turnOffOven();
  }
}
```

```java
// you can decide to not handle the exception with try/catch
// but rather just simply toss another "throws exception"
// now this becomes the NEW risky method 
public void cook() throws StartFireException {
  this.turnOven();
}
```

##### Generics 
```
Generics allows users to create type-safe collections, classes, methods, variables in a flexible way.

With ArrayList<Object>, we can put any object into the arraylist. 
We don't want that, but we shouldn't go around creating ArrayList<Fish>, ArrayList<Car> for each object. 

Generics act as a placeholder for the object you want to use.    
```

```java
// Generics act as a placeholder for the object you want to use

// 'E' is declared by class declaration 
// 'E' will be replaced by whatever the user inputs 
public class ArrayList<E> extends AbstractList<E> implements List<E> {
	// you can declare 'E' here only because it's been declared by the class 
	public boolean add(E o)

	// this is to declare something NOT defined in class declaration 
	// 'T' must be a subtype of Animal
	public <T extends Animal> void something(ArrayList<T> arrList)

	/* 
       By polymorphism, Dog is an Animal 
       ArrayList<Dog> is not an ArrayList<Animal> 
       therefore, we cannot pass Animal<Dog> into the parameter
	*/
	public void otherthing(ArrayList<Animal> arrList)
}
```

##### Iterator
```java
LinkedList<String> list = new LinkedList<>();
Iterator<String> iter = list.iterator();             // starting from head
Iterator<String> iter2 = list.descendingIterator();  // starting from tail

while (iter.hasNext()) {
  String s = iter.next();
  System.out.println(s);
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

##### Serializable 
```java
/* 
   if output file will be used by the Java program, serialization can optimize I/O
   otherwise, just create a plain text file  

   read official documentations on ObjectOutputStream and writeObject()
*/
public class Box implements Serializable { 
  /* 
     used for version control 
     when deserializing, this makes sure no changes have been made since serialization 
     if changes (deleting instance variable, changing non-transient to transient) exists, JVM stops deserialization 
  */
	static final long serialVersionUID = 512312516L;

	// class Present must also implement Serializable 
  private Present present = new Present();

  // transient ignores this object during the serialization process 
  transient Tag tag = new Tag();
}
```

##### BufferedWriter 
```java
// FileWriter writes each thing you pass to the file each time 
// BufferedWriter reduces that overhead by passing only when the buffer is full 

BufferedWriter writer = new BufferedWriter(new FileWriter(file));

// used to send data before buffer is full 
writer.flush();
```
