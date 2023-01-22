package oop;

public class Dog extends Animal implements Pet {
	String type;
	String color;
	boolean friendly;
	
	/* 
	 	super() is used to call the constructors of everything above the hierarchy 
	 	this() can be used to call a constructor within the same class 
	 	
	 	this() can be used when overloaded constructors all have to perform the same functionality 
	 	and you would like to reduce duplicate code
	 	
	 	you CANNOT use both super() and this() in the same constructor
	 */

	public Dog() {
		super(4);
		System.out.println("Bark! I am a doggo!");
		this.friendly = true;
		
		// I can do a bunch more here in this constructor 
	}
	
	public Dog(String type) {
		// calls the Dog() constructor
		this();
		this.type = type;
	}
	
	public Dog(String type, String color) {
		// calls the Dog(type) constructor
		this(type);
		this.color = color; 
	}

	public void sleep() {
		System.out.println("Sleepy doggo");
	}

	public void play() {
		System.out.println("Playful doge");
	}
	
	public String getType() {
		return this.type;
	}
}
 