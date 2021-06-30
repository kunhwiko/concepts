package oop;

/* 
 	Java does not allow for multiple inheritance due to potential ambiguities
 	e.g. What if superclass method A collides with superclass method B? Which do I use?
 	
 	Interfaces make all methods abstract, and multiple interfaces can be implemented  
 	All interface methods are implicitly public and abstract 
 	Interface methods cannot collide since there is no specific implementation, 
 	and must be implemented by the subclass
 	
 	Both abstract classes and interfaces act as necessary blueprints and also 
 	is great for polymorphism (we can pass in a parameter of type Pet rather than a specific type) 
 	
 	Use an abstract class: need to define a template for a group of subclasses 
 	and have some implementation code all subclasses can use 
 	
 	Use an interface: when there is need to define a 'role' that other classes can play,
 	regardless of where they are in the inheritance tree
 */

public interface Pet {
	
	void play();

}
