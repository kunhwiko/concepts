package oop;

// subclass "extends" superclass 
public class Lion extends Animal {
	
    public Lion() {
    	// pass in parameters to the super constructor
		super(4);
	}

	public void countLegs() {
        // perform some function, then add superclass version of method
		System.out.println("I am a lion, and this is how many legs I have");
		super.countLegs();
    } 
	
	/* 
		Method Overloading: difference in parameters help differentiate functions with the same name 
	 	This is not the same as method overriding
	 	If and only if you change the parameters, you can also change the return type
		You can change the access level 
	 */
	protected String doesSwim(String word) {
		return "I don't swim, I prefer to " + word;
	}
	
	public void sleep() {
		System.out.println("I like to sleep for 9 hours a day");
	}
}
