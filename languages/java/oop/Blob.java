package oop;

public class Blob extends Shape {
    public Blob(int edges) {
		super(edges);
	}
    
    // Method Overriding: overrides the superclass method 
	public void showEdges() {
        System.out.println("I do not have a shape");
    }
}