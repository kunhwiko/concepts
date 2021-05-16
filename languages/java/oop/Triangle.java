package oop;

//inheritance 
public class Triangle extends Shape {
    public Triangle(int edges) {
		super(edges);
	}

	public void showEdges() {
        // perform superclass version of method, then add subclass version 
		System.out.println("I am a triangle");
		super.showEdges();
    } 
}

