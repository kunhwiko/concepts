package oop;

// subclass "extends" superclass 
// subclass "implements" Pet
public class Fish extends Animal implements Pet {
	
	/* 
	  	Constructor Chaining: calls ALL constructors up the hierarchy
	  	
	  	superclass constructors are built before subclass constructors 
	  	child cannot exist before the parent  	
	  	super() MUST be the first statement in each constructor 
	 */
    public Fish() {
    	super();
	}
    
    // Method Overriding: overrides the superclass method 
	public void countLegs() {
        System.out.println("I do not have legs");
    }

	public void sleep() {
		System.out.println("I like to sleep for 16 hours a day");
	}

	public void play() {
		System.out.println("Hello, I am here to play with you");	
	}
	
	public void hello() {
		System.out.println("Hello, I am a fish");
	}
}