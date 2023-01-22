package oop;

public class Main {
	public static void main(String[] args) {
		// abstract objects cannot be instantiated 
		// Animal a = new Animal(4);
		
		Lion l = new Lion();
		l.countLegs();
		System.out.println(l.doesSwim("hunt"));
		
		// Polymorphism
		Animal f = new Fish();
		f.countLegs();
		f.doesSwim();
		
		// 'f' references a Fish object, but is an Animal type, so hello() is undefined
		// Compiler checks the 'reference' type and not the actual 'object' type
		// f.hello();
		
		if (f instanceof Fish) {
			Fish f2 = (Fish) f;
			f2.hello();
		}
		
		Dog d = new Dog("Samoyed", "White");
		System.out.println(d.getType());
	}
}
