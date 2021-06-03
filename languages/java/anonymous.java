// what are anonymous classes/functions?

// locally defined classes that are not named
// enables declaration + instantiation at the same time 
// use when local class will only be used once 

public class Hello {
    public void sayHello() {
        // inner class method
        class EnglishGreetings implements Greetings {
            public void greet() { 
                System.out.println("Hello there"); 
            }
        }
        Greetings englishGreetings = new EnglishGreetings();
    
        // anonymous class method 
        // here, we keep things simplier by not defining a new class but still implements Greetings
        // however, the class cannot be reused 
        Greetings frenchGreetings = new Greetings() {
            public void greet() { 
                System.out.println("Bonjour"); 
            }
        }
    }
}

public class HelloButton {
    public void start() {
        Button button = new Button();
        button.setText("Press Me");
        
        // Normally, the EventHandler interface must be implemented with a new class 
        // We can use an anonymous class to quickly implement the interface without having to declare a new class
        button.setAction(new EventHandler<ActionEvent>() {
            @Override
            public void handle(ActionEvent event) {
                System.out.println("Hello!");
            }
        });
    }
}
```