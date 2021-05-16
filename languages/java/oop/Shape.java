package oop;

//subclass "extends" superclass 

public class Shape {
    // instance variables 
    private int edges; 
    
    // constructor
    public Shape(int edges) {
        this.edges = edges; 
     }

    // methods 
    public void showEdges() {
        System.out.println(this.edges);
    }
}
