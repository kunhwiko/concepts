package oop;

/* 
 	Java does not allow for multiple inheritance due to potential ambiguities
 	e.g. What if superclass method A collides with superclass method B? Which do I use?
 	
 	Interfaces make all methods abstract, and multiple interfaces can be implemented  
 	All interface methods are implicitly public and abstract 
 	Interface methods cannot collide since there is no specific implementation, 
 	and must be implemented on the subclass method side  
 	
 	Both abstract classes and interfaces act as necessary blueprints and also 
 	is great for polymorphism (we can pass in a parameter of type Pet rather than a specific type) 
 	
 	For class inheritance, only subclasses can be of superclass type A 
 	For interfaces, any class that implements the interface will become interface type A
 	
 	Use an abstract class: need to define a template for a group of subclasses 
 	and have some implementation code all subclasses can use 
 	
 	Use an interface: when there is need to define a 'role' that other classes can play,
 	regardless of where they are in the inheritance tree
 */

public interface Pet {
	
	void play();

}
