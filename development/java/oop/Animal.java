package oop;

// abstract classes prevent classes from becoming instantiated 
// like what would an Animal object even be...?

abstract class Animal {
    // instance variables 
    private int legs; 
    
    // overloaded constructors (multiple constructors)
    public Animal() {
    	this.legs = 0;
    }
    
    public Animal(int legs) {
        this.legs = legs; 
     }

    // methods 
    public void countLegs() {
        System.out.println(this.legs);
    }
    
    public void doesSwim() {
    	System.out.println("This animal does swim");
    }
    
    // abstract methods are blueprints for what an animal does
    // these methods are intended to be overridden
    public abstract void sleep();
}
